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
	key := findStringSubmatch(":(.*)$", opts.RoutePrefixs.One(name))[0]
	id, err := strconv.Atoi(c.Param(key))
	q := NewQuery()
	if err != nil {
		addition.RushLogger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if err := c.BindQuery(&q.Query); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	cond := map[string]interface{}{"deleted_at": map[string]interface{}{"$exists": false}, "id": id}
	q.Cond = opts.RouteHooks.One.Cond(cond, struct{ name string }{name: name})

	if err := q.Build(q.Cond); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	db = db.Scopes(func(db *gorm.DB) *gorm.DB {
		if q.Select != "" {
			return db.Select(q.Select)
		}
		return db
	}).Scopes(func(db *gorm.DB) *gorm.DB {
		funk.ForEach(q.Preload, func(pre string) {
			if pre != "" {
				db = db.Preload(pre)
			}
		})
		return db
	}).Scopes(func(db *gorm.DB) *gorm.DB {
		return db.Where(q.SQL)
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
	if err := c.BindQuery(&q.Query); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	cond := map[string]interface{}{"deleted_at": map[string]interface{}{"$exists": false}}
	q.Cond = opts.RouteHooks.List.Cond(cond, struct{ name string }{name: name})

	if err := q.Build(q.Cond); err != nil {
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
		if q.Order != "" {
			db = db.Order(q.Order)
		}
		return db
	}).Scopes(func(db *gorm.DB) *gorm.DB {
		funk.ForEach(q.Preload, func(pre string) {
			if pre != "" {
				db = db.Preload(pre)
			}
		})
		return db
	}).Scopes(func(db *gorm.DB) *gorm.DB {
		if q.SQL != "" {
			db = db.Where(q.SQL)
		}
		return db
	}).Find(list)
	if err := db.Error; err != nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	if err := ext.DB.Model(one).Where(q.SQL).Count(&totalrecords).Error; err != nil {
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
			"cond":         q.Cond,
			"select":       q.Select,
			"preload":      q.Preload,
			"list":         list,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"range":        q.Range,
			"totalrecords": totalrecords,
			"cond":         q.Cond,
			"select":       q.Select,
			"preload":      q.Preload,
			"list":         list,
		})
	}
}

func create(name string, c *gin.Context, ext *GORM, opts *Opts) {
	var form form
	list := ext.Vars(name)
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
	if err := json.Unmarshal(docsByte, list); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	validate := validator.New()
	listValue := reflect.ValueOf(list).Elem()
	count := listValue.Len()
	tx := ext.DB.Begin()
	for index := 0; index < count; index++ {
		item := listValue.Index(index).Interface()
		if err := validate.Struct(item); err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		newValue := reflect.New(reflect.TypeOf(item))
		newValue.Elem().Set(reflect.ValueOf(item))
		if err := tx.Create(newValue.Interface()).Error; err != nil {
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

func remove(name string, c *gin.Context, ext *GORM, opts *Opts) {
	var form form
	bind := ext.Var(name)
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	form = opts.RouteHooks.Delete.Form(form)
	tx := ext.DB.Begin()
	for _, item := range form.Docs {
		id, ok := item["ID"]
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "no id found!"})
			return
		}
		if err := tx.Model(bind).Where("id=?", id).Update(map[string]interface{}{"deleted_at": time.Now()}).Error; err != nil {
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

func update(name string, c *gin.Context, ext *GORM, opts *Opts) {
	var form form
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	form = opts.RouteHooks.Create.Form(form)
	tx := ext.DB.Begin()
	for _, item := range form.Docs {
		one := ext.Var(name)
		id, ok := item["ID"]
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "no id found!"})
			return
		}
		// load data as default
		if tx.Where("id=?", id).First(one).Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "id=" + id.(string) + " no found"})
			return
		}
		jsonByte, err := json.Marshal(item)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		// set change value
		json.Unmarshal(jsonByte, one)
		if err := tx.Model(one).Update(one).Error; err != nil {
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
