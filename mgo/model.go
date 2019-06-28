// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgoext

// Model common fields
type Model struct {
	Created  int `bson:"_created" form:"_created" json:"_created" xml:"_created"`
	Modified int `bson:"_modified" form:"_modified" json:"_modified" xml:"_modified"`
	Deleted  int `bson:"_deleted" form:"_deleted" json:"_deleted" xml:"_deleted"`
}
