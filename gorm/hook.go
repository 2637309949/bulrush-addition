// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gormext

import (
	"github.com/gin-gonic/gin"
)

type (
	// HandlerAuthFunc defined auth type
	HandlerAuthFunc func(c *gin.Context) bool
	// Hook for API
	Hook struct {
		auth    HandlerAuthFunc
		handler gin.HandlerFunc
		gorm    *GORM
		pre     gin.HandlerFunc
		post    gin.HandlerFunc
		R       gin.HandlerFunc
	}
)

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
			}
		}
		if h.post != nil {
			h.post(c)
		}
	}
}

// Pre defined pre handler
func (h *Hook) Pre(pre func(*gin.Context)) *Hook {
	if h.pre == nil {
		h.pre = pre
	}
	return h
}

// Post defined pre handler
func (h *Hook) Post(post func(*gin.Context)) *Hook {
	if h.post == nil {
		h.post = post
	}
	return h
}

// Auth defined pre handler
func (h *Hook) Auth(auth func(*gin.Context) bool) *Hook {
	if h.auth == nil {
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
	h.R = h.route()
	return h
}
