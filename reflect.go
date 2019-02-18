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

// createSlice create slice by variable
func createSlice(target interface{}) interface{} {
	tagetType := reflect.TypeOf(target)
	if tagetType.Kind() == reflect.Ptr {
		tagetType = tagetType.Elem()
	}
	targetSlice := reflect.MakeSlice(reflect.SliceOf(tagetType), 0, 0).Interface()
	return targetSlice
}

// createObject create object by variable
func createObject(target interface{}) interface{} {
	tagetType := reflect.TypeOf(target)
	if tagetType.Kind() == reflect.Ptr {
		tagetType = tagetType.Elem()
	}
	targetObject := reflect.New(tagetType).Interface()
	return targetObject
}
