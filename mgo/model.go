// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgoext

// Model common fields
type Model struct {
	Created  int `bson:"_created,comment:创建时间,"`
	Modified int `bson:"_modified,comment:修改时间,"`
	Deleted  int `bson:"_deleted,comment:删除时间,"`
}
