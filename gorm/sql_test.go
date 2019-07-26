// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gormext

import "testing"

func Test_map2sql(t *testing.T) {
	type args struct {
		value map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		struct {
			name    string
			args    args
			want    string
			wantErr bool
		}{
			name: "map2sql",
			args: struct {
				value map[string]interface{}
			}{
				value: map[string]interface{}{
					"Age": map[string]interface{}{"$gte": 30},
					"$or": []map[string]interface{}{
						map[string]interface{}{"Weight": map[string]interface{}{"$lte": 200}},
						map[string]interface{}{
							"Height": 230,
							"$or": []map[string]interface{}{
								map[string]interface{}{"Weight2": map[string]interface{}{"$lte": 200}},
								map[string]interface{}{"Age2": 88},
							},
						},
					},
				},
			},
			want: "13",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := map2sql(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("map2sql() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("map2sql() = %v, want %v", got, tt.want)
			}
		})
	}
}
