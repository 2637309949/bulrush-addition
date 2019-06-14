// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgo

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type (
	// API type defined
	api struct {
		mgo *Mongo
	}
)

// One hook auto generate api
func (api *api) One(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	handler := func(c *gin.Context) {
		one(name, api.mgo, c)
	}
	h := createHooks(api.mgo, handler)
	handlers = append(handlers, h.R)
	r.GET(fmt.Sprintf("/%s/:id", name), handlers...)
	return h
}

// List hook auto generate api
func (api *api) List(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	handler := func(c *gin.Context) {
		list(name, api.mgo, c)
	}
	h := createHooks(api.mgo, handler)
	handlers = append(handlers, h.R)
	r.GET(fmt.Sprintf("/%s", name), handlers...)
	return h
}

// Create hook auto generate api
func (api *api) Create(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	handler := func(c *gin.Context) {
		create(name, api.mgo, c)
	}
	h := createHooks(api.mgo, handler)
	handlers = append(handlers, h.R)
	r.POST(fmt.Sprintf("/%s", name), handlers...)
	return h
}

// Update hook auto generate api
func (api *api) Update(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	handler := func(c *gin.Context) {
		update(name, api.mgo, c)
	}
	h := createHooks(api.mgo, handler)
	handlers = append(handlers, h.R)
	r.PUT(fmt.Sprintf("/%s", name), handlers...)
	return h
}

// Delete hook auto generate api
func (api *api) Delete(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	handler := func(c *gin.Context) {
		delete(name, api.mgo, c)
	}
	h := createHooks(api.mgo, handler)
	handlers = append(handlers, h.R)
	r.DELETE(fmt.Sprintf("/%s", name), handlers...)
	return h
}

// ALL hook auto generate api
func (api *api) ALL(r *gin.RouterGroup, name string) {
	r.GET(fmt.Sprintf("/%s", name), func(c *gin.Context) {
		list(name, api.mgo, c)
	})
	r.GET(fmt.Sprintf("/%s/:id", name), func(c *gin.Context) {
		one(name, api.mgo, c)
	})
	r.POST(fmt.Sprintf("/%s", name), func(c *gin.Context) {
		create(name, api.mgo, c)
	})
	r.PUT(fmt.Sprintf("/%s", name), func(c *gin.Context) {
		update(name, api.mgo, c)
	})
	r.DELETE(fmt.Sprintf("/%s", name), func(c *gin.Context) {
		delete(name, api.mgo, c)
	})
}
