// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gorm

import (
	"encoding/json"
	"math"
	"net/http"
	"net/url"
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
	db := gorm.DB.Where(sql).Offset((page - 1) * size).Limit(size)
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
	db.Find(list)
	gorm.DB.Model(one).Where(sql).Count(&totalrecords)
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
}

func create(name string, gorm *GORM, c *gin.Context) {
}

func delete(name string, gorm *GORM, c *gin.Context) {
}

func update(name string, gorm *GORM, c *gin.Context) {

}
