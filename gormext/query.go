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
	Where    string `form:"where" json:"where" xml:"where"`
	WhereMap map[string]interface{}
	Select   string `form:"select" json:"select" xml:"select"`
	Related  string `form:"related" json:"related" xml:"related"`
	Order    string `form:"order" json:"order" xml:"order"`
	Page     int    `form:"page" json:"page" xml:"page"`
	Size     int    `form:"size" json:"size" xml:"size"`
	Range    string `form:"range" json:"range" xml:"range"`
}

// BuildWhere defined select sql
func (q *Query) BuildWhere() (string, error) {
	var where map[string]interface{}
	var cloneWhere map[string]interface{}
	if q.Where == "" {
		q.Where = "%7B%7D"
	}
	unescapeWhere, err := url.QueryUnescape(q.Where)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal([]byte(unescapeWhere), &where)
	if err != nil {
		return "", err
	}
	err = addition.CopyMap(where, &cloneWhere)
	if err != nil {
		return "", err
	}
	q.WhereMap = cloneWhere
	flatWhere, err := flatAndToOr(cloneWhere)
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
	return strings.Split(q.Related, ",")
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
