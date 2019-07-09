// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgoext

import (
	"encoding/json"
	"math"
	"net/http"
	"reflect"
	"time"

	"github.com/thoas/go-funk"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

type form struct {
	Docs     []map[string]interface{} `form:"docs" json:"docs" xml:"docs"`
	Category interface{}              `form:"category" json:"category" xml:"category" `
}

func one(name string, c *gin.Context, mgo *Mongo, opts *Opts) {
	id := c.Param("id")
	Model := mgo.Model(name)
	one := mgo.Var(name)
	if !bson.IsObjectIdHex(id) {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "not a valid id"})
		return
	}
	q := NewQuery()
	q.name = name
	q.model = one

	err := c.ShouldBindQuery(q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	pipe, err := q.BuildPipe(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = Model.Pipe(pipe).One(one)
	if err != nil {
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
	pipe, err := q.BuildPipe("")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = Model.Pipe(pipe).All(list)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	match, err = q.BuildCond("")
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
			"list":         list,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"range":        q.Range,
			"totalrecords": totalrecords,
			"cond":         q.CondMap,
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

	docsByte, err := json.Marshal(form.Docs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := json.Unmarshal(docsByte, list); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	docs := funk.Map(reflect.ValueOf(list).Elem().Interface(), func(x interface{}) interface{} {
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
