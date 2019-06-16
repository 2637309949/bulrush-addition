// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgo

import (
	"encoding/json"
	"math"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

type puFormat struct {
	Cond map[string]interface{} `bson:"cond" form:"cond" json:"cond" xml:"cond"`
	Muti bool                   `bson:"muti" form:"muti" json:"muti" xml:"muti"`
	Doc  interface{}            `bson:"doc" form:"doc" json:"doc" xml:"doc" `
}

func one(name string, mgo *Mongo, c *gin.Context) {
	id := c.Param("id")
	Model := mgo.Model(name)
	one := mgo.Var(name)
	isOj := bson.IsObjectIdHex(id)
	if !isOj {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "not a valid id"})
		return
	}
	err := Model.FindId(bson.ObjectIdHex(id)).One(one)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, one)
}

func list(name string, mgo *Mongo, c *gin.Context) {
	var match map[string]interface{}
	Model := mgo.Model(name)
	list := mgo.Vars(name)
	cond := c.DefaultQuery("cond", "%7B%7D")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))
	_range := c.DefaultQuery("range", "PAGE")
	unescapeCond, err := url.QueryUnescape(cond)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = json.Unmarshal([]byte(unescapeCond), &match)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	query := Model.Find(match)
	totalrecords, _ := query.Count()
	if _range != "ALL" {
		query = query.Skip((page - 1) * size).Limit(size)
	}
	err = query.All(list)
	totalpages := math.Ceil(float64(totalrecords) / float64(size))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"range":        _range,
		"page":         page,
		"totalpages":   totalpages,
		"size":         size,
		"totalrecords": totalrecords,
		"cond":         match,
		"list":         list,
	})
}

func create(name string, mgo *Mongo, c *gin.Context) {
	Model := mgo.Model(name)
	binds := mgo.Var(name)
	if error := c.ShouldBind(&binds); error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": error.Error()})
		return
	}
	switch binds.(type) {
	case []interface{}:
	case interface{}:
		binds = []interface{}{binds}
	}
	if error := Model.Insert(binds.([]interface{})...); error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func delete(name string, mgo *Mongo, c *gin.Context) {
	Model := mgo.Model(name)
	var puDate puFormat
	if error := c.ShouldBind(&puDate); error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": error.Error()})
		return
	}
	if puDate.Muti {
		if _, error := Model.RemoveAll(puDate.Cond); error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": error.Error()})
			return
		}
	} else {
		if error := Model.Remove(puDate.Cond); error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": error.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func update(name string, mgo *Mongo, c *gin.Context) {
	Model := mgo.Model(name)
	var puDate puFormat
	if error := c.ShouldBind(&puDate); error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": error.Error()})
		return
	}

	if puDate.Muti {
		if _, error := Model.UpdateAll(puDate.Cond, puDate.Doc); error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": error.Error()})
			return
		}
	} else {
		if error := Model.Update(puDate.Cond, puDate.Doc); error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": error.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
