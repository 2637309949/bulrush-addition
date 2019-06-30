// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgoext

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/gin-gonic/gin"
)

func findStringSubmatch(matcher string, s string) []string {
	var rgx = regexp.MustCompile(matcher)
	rs := rgx.FindStringSubmatch(s)
	if rs != nil {
		return rs[1:]
	}
	return []string{}
}

func combineHF(handler gin.HandlerFunc, handlers []gin.HandlerFunc) []gin.HandlerFunc {
	h := append(handlers, handler)
	return h
}

func fieldTag(target interface{}, field string, tagname string) string {
	elementType := reflect.TypeOf(target)
	if elementType.Kind() == reflect.Ptr {
		elementType = elementType.Elem()
	}
	fmt.Println("elementType = ", elementType)

	tag := ""
	for i := 0; i < elementType.NumField(); i++ {
		bsonFieldTag := elementType.Field(i).Tag.Get("bson")
		if bsonFieldTag == field {
			tag = elementType.Field(i).Tag.Get("ref")
		}
	}
	if tag == "" {
		if elementType.Field(0).Type.Kind() == reflect.Struct {
			tag = fieldTag(reflect.New(elementType.Field(0).Type).Interface(), field, tagname)
		}
	}
	return tag
}
