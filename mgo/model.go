// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgoext

import (
	"time"

	"github.com/globalsign/mgo/bson"
)

// Model common fields
type Model struct {
	ID      bson.ObjectId `bson:"_id,omitempty" br:"comment:'模型ID'"`
	Created *time.Time    `bson:"_created" br:"comment:'创建时间'"`
	Updated *time.Time    `bson:"_updated" br:"comment:'修改时间'"`
	Deleted *time.Time    `bson:"_deleted" br:"comment:'删除时间'"`
}
