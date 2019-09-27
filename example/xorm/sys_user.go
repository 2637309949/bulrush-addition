// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"time"

	"gopkg.in/guregu/null.v3"
)

// User defined user info
type User struct {
	ID        uint64                 `xorm:"pk comment('模型ID')"`
	CreatorID uint64                 `xorm:"comment('创建人ID')"`
	Creator   *User                  `xorm:"-"`
	CreatedAt *time.Time             `xorm:"comment('创建时间')"`
	UpdatorID uint64                 `xorm:"comment('修改人ID')"`
	Updator   *User                  `xorm:"-"`
	UpdatedAt *time.Time             `xorm:"comment('更新时间')"`
	DeleterID uint64                 `xorm:"comment('删除人ID')"`
	Deleter   *User                  `xorm:"-"`
	DeletedAt *time.Time             `xorm:"comment('删除时间')"`
	Name      string                 `xorm:"comment('名称') unique not null"`
	Password  string                 `xorm:"comment('密码') not null"`
	Salt      string                 `xorm:"comment('盐噪点') not null"`
	Age       uint                   `xorm:"comment('年龄')"`
	Birthday  *time.Time             `xorm:"comment('生日')"`
	Mobile    string                 `xorm:"comment('手机')"`
	Email     null.String            `xorm:"comment('邮箱')"`
	Ignored   *struct{ Name string } `xorm:"-"`
}
