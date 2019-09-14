// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gormext

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"reflect"
	"strconv"
	"time"

	addition "github.com/2637309949/bulrush-addition"
	"github.com/2637309949/bulrush-utils/funcs"
	"github.com/2637309949/bulrush-utils/regex"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/thoas/go-funk"
	"gopkg.in/go-playground/validator.v9"
)

type form struct {
	Docs     []map[string]interface{} `form:"docs" json:"docs" xml:"docs" binding:"required"`
	Category interface{}              `form:"category" json:"category" xml:"category" binding:"required"`
}

func one(name string, c *gin.Context, ext *GORM, opts *Opts) {
	ret, err := funcs.Chain(func(ret interface{}) (interface{}, error) {
		key := regex.FindStringSubmatch(":(.*)$", opts.RoutePrefixs.One(name))[0]
		id, err := strconv.Atoi(c.Param(key))
		q := NewQuery()
		if err != nil {
			return nil, err
		}
		if err := c.BindQuery(&q.Query); err != nil {
			return nil, err
		}
		q.Cond = opts.RouteHooks.One.Cond(map[string]interface{}{"deletedAt": map[string]interface{}{"$exists": false}, "id": id}, c, struct{ Name string }{Name: name})
		if err := q.Build(q.Cond); err != nil {
			return nil, err
		}
		return q, nil
	}, func(ret interface{}) (interface{}, error) {
		one := ext.Var(name)
		q := ret.(*Query)
		err := ext.DB.
			Scopes(func(db *gorm.DB) *gorm.DB {
				if q.Select != "" {
					return db.Select(q.Select)
				}
				return db
			}).
			Scopes(func(db *gorm.DB) *gorm.DB {
				funk.ForEach(q.Preload, func(pre string) {
					if pre != "" {
						db = db.Preload(pre)
					}
				})
				return db
			}).
			Scopes(func(db *gorm.DB) *gorm.DB {
				return db.Where(q.SQL)
			}).
			First(one).Error
		return one, err
	})
	if err != nil {
		addition.RushLogger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ret)
}

func list(name string, c *gin.Context, ext *GORM, opts *Opts) {
	ret, err := funcs.Chain(func(ret interface{}) (interface{}, error) {
		q := NewQuery()
		if err := c.BindQuery(&q.Query); err != nil {
			return nil, err
		}
		q.Cond = opts.RouteHooks.List.Cond(map[string]interface{}{"deletedAt": map[string]interface{}{"$exists": false}}, c, struct{ Name string }{Name: name})
		if err := q.Build(q.Cond); err != nil {
			return nil, err
		}
		return q, nil
	}, func(ret interface{}) (interface{}, error) {
		totalrecords := 0
		one := ext.Var(name)
		list := ext.Vars(name)
		q := ret.(*Query)
		db := ext.DB
		db = db.
			Scopes(func(db *gorm.DB) *gorm.DB {
				if q.Range != "ALL" {
					db = db.Offset((q.Page - 1) * q.Size).Limit(q.Size)
				}
				return db
			}).
			Scopes(func(db *gorm.DB) *gorm.DB {
				if q.Select != "" {
					db = db.Select(q.Select)
				}
				return db
			}).
			Scopes(func(db *gorm.DB) *gorm.DB {
				if q.Order != "" {
					db = db.Order(q.Order)
				}
				return db
			}).
			Scopes(func(db *gorm.DB) *gorm.DB {
				funk.ForEach(q.Preload, func(pre string) {
					if pre != "" {
						db = db.Preload(pre)
					}
				})
				return db
			}).
			Scopes(func(db *gorm.DB) *gorm.DB {
				if q.SQL != "" {
					db = db.Where(q.SQL)
				}
				return db
			}).Find(list)

		if db.Error != nil {
			return nil, db.Error
		}
		if err := ext.DB.Model(one).Where(q.SQL).Count(&totalrecords).Error; err != nil {
			if err != nil {
				return nil, err
			}
		}

		if q.Range != "ALL" {
			totalpages := math.Ceil(float64(totalrecords) / float64(q.Size))
			return gin.H{
				"page":         q.Query.Page,
				"size":         q.Query.Size,
				"totalpages":   totalpages,
				"range":        q.Query.Range,
				"totalrecords": totalrecords,
				"cond":         q.Cond,
				"select":       q.Query.Select,
				"preload":      q.Query.Preload,
				"list":         list,
			}, nil
		}
		return gin.H{
			"range":        q.Query.Range,
			"totalrecords": totalrecords,
			"cond":         q.Cond,
			"select":       q.Query.Select,
			"preload":      q.Query.Preload,
			"list":         list,
		}, nil
	})
	if err != nil {
		addition.RushLogger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ret)

}

func create(name string, c *gin.Context, ext *GORM, opts *Opts) {
	ret, err := funcs.Chain(func(ret interface{}) (interface{}, error) {
		var form form
		list := ext.Vars(name)
		if err := c.ShouldBindJSON(&form); err != nil {
			addition.RushLogger.Error(err.Error())
			return nil, err
		}
		form = opts.RouteHooks.Create.Form(form)
		docsByte, err := json.Marshal(form.Docs)
		if err != nil {
			addition.RushLogger.Error(err.Error())
			return nil, err
		}
		if err := json.Unmarshal(docsByte, list); err != nil {
			addition.RushLogger.Error(err.Error())
			return nil, err
		}
		return list, nil
	}, func(list interface{}) (interface{}, error) {
		// save docs
		rowsAffected := make([]int64, 0)
		validate := validator.New()
		listValue := reflect.ValueOf(list).Elem()
		count := listValue.Len()
		tx := ext.DB.Begin()
		// Business and security considerations
		// 	Only save reference if exists, no create ref and no update ref
		tx = tx.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false)
		for index := 0; index < count; index++ {
			item := listValue.Index(index).Interface()
			if err := validate.Struct(item); err != nil {
				tx.Rollback()
				addition.RushLogger.Error(err.Error())
				return nil, err
			}
			newValue := reflect.New(reflect.TypeOf(item))
			newValue.Elem().Set(reflect.ValueOf(item))
			ret := tx.Create(newValue.Interface())
			if err := ret.Error; err != nil {
				tx.Rollback()
				addition.RushLogger.Error(err.Error())
				return nil, err
			}
			rowsAffected = append(rowsAffected, ret.RowsAffected)
		}
		if err := tx.Commit().Error; err != nil {
			addition.RushLogger.Error(err.Error())
			return nil, err
		}
		return rowsAffected, nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ret)
}

func remove(name string, c *gin.Context, ext *GORM, opts *Opts) {
	var form form
	bind := ext.Var(name)
	rowsAffected := make([]int64, 0)
	q := NewQuery()
	if err := c.ShouldBindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	form = opts.RouteHooks.Delete.Form(form)
	// Business and security considerations
	// 	Only save reference if exists, no create ref and no update ref
	tx := ext.DB.Begin()
	tx = tx.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false)
	for _, item := range form.Docs {
		id, ok := item["ID"]
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "no id found",
			})
			return
		}
		q.Cond = opts.RouteHooks.Delete.Cond(map[string]interface{}{"ID": id}, c, struct{ Name string }{Name: name})
		if err := q.Build(q.Cond); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		ret := tx.Model(bind).Where(q.SQL).Update(map[string]interface{}{"deleted_at": time.Now()})
		if err := ret.Error; err != nil {
			tx.Rollback()
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		rowsAffected = append(rowsAffected, ret.RowsAffected)
	}
	err := tx.Commit().Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, rowsAffected)
}

func update(name string, c *gin.Context, ext *GORM, opts *Opts) {
	ret, err := funcs.Chain(func(ret interface{}) (interface{}, error) {
		var form form
		if err := c.ShouldBindJSON(&form); err != nil {
			addition.RushLogger.Error(err.Error())
			return nil, err
		}
		form = opts.RouteHooks.Create.Form(form)
		return &form, nil
	}, func(ret interface{}) (interface{}, error) {
		rowsAffected := make([]int64, 0)
		form := ret.(*form)
		q := NewQuery()
		// Business and security considerations
		// 	Only save reference if exists, no create ref and no update ref
		tx := ext.DB.Begin()
		tx = tx.Set("gorm:association_autoupdate", false).Set("gorm:association_autocreate", false)
		for _, doc := range form.Docs {
			one := ext.Var(name)
			id, ok := doc["ID"]
			if !ok {
				addition.RushLogger.Error("no id found")
				return nil, errors.New("no id found")
			}
			jsonByte, err := json.Marshal(doc)
			if err != nil {
				addition.RushLogger.Error(err.Error())
				return nil, err
			}
			if err := json.Unmarshal(jsonByte, one); err != nil {
				addition.RushLogger.Error(err.Error())
				return nil, err
			}
			q.Cond = opts.RouteHooks.Update.Cond(map[string]interface{}{"ID": id}, c, struct{ Name string }{Name: name})
			if err := q.Build(q.Cond); err != nil {
				addition.RushLogger.Error(err.Error())
				return nil, err
			}
			ret := tx.Model(one).Where(q.SQL).Update(one)
			if err := ret.Error; err != nil {
				tx.Rollback()
				addition.RushLogger.Error(err.Error())
				return nil, err
			}
			rowsAffected = append(rowsAffected, ret.RowsAffected)
		}
		err := tx.Commit().Error
		return rowsAffected, err
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ret)
}
