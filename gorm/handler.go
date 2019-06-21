// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gorm

import (
	"encoding/json"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"time"

	addition "github.com/2637309949/bulrush-addition"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type form struct {
	Docs     []map[string]interface{} `form:"docs" json:"docs" xml:"docs"`
	Category interface{}              `form:"category" json:"category" xml:"category" `
}

func one(name string, gorm *GORM, c *gin.Context) {
	db := gorm.DB
	one := gorm.Var(name)
	q := Query{}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		addition.RushLogger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = c.BindQuery(&q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if q.Select != "" {
		db = db.Select(q.Select)
	}
	for _, pre := range q.BuildRelated() {
		if pre != "" {
			db = db.Preload(pre)
		}
	}
	if err := db.First(one, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, one)
}

func list(name string, gorm *GORM, c *gin.Context) {
	db := gorm.DB
	q := Query{}
	totalrecords := 0
	list := gorm.Vars(name)
	one := gorm.Var(name)
	err := c.BindQuery(&q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	sql, err := q.BuildWhere()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if q.Range != "ALL" {
		q.Range = "PAGE"
		if q.Page == 0 {
			q.Page = 1
		}
		if q.Size == 0 {
			q.Size = 20
		}
		db = db.Offset((q.Page - 1) * q.Size).Limit(q.Size)
	}
	if q.Select != "" {
		db = db.Select(q.Select)
	}
	if q.BuildOrder() != "" {
		db = db.Order(q.Order)
	}
	for _, pre := range q.BuildRelated() {
		if pre != "" {
			db = db.Preload(pre)
		}
	}
	if sql != "" {
		db = db.Where(sql)
	}
	if err := db.Find(list).Error; err != nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	if err := gorm.DB.Model(one).Where(sql).Count(&totalrecords).Error; err != nil {
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	if q.Select != "" {
		list, err = q.BuildSelect(list)
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
			"where":        q.WhereMap,
			"select":       q.Select,
			"related":      q.Related,
			"list":         list,
			"sql":          sql,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"range":        q.Range,
			"totalrecords": totalrecords,
			"where":        q.WhereMap,
			"select":       q.Select,
			"related":      q.Related,
			"list":         list,
			"sql":          sql,
		})
	}
}

func create(name string, gorm *GORM, c *gin.Context) {
	var form form
	binds := gorm.Vars(name)
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
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

func remove(name string, gorm *GORM, c *gin.Context) {
	var form form
	bind := gorm.Var(name)
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	tx := gorm.DB.Begin()
	for _, item := range form.Docs {
		id, ok := item["id"]
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "no id found!"})
			return
		}
		if err := tx.Model(bind).Where("id=?", id).Update(map[string]interface{}{"deletedAt": time.Now()}).Error; err != nil {
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

func update(name string, gorm *GORM, c *gin.Context) {
	var form form
	bind := gorm.Var(name)
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
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
