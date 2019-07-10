// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gormext

import (
	"encoding/json"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"time"

	addition "github.com/2637309949/bulrush-addition"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/thoas/go-funk"
	"gopkg.in/go-playground/validator.v9"
)

type form struct {
	Docs     []map[string]interface{} `form:"docs" json:"docs" xml:"docs"`
	Category interface{}              `form:"category" json:"category" xml:"category" `
}

func one(name string, c *gin.Context, ext *GORM, opts *Opts) {
	db := ext.DB
	one := ext.Var(name)
	q := NewQuery()
	key := findStringSubmatch(":(.*)$", opts.RoutePrefixs.One(name))[0]
	id, err := strconv.Atoi(c.Param(key))
	if err != nil {
		addition.RushLogger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if err = c.BindQuery(&q); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if err := q.BuildCond(map[string]interface{}{"deleted_at": map[string]interface{}{"$exists": false}, "id": id}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	q.CondMap = opts.RouteHooks.One.Cond(q.CondMap, struct{ name string }{name: name})
	sql, err := q.FlatWhere()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	db = db.Scopes(func(db *gorm.DB) *gorm.DB {
		if q.Select != "" {
			return db.Select(q.Select)
		}
		return db
	}).Scopes(func(db *gorm.DB) *gorm.DB {
		funk.ForEach(q.BuildRelated(), func(pre string) {
			if pre != "" {
				db = db.Preload(pre)
			}
		})
		return db
	}).Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where(sql)
	}).First(one)

	if err := db.Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, one)
}

func list(name string, c *gin.Context, ext *GORM, opts *Opts) {
	list := ext.Vars(name)
	one := ext.Var(name)
	totalrecords := 0
	db := ext.DB
	q := NewQuery()
	if err := c.BindQuery(q); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if err := q.BuildCond(map[string]interface{}{"deleted_at": map[string]interface{}{"$exists": false}}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	q.CondMap = opts.RouteHooks.List.Cond(q.CondMap, struct{ name string }{name: name})
	sql, err := q.FlatWhere()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	db = db.Scopes(func(db *gorm.DB) *gorm.DB {
		if q.Range != "ALL" {
			db = db.Offset((q.Page - 1) * q.Size).Limit(q.Size)
		}
		return db
	}).Scopes(func(db *gorm.DB) *gorm.DB {
		if q.Select != "" {
			db = db.Select(q.Select)
		}
		return db
	}).Scopes(func(db *gorm.DB) *gorm.DB {
		if q.BuildOrder() != "" {
			db = db.Order(q.Order)
		}
		return db
	}).Scopes(func(db *gorm.DB) *gorm.DB {
		funk.ForEach(q.BuildRelated(), func(pre string) {
			if pre != "" {
				db = db.Preload(pre)
			}
		})
		return db
	}).Scopes(func(db *gorm.DB) *gorm.DB {
		if sql != "" {
			db = db.Where(sql)
		}
		return db
	}).Find(list)
	if err := db.Error; err != nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	if err := ext.DB.Model(one).Where(sql).Count(&totalrecords).Error; err != nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	if q.Range != "ALL" {
		totalpages := math.Ceil(float64(totalrecords) / float64(q.Size))
		c.JSON(http.StatusOK, gin.H{
			"range":        q.Range,
			"page":         q.Page,
			"totalpages":   totalpages,
			"size":         q.Size,
			"totalrecords": totalrecords,
			"cond":         q.CondMap,
			"select":       q.Select,
			"preload":      q.Preload,
			"list":         list,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"range":        q.Range,
			"totalrecords": totalrecords,
			"cond":         q.CondMap,
			"select":       q.Select,
			"preload":      q.Preload,
			"list":         list,
		})
	}
}

func create(name string, c *gin.Context, gorm *GORM, opts *Opts) {
	var form form
	binds := gorm.Vars(name)
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	form = opts.RouteHooks.Create.Form(form)
	docsByte, err := json.Marshal(form.Docs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if err := json.Unmarshal(docsByte, binds); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	validate := validator.New()
	count := reflect.ValueOf(binds).Elem().Len()
	tx := gorm.DB.Begin()
	for count > 0 {
		count = count - 1
		ele := reflect.ValueOf(binds).Elem().Index(count).Interface()
		ptr := reflect.New(reflect.TypeOf(ele))
		ptr.Elem().Set(reflect.ValueOf(ele))
		err := validate.Struct(ptr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		if err := tx.Create(ptr.Interface()).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	err = tx.Commit().Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func remove(name string, c *gin.Context, gorm *GORM, opts *Opts) {
	var form form
	bind := gorm.Var(name)
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	form = opts.RouteHooks.Delete.Form(form)
	tx := gorm.DB.Begin()
	for _, item := range form.Docs {
		id, ok := item["id"]
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "no id found!"})
			return
		}
		if err := tx.Model(bind).Where("id=?", id).Update(item).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	err := tx.Commit().Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func update(name string, c *gin.Context, gorm *GORM, opts *Opts) {
	var form form
	bind := gorm.Var(name)
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	form = opts.RouteHooks.Create.Form(form)
	tx := gorm.DB.Begin()
	for _, item := range form.Docs {
		id, ok := item["id"]
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "no id found!"})
			return
		}
		item["updatedAt"] = time.Now()
		delete(item, "createdAt")
		if err := tx.Model(bind).Where("id=?", id).Update(item).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	err := tx.Commit().Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
