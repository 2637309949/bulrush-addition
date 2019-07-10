// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gormext

import (
	"reflect"
	"regexp"

	"github.com/gin-gonic/gin"
)

func combineHF(handler gin.HandlerFunc, handlers []gin.HandlerFunc) []gin.HandlerFunc {
	h := append(handlers, handler)
	return h
}

func findStringSubmatch(matcher string, s string) []string {
	var rgx = regexp.MustCompile(matcher)
	rs := rgx.FindStringSubmatch(s)
	if rs != nil {
		return rs[1:]
	}
	return []string{}
}

func findFieldStruct(vType reflect.Type, name string) *reflect.StructField {
	if vType.Kind() == reflect.Ptr {
		vType = vType.Elem()
	}
	if vType.Kind() == reflect.Struct {
		numField := vType.NumField()
		if numField > 0 {
			field, ok := vType.FieldByName(name)
			if ok {
				return &field
			}
			field = vType.Field(0)
			return findFieldStruct(field.Type, name)
		}
	}
	return nil
}

func createStruct(sfs []reflect.StructField) interface{} {
	return reflect.New(reflect.StructOf(sfs)).Interface()
}
