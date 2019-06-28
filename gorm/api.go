// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gormext

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type (
	// API type defined
	API struct {
		gorm *GORM
		Opts *APIOpts
	}
	// APIOpts defined api params
	APIOpts struct {
		Prefix        string
		FeaturePrefix string
		RoutePrefixs  *RoutePrefixs
		RouteHooks    *RouteHooks
	}
	// RoutePrefixs defined route prefixs
	RoutePrefixs struct {
		One    func(string) string
		List   func(string) string
		Create func(string) string
		Update func(string) string
		Delete func(string) string
	}
	// RouteHooks defined route hooks
	RouteHooks struct {
		One    *OneHookOpts
		List   *ListHookOpts
		Create *CreateHookOpts
		Delete *DeleteHookOpts
	}
	// OneHookOpts defined one hook opts
	OneHookOpts struct {
		Cond func(map[string]interface{}) map[string]interface{}
	}
	// ListHookOpts defined list hook opts
	ListHookOpts struct {
	}
	// CreateHookOpts defined create hook opts
	CreateHookOpts struct {
	}
	// DeleteHookOpts defined delete hook opts
	DeleteHookOpts struct{}
)

func (opts *APIOpts) prefix() string {
	if opts.Prefix == "" {
		return "/gorm"
	}
	return opts.Prefix
}

func (opts *APIOpts) mergeOpts(upOpts *APIOpts) *APIOpts {
	return &APIOpts{
		Prefix:        upOpts.Prefix,
		FeaturePrefix: upOpts.FeaturePrefix,
		RoutePrefixs:  upOpts.RoutePrefixs,
		RouteHooks:    upOpts.RouteHooks,
	}
}

func (opts *APIOpts) featurePrefix() string {
	if opts.FeaturePrefix == "" {
		return ""
	}
	return opts.FeaturePrefix
}

func (opts *APIOpts) routeHooks() *RouteHooks {
	if opts.RouteHooks == nil {
		return &RouteHooks{
			One:    &OneHookOpts{},
			List:   &ListHookOpts{},
			Create: &CreateHookOpts{},
			Delete: &DeleteHookOpts{},
		}
	}
	return opts.RouteHooks
}

func (opts *APIOpts) routePrefixs() *RoutePrefixs {
	if opts.RoutePrefixs == nil {
		return &RoutePrefixs{
			One: func(name string) string {
				return opts.prefix() + opts.featurePrefix() + "/" + name + "/:id"
			},
			List: func(name string) string {
				return opts.prefix() + opts.featurePrefix() + "/" + name
			},
			Create: func(name string) string {
				return opts.prefix() + opts.featurePrefix() + "/" + name
			},
			Update: func(name string) string {
				return opts.prefix() + opts.featurePrefix() + "/" + name
			},
			Delete: func(name string) string {
				return opts.prefix() + opts.featurePrefix() + "/" + name
			},
		}
	}
	return opts.RoutePrefixs
}

// Feature defined feature api
func (ai *API) Feature(name string) *API {
	feature := &API{gorm: ai.gorm, Opts: &APIOpts{FeaturePrefix: ai.Opts.FeaturePrefix + "/" + name}}
	return feature
}

// FeatureWithOpts defined feature api with opts
func (ai *API) FeatureWithOpts(opts *APIOpts) *API {
	feature := &API{gorm: ai.gorm, Opts: opts}
	return feature
}

// One hook auto generate api
func (ai *API) One(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	profile := ai.gorm.Profile(name)
	opts := ai.Opts
	if profile == nil {
		panic(fmt.Errorf("manifest %s not found", name))
	}
	if item, ok := profile["Opts"].(*APIOpts); ok {
		opts = opts.mergeOpts(item)
	}
	routePrefixs := opts.routePrefixs()
	handler := func(c *gin.Context) {
		one(name, ai.gorm, c)
	}
	h := createHooks(ai.gorm, handler)
	handlers = append(handlers, h.R)
	r.GET(routePrefixs.One(name), handlers...)
	return h
}

// List hook auto generate api
func (ai *API) List(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	profile := ai.gorm.Profile(name)
	opts := ai.Opts
	if profile == nil {
		panic(fmt.Errorf("manifest %s not found", name))
	}
	if item, ok := profile["Opts"].(*APIOpts); ok {
		opts = opts.mergeOpts(item)
	}
	routePrefixs := opts.routePrefixs()
	handler := func(c *gin.Context) {
		list(name, ai.gorm, c)
	}
	h := createHooks(ai.gorm, handler)
	handlers = append(handlers, h.R)
	r.GET(routePrefixs.List(name), handlers...)
	return h
}

// Create hook auto generate api
func (ai *API) Create(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	profile := ai.gorm.Profile(name)
	opts := ai.Opts
	if profile == nil {
		panic(fmt.Errorf("manifest %s not found", name))
	}
	if item, ok := profile["Opts"].(*APIOpts); ok {
		opts = opts.mergeOpts(item)
	}
	routePrefixs := opts.routePrefixs()
	handler := func(c *gin.Context) {
		create(name, ai.gorm, c)
	}
	h := createHooks(ai.gorm, handler)
	handlers = append(handlers, h.R)
	r.POST(routePrefixs.Create(name), handlers...)
	return h
}

// Update hook auto generate api
func (ai *API) Update(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	profile := ai.gorm.Profile(name)
	opts := ai.Opts
	if profile == nil {
		panic(fmt.Errorf("manifest %s not found", name))
	}
	if item, ok := profile["Opts"].(*APIOpts); ok {
		opts = opts.mergeOpts(item)
	}
	routePrefixs := opts.routePrefixs()
	handler := func(c *gin.Context) {
		update(name, ai.gorm, c)
	}
	h := createHooks(ai.gorm, handler)
	handlers = append(handlers, h.R)
	r.PUT(routePrefixs.Update(name), handlers...)
	return h
}

// Delete hook auto generate api
func (ai *API) Delete(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	profile := ai.gorm.Profile(name)
	opts := ai.Opts
	if profile == nil {
		panic(fmt.Errorf("manifest %s not found", name))
	}
	if item, ok := profile["Opts"].(*APIOpts); ok {
		opts = opts.mergeOpts(item)
	}
	routePrefixs := opts.routePrefixs()
	handler := func(c *gin.Context) {
		remove(name, ai.gorm, c)
	}
	h := createHooks(ai.gorm, handler)
	handlers = append(handlers, h.R)
	r.DELETE(routePrefixs.Delete(name), handlers...)
	return h
}

// ALL hook auto generate api
func (ai *API) ALL(r *gin.RouterGroup, name string) {
	profile := ai.gorm.Profile(name)
	opts := ai.Opts
	if profile == nil {
		panic(fmt.Errorf("manifest %s not found", name))
	}
	if item, ok := profile["Opts"].(*APIOpts); ok {
		opts = opts.mergeOpts(item)
	}
	routePrefixs := opts.routePrefixs()
	r.GET(routePrefixs.One(name), func(c *gin.Context) {
		one(name, ai.gorm, c)
	})
	r.GET(routePrefixs.List(name), func(c *gin.Context) {
		list(name, ai.gorm, c)
	})
	r.POST(routePrefixs.Create(name), func(c *gin.Context) {
		create(name, ai.gorm, c)
	})
	r.PUT(routePrefixs.Update(name), func(c *gin.Context) {
		update(name, ai.gorm, c)
	})
	r.DELETE(routePrefixs.Delete(name), func(c *gin.Context) {
		remove(name, ai.gorm, c)
	})
}
