// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package xormext

import (
	"errors"
	"math"
	"net/http"
	"strconv"

	addition "github.com/2637309949/bulrush-addition"
	"github.com/2637309949/bulrush-utils/funcs"
	"github.com/2637309949/bulrush-utils/regex"
	"github.com/gin-gonic/gin"
)

type form struct {
	Docs     []map[string]interface{} `form:"docs" json:"docs" xml:"docs" binding:"required"`
	Category interface{}              `form:"category" json:"category" xml:"category" binding:"required"`
}

func one(name string, c *gin.Context, ext *XORM, opts *Opts) {
	ret, err := funcs.Chain(
		func(ret interface{}) (interface{}, error) {
			key := regex.FindStringSubmatch(":(.*)$", opts.RoutePrefixs.One(name))[0]
			id, err := strconv.Atoi(c.Param(key))
			q := NewQuery()
			if err != nil {
				return nil, err
			}
			if err := c.BindQuery(&q.Query); err != nil {
				return nil, err
			}
			q.Cond = opts.RouteHooks.One.Cond(map[string]interface{}{"deletedAt": map[string]interface{}{"$exists": false}, "i_d": id}, c, struct{ Name string }{Name: name})
			if err := q.Build(q.Cond); err != nil {
				return nil, err
			}
			return q, nil
		},
		func(ret interface{}) (interface{}, error) {
			one := ext.Var(name)
			q := ret.(*Query)
			if ok, err := ext.DB.Where(q.SQL).Get(one); err != nil || !ok {
				if err != nil {
					return nil, err
				}
				return nil, errors.New("not found")
			}
			return one, nil
		},
	)
	if err != nil {
		addition.RushLogger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ret)
}

func list(name string, c *gin.Context, ext *XORM, opts *Opts) {
	ret, err := funcs.Chain(
		func(ret interface{}) (interface{}, error) {
			q := NewQuery()
			if err := c.BindQuery(&q.Query); err != nil {
				return nil, err
			}
			q.Cond = opts.RouteHooks.List.Cond(map[string]interface{}{"deletedAt": map[string]interface{}{"$exists": false}}, c, struct{ Name string }{Name: name})
			if err := q.Build(q.Cond); err != nil {
				return nil, err
			}
			return q, nil
		},
		func(ret interface{}) (interface{}, error) {
			var err error
			totalrecords := int64(0)
			one := ext.Var(name)
			list := ext.Vars(name)
			q := ret.(*Query)
			if err := ext.DB.Where(q.SQL).Find(list); err != nil {
				return nil, err
			}
			if totalrecords, err = ext.DB.Where(q.SQL).Count(one); err != nil {
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
		},
	)
	if err != nil {
		addition.RushLogger.Error(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ret)
}

func create(name string, c *gin.Context, ext *XORM, opts *Opts) {
}

func remove(name string, c *gin.Context, ext *XORM, opts *Opts) {
}

func update(name string, c *gin.Context, ext *XORM, opts *Opts) {
}
