// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gormext

import "github.com/gin-gonic/gin"

func combineHF(handler gin.HandlerFunc, handlers []gin.HandlerFunc) []gin.HandlerFunc {
	h := append(handlers, handler)
	return h
}