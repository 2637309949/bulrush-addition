// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgoext

import (
	"math"
	"net/http"
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
	ret := map[string]interface{}{}
	err = Model.Pipe(pipe).One(&ret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"pipe": pipe,
		"one":  ret,
	})
}

func list(name string, c *gin.Context, mgo *Mongo, opts *Opts) {
	var match map[string]interface{}
	Model := mgo.Model(name)
	one := mgo.Var(name)
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

	list := []map[string]interface{}{}
	err = Model.Pipe(pipe).All(&list)
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
	Model := mgo.Model(name)
	if error := c.ShouldBind(&form); error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": error.Error()})
		return
	}
	docs := funk.Map(form.Docs, func(x interface{}) interface{} {
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
	Model := mgo.Model(name)
	if error := c.ShouldBind(&form); error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": error.Error()})
		return
	}

	ids := funk.Map(form.Docs, func(item map[string]interface{}) bson.ObjectId {
		id := item["_id"].(string)
		return bson.ObjectIdHex(id)
	}).([]bson.ObjectId)
	if err := Model.Update(bson.M{"_id": bson.M{"$in": ids}}, bson.M{"$set": bson.M{"_deleted": time.Now().Unix()}}); err != nil {
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
	for _, item := range form.Docs {
		id := item["_id"].(string)
		delete(item, "_id")
		if err := Model.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": item}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
