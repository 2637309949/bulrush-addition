// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package addition

import (
	"reflect"
)

// CreateSlice reflect and create
func CreateSlice(target interface{}) interface{} {
	tType := reflect.ValueOf(target).Type()
	if tType.Kind() == reflect.Ptr {
		tType = tType.Elem()
	}
	tSlice := reflect.MakeSlice(reflect.SliceOf(tType), 0, 0).Interface()
	return tSlice
}

// CreateObject reflect and create
func CreateObject(target interface{}) interface{} {
	tType := reflect.ValueOf(target).Type()
	if tType.Kind() == reflect.Ptr {
		tType = tType.Elem()
	}
	tObject := reflect.New(tType).Interface()
	return tObject
}
