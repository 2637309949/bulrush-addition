// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gormext

import "time"

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models
type Model struct {
	ID        uint       `gorm:"comment:主键ID;primary_key"`
	CreatedAt *time.Time `gorm:"comment:创建时间;"`
	UpdatedAt *time.Time `gorm:"comment:更新时间;"`
	DeletedAt *time.Time `gorm:"comment:删除时间;"`
}
