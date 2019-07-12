// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgoext

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
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

func preloadInfo(target interface{}, preload string) *PreloadInfo {
	elementType := reflect.TypeOf(target)
	if elementType.Kind() == reflect.Ptr {
		elementType = elementType.Elem()
	}
	field, ok := elementType.FieldByName(preload)
	if !ok {
		return nil
	}
	bsonName := strings.Split(field.Tag.Get("bson"), ",")[0]
	refStr := findStringSubmatch(`ref\((.*?)\)`, field.Tag.Get("br"))
	upStr := findStringSubmatch(`up\((.*?)\)`, field.Tag.Get("br"))
	if len(refStr) > 0 {
		refInfo := strings.Split(refStr[0], ",")
		if len(refInfo) >= 2 {
			info := &PreloadInfo{}
			info.BsonName = preload
			info.Coll = refInfo[0]
			info.Local = refInfo[1]
			if bsonName != "" {
				info.BsonName = bsonName
			}
			if len(refInfo) > 2 {
				info.Foreign = refInfo[2]
			} else {
				info.Foreign = "_id"
			}
			if len(upStr) > 0 {
				upArr := strings.Split(upStr[0], ",")
				upArr = funk.Map(upArr, func(x string) string {
					return fmt.Sprintf("%s.%s", info.BsonName, x)
				}).([]string)
				info.Up = strings.Join(upArr, ",")
			}
			if field.Type.Kind() == reflect.Ptr {
				info.IsArray = field.Type.Elem().Kind() == reflect.Slice
			} else {
				info.IsArray = field.Type.Kind() == reflect.Slice
			}
			return info
		}
	}
	return nil
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
