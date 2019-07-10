// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgoext

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"reflect"
	"time"

	addition "github.com/2637309949/bulrush-addition"
	"github.com/thoas/go-funk"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

type form struct {
	Docs     []map[string]interface{} `form:"docs" json:"docs" xml:"docs"`
	Category interface{}              `form:"category" json:"category" xml:"category" `
}

func one(name string, c *gin.Context, mgo *Mongo, opts *Opts) {
	Model := mgo.Model(name)
	one := mgo.Var(name)
	key := findStringSubmatch(":(.*)$", opts.RoutePrefixs.One(name))[0]
	id := c.Param(key)
	if !bson.IsObjectIdHex(id) {
		addition.RushLogger.Error("not a valid id")
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "not a valid id"})
		return
	}
	q := NewQuery()
	q.name = name
	q.model = one
	if err := c.BindQuery(q); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	cond := opts.RouteHooks.One.Cond(map[string]interface{}{"_id": map[string]interface{}{"$oid": id}}, struct{ name string }{name: name})
	fmt.Println("cond = ", cond)
	pipe, err := q.BuildPipe(cond)
	if err != nil {
		addition.RushLogger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if err = Model.Pipe(pipe).One(one); err != nil {
		addition.RushLogger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, one)
}

func list(name string, c *gin.Context, mgo *Mongo, opts *Opts) {
	var match map[string]interface{}
	Model := mgo.Model(name)
	one := mgo.Var(name)
	list := mgo.Vars(name)
	q := NewQuery()
	q.name = name
	q.model = one

	err := c.ShouldBindQuery(q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	cond := opts.RouteHooks.List.Cond(map[string]interface{}{"Deleted": map[string]interface{}{"$eq": nil}}, struct{ name string }{name: name})
	pipe, err := q.BuildPipe(cond)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = Model.Pipe(pipe).All(list)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	match, err = q.BuildCond(cond)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	totalrecords, _ := Model.Find(match).Count()
	if q.Range != "ALL" {
		totalpages := math.Ceil(float64(totalrecords) / float64(q.Size))
		c.JSON(http.StatusOK, gin.H{
			"range":        q.Range,
			"page":         q.Page,
			"totalpages":   totalpages,
			"size":         q.Size,
			"totalrecords": totalrecords,
			"cond":         q.CondMap,
			"preload":      q.Preload,
			"list":         list,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"range":        q.Range,
			"totalrecords": totalrecords,
			"cond":         q.CondMap,
			"preload":      q.Preload,
			"list":         list,
		})
	}
}

func create(name string, c *gin.Context, mgo *Mongo, opts *Opts) {
	var form form
	var list = mgo.Vars(name)
	Model := mgo.Model(name)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	form = opts.RouteHooks.Create.Form(form)
	docsByte, err := json.Marshal(form.Docs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if err := json.Unmarshal(docsByte, list); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	list = reflect.ValueOf(list).Elem().Interface()
	docs := funk.Map(list, func(x interface{}) interface{} {
		return x
	}).([]interface{})
	if err := Model.Insert(docs...); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func remove(name string, c *gin.Context, mgo *Mongo, opts *Opts) {
	var form form
	var one = mgo.Var(name)
	Model := mgo.Model(name)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	form = opts.RouteHooks.Delete.Form(form)
	docs := funk.Map(form.Docs, func(x map[string]interface{}) interface{} {
		fieldStructs := []reflect.StructField{}
		for k := range x {
			fieldStruct := findFieldStruct(reflect.TypeOf(one), k)
			if fieldStruct == nil {
				return nil
			}
			fieldStructs = append(fieldStructs, *fieldStruct)
		}
		nx := createStruct(fieldStructs)
		jsonByte, err := json.Marshal(x)
		err = json.Unmarshal(jsonByte, nx)
		if err != nil {
			return nil
		}
		return nx
	}).([]interface{})
	docs = funk.Filter(docs, func(x interface{}) bool {
		return x != nil
	}).([]interface{})

	if len(docs) != len(form.Docs) {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "not exists this key in model"})
		return
	}
	ids := funk.Map(docs, func(x interface{}) bson.ObjectId {
		out := map[string]interface{}{}
		bsonByte, err := bson.Marshal(x)
		bson.Unmarshal(bsonByte, &out)
		if err != nil {
			return bson.NewObjectId()
		}
		return out["_id"].(bson.ObjectId)
	}).([]bson.ObjectId)

	if err := Model.Update(bson.M{"_id": bson.M{"$in": ids}}, bson.M{"$set": bson.M{"_deleted": time.Now()}}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func update(name string, c *gin.Context, mgo *Mongo, opts *Opts) {
	var form form
	Model := mgo.Model(name)
	var one = mgo.Var(name)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	form = opts.RouteHooks.Update.Form(form)
	docs := funk.Map(form.Docs, func(x map[string]interface{}) interface{} {
		fieldStructs := []reflect.StructField{}
		for k := range x {
			fieldStruct := findFieldStruct(reflect.TypeOf(one), k)
			if fieldStruct == nil {
				return nil
			}
			fieldStructs = append(fieldStructs, *fieldStruct)
		}
		nx := createStruct(fieldStructs)
		jsonByte, err := json.Marshal(x)
		err = json.Unmarshal(jsonByte, nx)
		if err != nil {
			return nil
		}
		return nx
	}).([]interface{})
	docs = funk.Filter(docs, func(x interface{}) bool {
		return x != nil
	}).([]interface{})

	if len(docs) != len(form.Docs) {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "not exists this key in model"})
		return
	}
	funk.ForEach(docs, func(x interface{}) {
		out := map[string]interface{}{}
		bsonByte, err := bson.Marshal(x)
		bson.Unmarshal(bsonByte, &out)
		if err = Model.Update(bson.M{"_id": out["_id"]}, bson.M{"$set": out}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	})
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
