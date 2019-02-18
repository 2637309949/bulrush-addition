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

// CreateSlice create slice by variable
func CreateSlice(target interface{}) interface{} {
	tagetType := reflect.TypeOf(target)
	if tagetType.Kind() == reflect.Ptr {
		tagetType = tagetType.Elem()
	}
	targetSlice := reflect.MakeSlice(reflect.SliceOf(tagetType), 0, 0).Interface()
	return targetSlice
}

// CreateObject create object by variable
func CreateObject(target interface{}) interface{} {
	tagetType := reflect.TypeOf(target)
	if tagetType.Kind() == reflect.Ptr {
		tagetType = tagetType.Elem()
	}
	targetObject := reflect.New(tagetType).Interface()
	return targetObject
}
