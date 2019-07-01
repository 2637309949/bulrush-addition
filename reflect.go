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
	tType = indirectType(tType)
	return reflect.New(reflect.SliceOf(tType)).Interface()
}

// CreateObject reflect and create
func CreateObject(target interface{}) interface{} {
	tType := reflect.ValueOf(target).Type()
	tType = indirectType(tType)
	return reflect.New(tType).Interface()
}

// indirect from ptr
func indirectValue(reflectValue reflect.Value) reflect.Value {
	if reflectValue.Kind() == reflect.Ptr && reflectValue.Elem().Kind() == reflect.Interface {
		reflectValue = reflectValue.Elem().Elem()
	}
	for reflectValue.Kind() == reflect.Slice || reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}
	return reflectValue
}

// indirect from ptr
func indirectType(reflectType reflect.Type) reflect.Type {
	if reflectType.Kind() == reflect.Ptr && reflectType.Elem().Kind() == reflect.Interface {
		reflectType = reflectType.Elem().Elem()
	}
	for reflectType.Kind() == reflect.Slice || reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}
	return reflectType
}
