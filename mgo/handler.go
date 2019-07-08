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
	"strings"
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
	var list = mgo.Vars(name)
	Model := mgo.Model(name)
	if error := c.ShouldBind(&form); error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": error.Error()})
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

	ids := funk.Map(reflect.ValueOf(list).Elem().Interface(), func(item interface{}) bson.ObjectId {
		if vaule := reflect.ValueOf(item).FieldByName("ID"); vaule.IsValid() {
			return vaule.Interface().(bson.ObjectId)
		}
		return bson.NewObjectId()
	}).([]bson.ObjectId)

	fmt.Println(ids)
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
	if error := c.ShouldBind(&form); error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": error.Error()})
		return
	}
	funk.ForEach(form.Docs, func(x map[string]interface{}) {
		// to struct
		var one = mgo.Var(name)
		docByte, err := json.Marshal(x)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		if err := json.Unmarshal(docByte, one); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		// to bson json
		bsonByte, err := bson.Marshal(one)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		jsonmap := bson.M{}
		if err := bson.Unmarshal(bsonByte, &jsonmap); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		// select keys
		for k := range jsonmap {
			valueType := reflect.TypeOf(one).Elem()
			field, _ := valueType.FieldByNameFunc(func(name string) bool {
				field, _ := valueType.FieldByName(name)
				tag := field.Tag.Get("bson")
				bsonName := strings.Split(tag, ",")[0]
				if bsonName == k {
					return true
				}
				return false
			})
			_, ok := x[field.Name]
			if !ok {
				delete(jsonmap, k)
			}
		}

		// update
		id := jsonmap["_id"].(bson.ObjectId)
		delete(jsonmap, "_id")
		if err := Model.Update(bson.M{"_id": id}, bson.M{"$set": jsonmap}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	})
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
