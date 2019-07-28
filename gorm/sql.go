// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gormext

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/thoas/go-funk"
)

func formatNull(key string, instruct string, value interface{}) string {
	return fmt.Sprintf("%s %s null", key, instruct)
}

func formatString(key string, instruct string, value interface{}) string {
	return fmt.Sprintf("%s %s '%v'", key, instruct, value)
}

func formatNumber(key string, instruct string, value interface{}) string {
	return fmt.Sprintf("%s %s %v", key, instruct, value)
}

func formatArray(key string, instruct string, value interface{}) string {
	items := funk.Map(value, func(item interface{}) string {
		if reflect.TypeOf(item).Kind() == reflect.String {
			return fmt.Sprintf("'%s'", item)
		}
		return fmt.Sprintf("%v", item)
	}).([]string)
	return fmt.Sprintf("%s %s (%v)", key, instruct, strings.Join(items, ","))
}

// check whether is least or not
func isLeastDirect(key string, value interface{}) bool {
	// Security filtering for key
	if strings.Contains(key, " ") {
		return false
	}
	// isLeastDirect
	if !strings.HasPrefix(key, "$") {
		if toMap, ok := value.(map[string]interface{}); ok {
			for key := range toMap {
				if !strings.HasPrefix(key, "$") {
					return false
				}
			}
		}
		return true
	}
	return false
}

// sql direct to sql string
func direct2Sql(key string, instruct string, value interface{}) string {
	switch true {
	case isNull(value):
		return formatNull(key, instruct, value)
	case isNumber(value):
		return formatNumber(key, instruct, value)
	case isString(value):
		return formatString(key, instruct, value)
	case isSlice(value):
		return formatArray(key, instruct, value)
	}
	andJoin := []string{}
	if vmap, ok := value.(map[string]interface{}); ok {
		for k, v := range vmap {
			if k == "$eq" {
				subItem := direct2Sql(key, "=", v)
				andJoin = append(andJoin, subItem)
			}
			if k == "$ne" {
				subItem := direct2Sql(key, "<>", v)
				andJoin = append(andJoin, subItem)
			}
			if k == "$gte" {
				subItem := direct2Sql(key, ">=", v)
				andJoin = append(andJoin, subItem)
			}
			if k == "$gt" {
				subItem := direct2Sql(key, ">", v)
				andJoin = append(andJoin, subItem)
			}
			if k == "$lte" {
				subItem := direct2Sql(key, "<=", v)
				andJoin = append(andJoin, subItem)
			}
			if k == "$lt" {
				subItem := direct2Sql(key, "<", v)
				andJoin = append(andJoin, subItem)
			}
			if k == "$in" {
				subItem := direct2Sql(key, "in", v)
				andJoin = append(andJoin, subItem)
			}
			if k == "$regex" {
				subItem := direct2Sql(key, "regexp", v)
				andJoin = append(andJoin, subItem)
			}
			if k == "$like" {
				subItem := direct2Sql(key, "like", v)
				andJoin = append(andJoin, subItem)
			}
			if k == "$exists" {
				vb, ok := v.(bool)
				if ok && vb {
					andJoin = append(andJoin, fmt.Sprintf("%s is %s null", key, "no"))
				} else {
					andJoin = append(andJoin, fmt.Sprintf("%s is %s null", key, ""))
				}
			}
		}
		return strings.Join(andJoin, " and ")
	}
	// should be panic
	return "stop sql"
}

// flaten $or direct to $and
func flatAndToOr(where map[string]interface{}) (map[string]interface{}, error) {
	var err error
	if or, ok := where["$or"]; ok && reflect.Slice == reflect.TypeOf(or).Kind() {
		newMap := map[string]interface{}{}
		newMap["$or"] = []map[string]interface{}{}
		orArr := toArrayInterface(or)
		for _, item := range orArr {
			subItem, ok := item.(map[string]interface{})
			_, ok = subItem["$or"]
			for wk, wv := range where {
				if wk != "$or" {
					subItem[wk] = wv
				}
			}
			if ok {
				subItem, err = flatAndToOr(subItem)
				if err != nil {
					return map[string]interface{}{}, errors.New("orMap2 error")
				}
			}
			newMap["$or"] = append(newMap["$or"].([]map[string]interface{}), subItem)
		}
		return newMap, nil
	}
	return where, nil
}

// flaten map to sql
func shuttle(key string, value interface{}) (string, error) {
	if key == "" {
		andJoin := []string{}
		if mapv, ok := value.(map[string]interface{}); ok {
			for k, v := range mapv {
				sub, err := shuttle(k, v)
				if err != nil {
					return "", errors.New("shuttle2 error")
				}
				if strings.Contains(sub, "or") {
					andJoin = append(andJoin, fmt.Sprintf("(%s)", sub))
				} else {
					andJoin = append(andJoin, sub)
				}
			}
			return strings.Join(andJoin, " and "), nil
		}
		return "", errors.New("shuttle1 error")
	}
	if key == "$or" {
		orJoin := []string{}
		if vArr, ok := value.([]map[string]interface{}); ok {
			for _, v := range vArr {
				sub, err := shuttle("", v)
				if err != nil {
					return "", errors.New("shuttle4 error")
				}
				if strings.Contains(sub, "and") {
					orJoin = append(orJoin, fmt.Sprintf("(%s)", sub))
				} else {
					orJoin = append(orJoin, sub)
				}
			}
			return strings.Join(orJoin, " or "), nil
		}
		return "", errors.New("shuttle3 error")
	}
	if isLeastDirect(key, value) {
		return direct2Sql(columnNamer(key), "=", value), nil
	}
	return "", errors.New("shuttle6 error")
}

func map2sql(value map[string]interface{}) (string, error) {
	flat, err := flatAndToOr(value)
	if err != nil {
		return "", err
	}
	sql, err := shuttle("", flat)
	if err != nil {
		return "", err
	}
	return sql, nil
}
