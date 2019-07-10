// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gormext

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/thoas/go-funk"

	addition "github.com/2637309949/bulrush-addition"
)

// Query defined query std
type Query struct {
	Cond    string `form:"cond" json:"cond" xml:"cond"`
	CondMap map[string]interface{}
	Select  string `form:"select" json:"select" xml:"select"`
	Preload string `form:"preload" json:"preload" xml:"preload"`
	Order   string `form:"order" json:"order" xml:"order"`
	Page    int    `form:"page" json:"page" xml:"page"`
	Size    int    `form:"size" json:"size" xml:"size"`
	Range   string `form:"range" json:"range" xml:"range"`
}

// BuildCond defined select sql
func (q *Query) BuildCond() error {
	var cond map[string]interface{}
	if q.Cond == "" {
		q.Cond = "%7B%7D"
	}
	unescapeCond, err := url.QueryUnescape(q.Cond)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(unescapeCond), &cond)
	if err != nil {
		return err
	}
	q.CondMap = cond
	return nil
}

// FlatWhere defined flat where to sql
func (q *Query) FlatWhere() (string, error) {
	var cloneCond map[string]interface{}
	err := addition.CopyMap(q.CondMap, &cloneCond)
	flatWhere, err := flatAndToOr(cloneCond)
	if err != nil {
		return "", err
	}
	sql, err := shuttle("", flatWhere)
	return sql, err
}

// BuildOrder defined order sql
func (q *Query) BuildOrder() string {
	var ordersWithDirect []string
	orders := strings.Split(q.Order, ",")
	for _, item := range orders {
		if strings.HasPrefix(item, "-") {
			ordersWithDirect = append(ordersWithDirect, fmt.Sprintf("%s %s", strings.Replace(item, "-", "", 1), "desc"))
		} else {
			ordersWithDirect = append(ordersWithDirect, strings.Replace(item, "+", "", 1))
		}
	}
	return strings.Join(ordersWithDirect, ",")
}

// BuildRelated defined related sql for preLoad
func (q *Query) BuildRelated() []string {
	return strings.Split(q.Preload, ",")
}

// BuildSelect dfined build select fields
func (q *Query) BuildSelect(list interface{}) ([]map[string]interface{}, error) {
	var jArr []map[string]interface{}
	jByte, err := json.Marshal(list)
	err = json.Unmarshal(jByte, &jArr)
	for _, jMap := range jArr {
		for k := range jMap {
			sels := strings.Split(q.Select, ",")
			_, ok := funk.FindString(sels, func(s string) bool {
				return s == k
			})
			if !ok {
				delete(jMap, k)
			}
		}
	}
	return jArr, err
}
