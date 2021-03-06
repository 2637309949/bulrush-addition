// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgoext

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin/binding"

	addition "github.com/2637309949/bulrush-addition"
	"github.com/2637309949/bulrush-utils/funcs"
	"github.com/2637309949/bulrush-utils/regex"
	"github.com/thoas/go-funk"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

type form struct {
	Docs     []map[string]interface{} `form:"docs" json:"docs" xml:"docs" binding:"required"`
	Category interface{}              `form:"category" json:"category" xml:"category" binding:"required"`
}

func one(name string, c *gin.Context, mgo *Mongo, opts *Opts) {
	ret, err := funcs.Chain(
		func(ret interface{}) (interface{}, error) {
			one := mgo.Var(name)
			key := regex.FindStringSubmatch(":(.*)$", opts.RoutePrefixs.One(name))[0]
			id := c.Param(key)
			if !bson.IsObjectIdHex(id) {
				return nil, errors.New("not a valid id")
			}
			q := NewQuery()
			q.name = name
			q.model = one
			if err := c.BindQuery(&q.Query); err != nil {
				return nil, errors.New("not a valid id")
			}
			q.Cond = opts.RouteHooks.One.Cond(map[string]interface{}{"DeletedAt": map[string]interface{}{"$eq": nil}, "ID": map[string]interface{}{"$oid": id}}, c, struct{ Name string }{Name: name})
			return q, nil
		},
		func(ret interface{}) (interface{}, error) {
			Model := mgo.Model(name)
			one := mgo.Var(name)
			q := ret.(*Query)
			if err := q.Build(q.Cond); err != nil {
				return nil, err
			}
			if err := Model.Pipe(q.Pipe).One(one); err != nil {
				return nil, err
			}
			return one, nil
		},
	)
	if err != nil {
		addition.RushLogger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ret)
}

func list(name string, c *gin.Context, mgo *Mongo, opts *Opts) {
	ret, err := funcs.Chain(
		func(ret interface{}) (interface{}, error) {
			one := mgo.Var(name)
			q := NewQuery()
			q.name = name
			q.model = one
			if err := c.ShouldBindQuery(&q.Query); err != nil {
				return nil, err
			}
			q.Cond = opts.RouteHooks.List.Cond(map[string]interface{}{"DeletedAt": map[string]interface{}{"$eq": nil}}, c, struct{ Name string }{Name: name})
			if err := q.Build(q.Cond); err != nil {
				return nil, err
			}
			return q, nil

		},
		func(ret interface{}) (interface{}, error) {
			Model := mgo.Model(name)
			list := mgo.Vars(name)
			q := ret.(*Query)
			if err := Model.Pipe(q.Pipe).All(list); err != nil {
				return nil, err
			}
			totalrecords, _ := Model.Find(q.Cond).Count()
			if q.Range != "ALL" {
				totalpages := math.Ceil(float64(totalrecords) / float64(q.Size))
				return gin.H{
					"page":         q.Query.Page,
					"size":         q.Query.Size,
					"totalpages":   totalpages,
					"range":        q.Query.Range,
					"totalrecords": totalrecords,
					"cond":         q.Cond,
					"preload":      q.Query.Preload,
					"list":         list,
				}, nil
			}
			return gin.H{
				"range":        q.Query.Range,
				"totalrecords": totalrecords,
				"cond":         q.Cond,
				"preload":      q.Query.Preload,
				"list":         list,
			}, nil
		},
	)
	if err != nil {
		addition.RushLogger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ret)
}

func create(name string, c *gin.Context, mgo *Mongo, opts *Opts) {
	var form form
	var list = mgo.Vars(name)
	Model := mgo.Model(name)
	if err := c.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	form = opts.RouteHooks.Create.Form(form)
	docsByte, err := json.Marshal(form.Docs)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := json.Unmarshal(docsByte, list); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	list = reflect.ValueOf(list).Elem().Interface()
	docs := funk.Map(list, func(x interface{}) interface{} {
		return x
	}).([]interface{})
	if err := Model.Insert(docs...); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func remove(name string, c *gin.Context, mgo *Mongo, opts *Opts) {
	var form form
	Model := mgo.Model(name)
	if err := c.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	form = opts.RouteHooks.Delete.Form(form)
	for _, item := range form.Docs {
		_, ok := item["ID"]
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "no id found!",
			})
			return
		}
	}
	ids := funk.Map(form.Docs, func(x map[string]interface{}) bson.ObjectId {
		return bson.ObjectIdHex(x["ID"].(string))
	}).([]bson.ObjectId)

	cond := opts.RouteHooks.Delete.Cond(bson.M{"_id": bson.M{"$in": ids}}, c, struct{ Name string }{Name: name})
	if err := Model.Update(cond, bson.M{"$set": bson.M{"_deleted": time.Now()}}); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func update(name string, c *gin.Context, mgo *Mongo, opts *Opts) {
	var form form
	Model := mgo.Model(name)
	one := mgo.Var(name)
	if err := c.ShouldBindBodyWith(&form, binding.JSON); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "not exists this key in model",
		})
		return
	}
	funk.ForEach(docs, func(x interface{}) {
		out := map[string]interface{}{}
		bsonByte, err := bson.Marshal(x)
		bson.Unmarshal(bsonByte, &out)
		cond := opts.RouteHooks.Update.Cond(bson.M{"_id": out["_id"]}, c, struct{ Name string }{Name: name})
		if err = Model.Update(cond, bson.M{"$set": out}); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	})
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
