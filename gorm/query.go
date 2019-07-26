// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gormext

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	addition "github.com/2637309949/bulrush-addition"
)

// Query defined query std
type Query struct {
	Query struct {
		Cond    string `form:"cond" json:"cond" xml:"cond"`
		Select  string `form:"select" json:"select" xml:"select"`
		Preload string `form:"preload" json:"preload" xml:"preload"`
		Order   string `form:"order" json:"order" xml:"order"`
		Page    int    `form:"page" json:"page" xml:"page"`
		Size    int    `form:"size" json:"size" xml:"size"`
		Range   string `form:"range" json:"range" xml:"range"`
	}
	Cond    map[string]interface{}
	SQL     string
	Select  string
	Preload []string
	Order   string
	Page    int
	Size    int
	Range   string
}

// NewQuery defined new query
func NewQuery() *Query {
	return &Query{
		Query: struct {
			Cond    string `form:"cond" json:"cond" xml:"cond"`
			Select  string `form:"select" json:"select" xml:"select"`
			Preload string `form:"preload" json:"preload" xml:"preload"`
			Order   string `form:"order" json:"order" xml:"order"`
			Page    int    `form:"page" json:"page" xml:"page"`
			Size    int    `form:"size" json:"size" xml:"size"`
			Range   string `form:"range" json:"range" xml:"range"`
		}{
			Cond:  "%7B%7D",
			Size:  20,
			Page:  1,
			Range: "PAGE",
			Order: "-created_at",
		},
		Cond:  map[string]interface{}{},
		Page:  1,
		Size:  20,
		Range: "PAGE",
		Order: "-created_at",
	}
}

// Build defined all build
func (q *Query) Build(pre map[string]interface{}) error {
	q.Page = q.Query.Page
	q.Size = q.Query.Size
	q.Range = q.Query.Range
	err := q.BuildCond(pre)
	q.BuildOrder()
	q.BuildRelated()
	return err
}

// BuildCond defined select sql
func (q *Query) BuildCond(cond map[string]interface{}) error {
	var clone map[string]interface{}
	q.Cond = cond
	unescapeCond, err := url.QueryUnescape(q.Query.Cond)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(unescapeCond), &q.Cond)
	if err != nil {
		return err
	}
	if err = addition.CopyMap(q.Cond, &clone); err != nil {
		return err
	}
	sql, err := map2sql(clone)
	if err != nil {
		return err
	}
	q.SQL = sql
	return nil
}

// BuildOrder defined order sql
func (q *Query) BuildOrder() {
	var ordersWithDirect []string
	orders := strings.Split(q.Query.Order, ",")
	for _, item := range orders {
		if strings.HasPrefix(item, "-") {
			ordersWithDirect = append(ordersWithDirect, fmt.Sprintf("%s %s", strings.Replace(item, "-", "", 1), "desc"))
		} else {
			ordersWithDirect = append(ordersWithDirect, strings.Replace(item, "+", "", 1))
		}
	}
	q.Order = strings.Join(ordersWithDirect, ",")
}

// BuildRelated defined related sql for preLoad
func (q *Query) BuildRelated() {
	q.Preload = strings.Split(q.Query.Preload, ",")
}
