// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgoext

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type (
	// API type defined
	API struct {
		mgo  *Mongo
		Opts *Opts
	}
	// Opts defined api params
	Opts struct {
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
		Update *UpdateHookOpts
		Delete *DeleteHookOpts
	}
	// OneHookOpts defined one hook opts
	OneHookOpts struct {
		Pre  gin.HandlerFunc
		Post gin.HandlerFunc
		Auth func(*gin.Context) bool
		Cond func(map[string]interface{}) map[string]interface{}
	}
	// ListHookOpts defined list hook opts
	ListHookOpts struct {
		Pre  gin.HandlerFunc
		Post gin.HandlerFunc
		Auth func(*gin.Context) bool
		Cond func(map[string]interface{}) map[string]interface{}
	}
	// CreateHookOpts defined create hook opts
	CreateHookOpts struct {
		Pre  gin.HandlerFunc
		Post gin.HandlerFunc
		Auth func(*gin.Context) bool
	}
	// UpdateHookOpts defined create hook opts
	UpdateHookOpts struct {
		Pre  gin.HandlerFunc
		Post gin.HandlerFunc
		Auth func(*gin.Context) bool
	}
	// DeleteHookOpts defined delete hook opts
	DeleteHookOpts struct {
		Pre  gin.HandlerFunc
		Post gin.HandlerFunc
		Auth func(*gin.Context) bool
	}
)

func (opts *Opts) prefix() string {
	if opts.Prefix == "" {
		return "/mgo"
	}
	return opts.Prefix
}

func (opts *Opts) mergeOpts(upOpts *Opts) *Opts {
	newOpts := &Opts{
		Prefix:        opts.Prefix,
		FeaturePrefix: opts.FeaturePrefix,
		RouteHooks: &RouteHooks{
			One:    &OneHookOpts{},
			List:   &ListHookOpts{},
			Create: &CreateHookOpts{},
			Update: &UpdateHookOpts{},
			Delete: &DeleteHookOpts{},
		},
		RoutePrefixs: &RoutePrefixs{},
	}
	// merge Prefix
	if upOpts.Prefix != "" {
		newOpts.Prefix = upOpts.Prefix
	}
	if upOpts.FeaturePrefix != "" {
		newOpts.FeaturePrefix = upOpts.FeaturePrefix
	}
	// merge RoutePrefixs
	if upOpts.RoutePrefixs != nil {
		if upOpts.RoutePrefixs.One != nil {
			newOpts.RoutePrefixs.One = upOpts.RoutePrefixs.One
		}
		if upOpts.RoutePrefixs.List != nil {
			newOpts.RoutePrefixs.List = upOpts.RoutePrefixs.List
		}
		if upOpts.RoutePrefixs.Create != nil {
			newOpts.RoutePrefixs.Create = upOpts.RoutePrefixs.Create
		}
		if upOpts.RoutePrefixs.Update != nil {
			newOpts.RoutePrefixs.Update = upOpts.RoutePrefixs.Update
		}
		if upOpts.RoutePrefixs.Delete != nil {
			newOpts.RoutePrefixs.Delete = upOpts.RoutePrefixs.Delete
		}
	}
	// merge RouteHooks
	if upOpts.RouteHooks != nil {
		// merge One
		if upOpts.RouteHooks.One != nil {
			if upOpts.RouteHooks.One.Pre != nil {
				newOpts.RouteHooks.One.Pre = upOpts.RouteHooks.One.Pre
			}
			if upOpts.RouteHooks.One.Post != nil {
				newOpts.RouteHooks.One.Post = upOpts.RouteHooks.One.Post
			}
			if upOpts.RouteHooks.One.Auth != nil {
				newOpts.RouteHooks.One.Auth = upOpts.RouteHooks.One.Auth
			}
			if upOpts.RouteHooks.One.Cond != nil {
				newOpts.RouteHooks.One.Cond = upOpts.RouteHooks.One.Cond
			}
		}
		// merge List
		if upOpts.RouteHooks.List != nil {
			if upOpts.RouteHooks.List.Pre != nil {
				newOpts.RouteHooks.List.Pre = upOpts.RouteHooks.List.Pre
			}
			if upOpts.RouteHooks.List.Post != nil {
				newOpts.RouteHooks.List.Post = upOpts.RouteHooks.List.Post
			}
			if upOpts.RouteHooks.List.Auth != nil {
				newOpts.RouteHooks.List.Auth = upOpts.RouteHooks.List.Auth
			}
			if upOpts.RouteHooks.List.Cond != nil {
				newOpts.RouteHooks.List.Cond = upOpts.RouteHooks.List.Cond
			}
		}
		// merge Create
		if upOpts.RouteHooks.Create != nil {
			if upOpts.RouteHooks.Create.Pre != nil {
				newOpts.RouteHooks.Create.Pre = upOpts.RouteHooks.Create.Pre
			}
			if upOpts.RouteHooks.Create.Post != nil {
				newOpts.RouteHooks.Create.Post = upOpts.RouteHooks.Create.Post
			}
			if upOpts.RouteHooks.Create.Auth != nil {
				newOpts.RouteHooks.Create.Auth = upOpts.RouteHooks.Create.Auth
			}
		}
		// merge Update
		if upOpts.RouteHooks.Update != nil {
			if upOpts.RouteHooks.Update.Pre != nil {
				newOpts.RouteHooks.Update.Pre = upOpts.RouteHooks.Update.Pre
			}
			if upOpts.RouteHooks.Update.Post != nil {
				newOpts.RouteHooks.Update.Post = upOpts.RouteHooks.Update.Post
			}
			if upOpts.RouteHooks.Update.Auth != nil {
				newOpts.RouteHooks.Update.Auth = upOpts.RouteHooks.Update.Auth
			}
		}
		// merge Delete
		if upOpts.RouteHooks.Delete != nil {
			if upOpts.RouteHooks.Delete.Pre != nil {
				newOpts.RouteHooks.Delete.Pre = upOpts.RouteHooks.Delete.Pre
			}
			if upOpts.RouteHooks.Delete.Post != nil {
				newOpts.RouteHooks.Delete.Post = upOpts.RouteHooks.Delete.Post
			}
			if upOpts.RouteHooks.Delete.Auth != nil {
				newOpts.RouteHooks.Delete.Auth = upOpts.RouteHooks.Delete.Auth
			}
		}
	}
	return newOpts
}

func (opts *Opts) featurePrefix() string {
	if opts.FeaturePrefix == "" {
		return ""
	}
	return opts.FeaturePrefix
}

func (opts *Opts) routeHooks() *RouteHooks {
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

func (opts *Opts) routePrefixs() *RoutePrefixs {
	one := func(name string) string {
		return opts.prefix() + opts.featurePrefix() + "/" + name + "/:id"
	}
	list := func(name string) string {
		return opts.prefix() + opts.featurePrefix() + "/" + name
	}
	create := func(name string) string {
		return opts.prefix() + opts.featurePrefix() + "/" + name
	}
	update := func(name string) string {
		return opts.prefix() + opts.featurePrefix() + "/" + name
	}
	delete := func(name string) string {
		return opts.prefix() + opts.featurePrefix() + "/" + name
	}

	newRoutePrefixs := &RoutePrefixs{}
	if opts.RoutePrefixs == nil {
		newRoutePrefixs = &RoutePrefixs{
			One:    one,
			List:   list,
			Create: create,
			Update: update,
			Delete: delete,
		}
	} else {
		if opts.RoutePrefixs.One == nil {
			newRoutePrefixs.One = one
		}
		if opts.RoutePrefixs.List == nil {
			newRoutePrefixs.List = list
		}
		if opts.RoutePrefixs.Create == nil {
			newRoutePrefixs.Create = create
		}
		if opts.RoutePrefixs.Update == nil {
			newRoutePrefixs.Update = update
		}
		if opts.RoutePrefixs.Delete == nil {
			newRoutePrefixs.Delete = delete
		}
	}

	return newRoutePrefixs
}

// Feature defined feature api
func (ai *API) Feature(name string) *API {
	feature := &API{
		mgo: ai.mgo,
		Opts: &Opts{
			Prefix:        ai.Opts.Prefix,
			FeaturePrefix: ai.Opts.FeaturePrefix + "/" + name,
			RoutePrefixs:  ai.Opts.RoutePrefixs,
			RouteHooks:    ai.Opts.RouteHooks,
		},
	}
	return feature
}

// FeatureWithOpts defined feature api with opts
func (ai *API) FeatureWithOpts(opts *Opts) *API {
	feature := &API{mgo: ai.mgo, Opts: opts}
	return feature
}

// One hook auto generate api
func (ai *API) One(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	profile := ai.mgo.Profile(name)
	opts := ai.Opts
	if profile == nil {
		panic(fmt.Errorf("manifest %s not found", name))
	}
	if profile.Opts != nil {
		opts = opts.mergeOpts(profile.Opts)
	}
	routePrefixs := opts.routePrefixs()
	routeHooks := opts.routeHooks()

	handler := func(c *gin.Context) {
		one(name, c, ai.mgo, opts)
	}
	h := createHooks(ai.mgo, handler)
	h.Pre(routeHooks.List.Pre)
	h.Post(routeHooks.List.Post)
	h.Auth(routeHooks.List.Auth)
	handlers = append(handlers, h.R)
	r.GET(routePrefixs.One(name), handlers...)
	return h
}

// List hook auto generate api
func (ai *API) List(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	profile := ai.mgo.Profile(name)
	opts := ai.Opts
	if profile == nil {
		panic(fmt.Errorf("manifest %s not found", name))
	}
	if profile.Opts != nil {
		opts = opts.mergeOpts(profile.Opts)
	}
	routePrefixs := opts.routePrefixs()
	routeHooks := opts.routeHooks()

	handler := func(c *gin.Context) {
		list(name, c, ai.mgo, opts)
	}
	h := createHooks(ai.mgo, handler)
	h.Pre(routeHooks.List.Pre)
	h.Post(routeHooks.List.Post)
	h.Auth(routeHooks.List.Auth)
	handlers = append(handlers, h.R)
	r.GET(routePrefixs.List(name), handlers...)
	return h
}

// Create hook auto generate api
func (ai *API) Create(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	profile := ai.mgo.Profile(name)
	opts := ai.Opts
	if profile == nil {
		panic(fmt.Errorf("manifest %s not found", name))
	}
	if profile.Opts != nil {
		opts = opts.mergeOpts(profile.Opts)
	}
	routePrefixs := opts.routePrefixs()
	routeHooks := opts.routeHooks()

	handler := func(c *gin.Context) {
		create(name, c, ai.mgo, opts)
	}
	h := createHooks(ai.mgo, handler)
	h.Pre(routeHooks.List.Pre)
	h.Post(routeHooks.List.Post)
	h.Auth(routeHooks.List.Auth)
	handlers = append(handlers, h.R)
	r.POST(routePrefixs.Create(name), handlers...)
	return h
}

// Update hook auto generate api
func (ai *API) Update(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	profile := ai.mgo.Profile(name)
	opts := ai.Opts
	if profile == nil {
		panic(fmt.Errorf("manifest %s not found", name))
	}
	if profile.Opts != nil {
		opts = opts.mergeOpts(profile.Opts)
	}
	routePrefixs := opts.routePrefixs()
	routeHooks := opts.routeHooks()

	handler := func(c *gin.Context) {
		update(name, c, ai.mgo, opts)
	}
	h := createHooks(ai.mgo, handler)
	h.Pre(routeHooks.List.Pre)
	h.Post(routeHooks.List.Post)
	h.Auth(routeHooks.List.Auth)
	handlers = append(handlers, h.R)
	r.PUT(routePrefixs.Update(name), handlers...)
	return h
}

// Delete hook auto generate api
func (ai *API) Delete(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	profile := ai.mgo.Profile(name)
	opts := ai.Opts
	if profile == nil {
		panic(fmt.Errorf("manifest %s not found", name))
	}
	if profile.Opts != nil {
		opts = opts.mergeOpts(profile.Opts)
	}
	routePrefixs := opts.routePrefixs()
	routeHooks := opts.routeHooks()

	handler := func(c *gin.Context) {
		remove(name, c, ai.mgo, opts)
	}
	h := createHooks(ai.mgo, handler)
	h.Pre(routeHooks.List.Pre)
	h.Post(routeHooks.List.Post)
	h.Auth(routeHooks.List.Auth)
	handlers = append(handlers, h.R)
	r.DELETE(routePrefixs.Delete(name), handlers...)
	return h
}

// ALL hook auto generate api
func (ai *API) ALL(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	profile := ai.mgo.Profile(name)
	opts := ai.Opts
	if profile == nil {
		panic(fmt.Errorf("manifest %s not found", name))
	}
	if profile.Opts != nil {
		opts = opts.mergeOpts(profile.Opts)
	}
	routePrefixs := opts.routePrefixs()
	routeHooks := opts.routeHooks()

	h := createHooks(ai.mgo, nil)
	r.GET(routePrefixs.One(name), combineHF(func(c *gin.Context) {
		handler := func(c *gin.Context) {
			one(name, c, ai.mgo, opts)
		}
		h1 := createHooks(ai.mgo, handler)
		h1.Pre(routeHooks.One.Pre)
		h1.Post(routeHooks.One.Post)
		h1.Auth(routeHooks.One.Auth)

		h1.Pre(h.pre)
		h1.Post(h.post)
		h1.Auth(h.auth)
		h1.R(c)
	}, handlers)...)
	r.GET(routePrefixs.List(name), combineHF(func(c *gin.Context) {
		handler := func(c *gin.Context) {
			list(name, c, ai.mgo, opts)
		}
		h1 := createHooks(ai.mgo, handler)
		h1.Pre(routeHooks.List.Pre)
		h1.Post(routeHooks.List.Post)
		h1.Auth(routeHooks.List.Auth)

		h1.Pre(h.pre)
		h1.Post(h.post)
		h1.Auth(h.auth)
		h1.R(c)
	}, handlers)...)
	r.POST(routePrefixs.Create(name), combineHF(func(c *gin.Context) {
		handler := func(c *gin.Context) {
			create(name, c, ai.mgo, opts)
		}
		h1 := createHooks(ai.mgo, handler)
		h1.Pre(routeHooks.Create.Pre)
		h1.Post(routeHooks.Create.Post)
		h1.Auth(routeHooks.Create.Auth)

		h1.Pre(h.pre)
		h1.Post(h.post)
		h1.Auth(h.auth)
		h1.R(c)
	}, handlers)...)
	r.PUT(routePrefixs.Update(name), combineHF(func(c *gin.Context) {
		handler := func(c *gin.Context) {
			update(name, c, ai.mgo, opts)
		}
		h1 := createHooks(ai.mgo, handler)
		h1.Pre(routeHooks.Update.Pre)
		h1.Post(routeHooks.Update.Post)
		h1.Auth(routeHooks.Update.Auth)

		h1.Pre(h.pre)
		h1.Post(h.post)
		h1.Auth(h.auth)
		h1.R(c)
	}, handlers)...)
	r.DELETE(routePrefixs.Delete(name), combineHF(func(c *gin.Context) {
		handler := func(c *gin.Context) {
			remove(name, c, ai.mgo, opts)
		}
		h1 := createHooks(ai.mgo, handler)
		h1.Pre(routeHooks.Delete.Pre)
		h1.Post(routeHooks.Delete.Post)
		h1.Auth(routeHooks.Delete.Auth)

		h1.Pre(h.pre)
		h1.Post(h.post)
		h1.Auth(h.auth)
		h1.R(c)
	}, handlers)...)
	return h
}
