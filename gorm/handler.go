// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gorm

import (
	"encoding/json"
	"math"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	addition "github.com/2637309949/bulrush-addition"
	"gopkg.in/go-playground/validator.v9"

	"github.com/gin-gonic/gin"
)

type puFormat struct {
	Cond map[string]interface{} `bson:"cond" form:"cond" json:"cond" xml:"cond"`
	Muti bool                   `bson:"muti" form:"muti" json:"muti" xml:"muti"`
	Doc  interface{}            `bson:"doc" form:"doc" json:"doc" xml:"doc" `
}

func one(name string, gorm *GORM, c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		addition.RushLogger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	one := gorm.Var(name)
	if err := gorm.DB.First(one, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, one)
}

func list(name string, gorm *GORM, c *gin.Context) {
	var whereMap map[string]interface{}
	list := gorm.Vars(name)
	one := gorm.Var(name)
	where := c.DefaultQuery("where", "%7B%7D")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "20"))
	_range := c.DefaultQuery("range", "PAGE")
	unescapeWhere, err := url.QueryUnescape(where)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = json.Unmarshal([]byte(unescapeWhere), &whereMap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	query := NewQuery(whereMap, c.DefaultQuery("select", ""), c.DefaultQuery("order", ""), c.DefaultQuery("related", ""))
	sel := query.BuildSelect()
	ord := query.BuildOrder()
	rel := query.BuildRelated()
	sql, err := query.BuildWhere()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	var totalrecords int
	db := gorm.DB.Where(sql)
	if _range != "ALL" {
		db = db.Offset((page - 1) * size).Limit(size)
	}
	if sel != "" {
		db = db.Select(sel)
	}
	if ord != "" {
		db = db.Order(ord)
	}
	for _, rl := range rel {
		if rl != "" {
			db = db.Preload(rl)
		}
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
	if _range != "ALL" {
		totalpages := math.Ceil(float64(totalrecords) / float64(size))
		c.JSON(http.StatusOK, gin.H{
			"range":        _range,
			"page":         page,
			"totalpages":   totalpages,
			"size":         size,
			"totalrecords": totalrecords,
			"where":        whereMap,
			"list":         list,
			"sql":          sql,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"range":        _range,
			"totalrecords": totalrecords,
			"where":        whereMap,
			"list":         list,
			"sql":          sql,
		})
	}
}

func create(name string, gorm *GORM, c *gin.Context) {
	var form interface{}
	binds := gorm.Vars(name)
	bind := gorm.Var(name)
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	buf = buf[0:num]
	err := json.Unmarshal(buf, &form)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if reflect.TypeOf(form).Kind() == reflect.Map {
		form = []interface{}{form}
	}

	// to model
	slicev := reflect.ValueOf(binds).Elem()
	slicev = slicev.Slice(0, slicev.Cap())
	for _, f := range form.([]interface{}) {
		mf, err := json.Marshal(f)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		err = json.Unmarshal(mf, bind)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		slicev = reflect.Append(slicev, reflect.ValueOf(bind).Elem())
		slicev = slicev.Slice(0, slicev.Cap())
	}
	reflect.ValueOf(binds).Elem().Set(slicev)

	// validate model
	validate := validator.New()
	count := reflect.ValueOf(binds).Elem().Len()
	for count > 0 {
		count = count - 1
		ele := reflect.ValueOf(binds).Elem().Index(count).Interface()
		err := validate.Struct(ele)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	// save model
	count = reflect.ValueOf(binds).Elem().Len()
	for count > 0 {
		count = count - 1
		ele := reflect.ValueOf(binds).Elem().Index(count).Interface()
		ptrEle := reflect.New(reflect.TypeOf(ele))
		ptrEle.Elem().Set(reflect.ValueOf(ele))
		if err := gorm.DB.Create(ele).Error; err != nil {
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func delete(name string, gorm *GORM, c *gin.Context) {
	var form interface{}
	binds := gorm.Vars(name)
	bind := gorm.Var(name)
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	buf = buf[0:num]
	err := json.Unmarshal(buf, &form)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if reflect.TypeOf(form).Kind() == reflect.Map {
		form = []interface{}{form}
	}

	// map interface{} type to model type
	slicev := reflect.ValueOf(binds).Elem()
	slicev = slicev.Slice(0, slicev.Cap())
	for _, f := range form.([]interface{}) {
		mf, err := json.Marshal(f)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		err = json.Unmarshal(mf, bind)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		slicev = reflect.Append(slicev, reflect.ValueOf(bind).Elem())
		slicev = slicev.Slice(0, slicev.Cap())
	}
	reflect.ValueOf(binds).Elem().Set(slicev)

	// delete model
	count := reflect.ValueOf(binds).Elem().Len()
	for count > 0 {
		count = count - 1
		ele := reflect.ValueOf(binds).Elem().Index(count).Interface()
		ptrEle := reflect.New(reflect.TypeOf(ele))
		ptrEle.Elem().Set(reflect.ValueOf(ele))
		if err := gorm.DB.Delete(ele).Error; err != nil {
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func update(name string, gorm *GORM, c *gin.Context) {
	var form interface{}
	binds := gorm.Vars(name)
	bind := gorm.Var(name)
	buf := make([]byte, 1024)
	num, _ := c.Request.Body.Read(buf)
	buf = buf[0:num]
	err := json.Unmarshal(buf, &form)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if reflect.TypeOf(form).Kind() == reflect.Map {
		form = []interface{}{form}
	}

	// map interface{} type to model type
	slicev := reflect.ValueOf(binds).Elem()
	slicev = slicev.Slice(0, slicev.Cap())
	for _, f := range form.([]interface{}) {
		mf, err := json.Marshal(f)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		err = json.Unmarshal(mf, bind)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		slicev = reflect.Append(slicev, reflect.ValueOf(bind).Elem())
		slicev = slicev.Slice(0, slicev.Cap())
	}
	reflect.ValueOf(binds).Elem().Set(slicev)

	// update model
	for _, f := range form.([]interface{}) {
		if err := gorm.DB.Model(bind).Updates(f).Error; err != nil {
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
