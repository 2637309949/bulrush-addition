// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gormext

import "time"

// Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models
type Model struct {
	ID        uint       `gorm:"primary_key" form:"id" json:"id" xml:"id"`
	CreatedAt time.Time  `form:"createdAt" json:"createdAt" xml:"createdAt"`
	UpdatedAt time.Time  `form:"updatedAt" json:"updatedAt" xml:"updatedAt"`
	DeletedAt *time.Time `sql:"index" form:"deletedAt" json:"deletedAt" xml:"deletedAt"`
}
