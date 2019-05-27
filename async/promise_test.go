/**
 * @author [Double]
 * @email [2637309949@qq.com.com]
 * @create date 2019-01-12 22:46:31
 * @modify date 2019-01-12 22:46:31
 * @desc [bulrush async plugin]
 */

package async

import (
	"reflect"
	"testing"
)

func Test_gPromise(t *testing.T) {
	type args struct {
		funk func(resolve func(interface{}), reject func(interface{}))
	}
	tests := []struct {
		name string
		args args
		want *promise
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gPromise(tt.args.funk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gPromise() = %v, want %v", got, tt.want)
			}
		})
	}
	gPromise(func(res func(interface{}), rej func(interface{})) {
		res(12)
		rej("test")
	})
	gPromise(func(res func(interface{})) {
		res(12)
	})
}
