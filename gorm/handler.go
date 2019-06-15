// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gorm

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type puFormat struct {
	Cond map[string]interface{} `bson:"cond" form:"cond" json:"cond" xml:"cond"`
	Muti bool                   `bson:"muti" form:"muti" json:"muti" xml:"muti"`
	Doc  interface{}            `bson:"doc" form:"doc" json:"doc" xml:"doc" `
}

func one(name string, gorm *GORM, c *gin.Context) {
}

func list(name string, gorm *GORM, c *gin.Context) {
	var match map[string]interface{}
	list := gorm.Vars(name)
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
	fmt.Println(reflect.TypeOf(list))
	var totalrecords int
	gorm.DB.Offset((page - 1) * size).Limit(size).Find(&list)
	gorm.DB.Count(&totalrecords)
	totalpages := math.Ceil(float64(totalrecords) / float64(size))
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

func create(name string, gorm *GORM, c *gin.Context) {
}

func delete(name string, gorm *GORM, c *gin.Context) {
}

func update(name string, gorm *GORM, c *gin.Context) {

}
