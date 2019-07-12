// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package addition

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
)

// Some get or a default value
func Some(target interface{}, initValue interface{}) interface{} {
	if target != nil && target != "" && target != 0 {
		return target
	}
	return initValue
}

// Find elements
func Find(arrs []interface{}, matcher func(interface{}) bool) interface{} {
	var target interface{}
	for _, item := range arrs {
		match := matcher(item)
		if match {
			target = item
			break
		}
	}
	return target
}

// ToStrArray -
func ToStrArray(t []interface{}) []string {
	s := make([]string, len(t))
	for i, v := range t {
		s[i] = fmt.Sprint(v)
	}
	return s
}

// ToIntArray -
func ToIntArray(t []interface{}) []int {
	s := make([]int, len(t))
	for i, v := range t {
		s[i] = v.(int)
	}
	return s
}

// RandString -
func RandString(n int) string {
	const seeds = "abcdefghijklmnopqrstuvwxyz1234567890"
	bytes := make([]byte, n)
	for i := range bytes {
		bytes[i] = seeds[rand.Intn(len(seeds))]
	}
	return string(bytes)
}

// LeftV -
func LeftV(left interface{}, right interface{}) interface{} {
	return left
}

// RightV -
func RightV(left interface{}, right interface{}) interface{} {
	return right
}

// LeftOkV -
func LeftOkV(left interface{}, right ...bool) interface{} {
	var r = true
	if len(right) != 0 {
		r = right[0]
	} else if left == "" || left == nil || left == 0 {
		r = false
	}
	if r {
		return left
	}
	return nil
}

// LeftSV -
func LeftSV(left interface{}, right error) interface{} {
	if right != nil {
		panic(right)
	}
	return left
}

// CopyMap -
func CopyMap(src map[string]interface{}, dest interface{}) error {
	jsonStr, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonStr, dest)
	if err != nil {
		return err
	}
	return nil
}

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
