/**
 * @author [Double]
 * @email [2637309949@qq.com.com]
 * @create date 2019-01-12 22:46:31
 * @modify date 2019-01-12 22:46:31
 * @desc [bulrush reflect]
 */

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
