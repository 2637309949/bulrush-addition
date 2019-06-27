// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgoext

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// API type defined
type api struct {
	mgo *Mongo
}

// Feature defined feature api
func (ai *api) Feature(name string) *api {
	feature := &api{mgo: ai.mgo}
	return feature
}

// FeatureWithOpts defined feature api with opts
func (ai *api) FeatureWithOpts() *api {
	feature := &api{mgo: ai.mgo}
	return feature
}

// One hook auto generate api
func (ai *api) One(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	handler := func(c *gin.Context) {
		one(name, ai.mgo, c)
	}
	h := createHooks(ai.mgo, handler)
	handlers = append(handlers, h.R)
	r.GET(fmt.Sprintf("/%s/:id", name), handlers...)
	return h
}

// List hook auto generate api
func (ai *api) List(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	handler := func(c *gin.Context) {
		list(name, ai.mgo, c)
	}
	h := createHooks(ai.mgo, handler)
	handlers = append(handlers, h.R)
	r.GET(fmt.Sprintf("/%s", name), handlers...)
	return h
}

// Create hook auto generate api
func (ai *api) Create(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	handler := func(c *gin.Context) {
		create(name, ai.mgo, c)
	}
	h := createHooks(ai.mgo, handler)
	handlers = append(handlers, h.R)
	r.POST(fmt.Sprintf("/%s", name), handlers...)
	return h
}

// Update hook auto generate api
func (ai *api) Update(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	handler := func(c *gin.Context) {
		update(name, ai.mgo, c)
	}
	h := createHooks(ai.mgo, handler)
	handlers = append(handlers, h.R)
	r.PUT(fmt.Sprintf("/%s", name), handlers...)
	return h
}

// Delete hook auto generate api
func (ai *api) Delete(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	handler := func(c *gin.Context) {
		remove(name, ai.mgo, c)
	}
	h := createHooks(ai.mgo, handler)
	handlers = append(handlers, h.R)
	r.DELETE(fmt.Sprintf("/%s", name), handlers...)
	return h
}

// ALL hook auto generate api
func (ai *api) ALL(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) {
	r.GET(fmt.Sprintf("/%s", name), combineHF(func(c *gin.Context) {
		list(name, ai.mgo, c)
	}, handlers)...)
	r.GET(fmt.Sprintf("/%s/:id", name), combineHF(func(c *gin.Context) {
		one(name, ai.mgo, c)
	}, handlers)...)
	r.POST(fmt.Sprintf("/%s", name), combineHF(func(c *gin.Context) {
		create(name, ai.mgo, c)
	}, handlers)...)
	r.PUT(fmt.Sprintf("/%s", name), combineHF(func(c *gin.Context) {
		update(name, ai.mgo, c)
	}, handlers)...)
	r.DELETE(fmt.Sprintf("/%s", name), combineHF(func(c *gin.Context) {
		remove(name, ai.mgo, c)
	}, handlers)...)
}
