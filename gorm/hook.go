// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gormext

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hook for API
type Hook struct {
	auth    func(c *gin.Context) bool
	handler func(c *gin.Context)
	gorm    *GORM
	pre     func(c *gin.Context)
	post    func(c *gin.Context)
	r       func(c *gin.Context)
}

// FailureHandler handlerss
var FailureHandler = func(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"message": "Internal Server Error",
		"stack":   "Access Denied, you don't have permission",
	})
}

// route defined gin middle
func (h *Hook) route() func(c *gin.Context) {
	return func(c *gin.Context) {
		if h.pre != nil {
			h.pre(c)
		}
		if h.handler != nil {
			if h.auth != nil && h.auth(c) {
				h.handler(c)
			} else if h.auth == nil {
				h.handler(c)
			} else {
				FailureHandler(c)
			}
		}
		if h.post != nil {
			h.post(c)
		}
	}
}

// Pre defined pre handler
func (h *Hook) Pre(pre func(*gin.Context)) *Hook {
	if pre != nil {
		h.pre = pre
	}
	return h
}

// Post defined pre handler
func (h *Hook) Post(post func(*gin.Context)) *Hook {
	if post != nil {
		h.post = post
	}
	return h
}

// Auth defined pre handler
func (h *Hook) Auth(auth func(*gin.Context) bool) *Hook {
	if auth != nil {
		h.auth = auth
	}
	return h
}

// createHooks alloc new hook
func createHooks(gorm *GORM, handler func(c *gin.Context)) *Hook {
	h := &Hook{
		gorm:    gorm,
		handler: handler,
	}
	h.r = h.route()
	return h
}
