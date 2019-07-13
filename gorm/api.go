// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gormext

import (
	"fmt"
	"time"

	utils "github.com/2637309949/bulrush-utils"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
)

type (
	// API type defined
	API struct {
		gorm *GORM
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
		opts   *Opts
		One    func(string) string
		List   func(string) string
		Create func(string) string
		Update func(string) string
		Delete func(string) string
	}
	// RouteHooks defined route hooks
	RouteHooks struct {
		One    *OneHook
		List   *ListHook
		Create *CreateHook
		Update *UpdateHook
		Delete *DeleteHook
	}
	// OneHook defined one hook opts
	OneHook struct {
		Pre  gin.HandlerFunc
		Post gin.HandlerFunc
		Auth func(c *gin.Context) bool
		Cond func(map[string]interface{}, struct{ name string }) map[string]interface{}
	}
	// ListHook defined list hook opts
	ListHook struct {
		Pre  gin.HandlerFunc
		Post gin.HandlerFunc
		Auth func(*gin.Context) bool
		Cond func(map[string]interface{}, struct{ name string }) map[string]interface{}
	}
	// CreateHook defined create hook opts
	CreateHook struct {
		Pre  gin.HandlerFunc
		Post gin.HandlerFunc
		Auth func(*gin.Context) bool
		Form func(form) form
	}
	// UpdateHook defined create hook opts
	UpdateHook struct {
		Pre  gin.HandlerFunc
		Post gin.HandlerFunc
		Form func(form) form
		Auth func(*gin.Context) bool
	}
	// DeleteHook defined delete hook opts
	DeleteHook struct {
		Pre  gin.HandlerFunc
		Post gin.HandlerFunc
		Auth func(*gin.Context) bool
		Form func(form) form
	}
)

func (one *OneHook) pre() gin.HandlerFunc {
	if one.Pre == nil {
		return func(c *gin.Context) {
		}
	}
	return one.Pre
}

func (one *OneHook) post() gin.HandlerFunc {
	if one.Post == nil {
		return func(c *gin.Context) {
		}
	}
	return one.Post
}

func (one *OneHook) auth() func(c *gin.Context) bool {
	if one.Auth == nil {
		return func(c *gin.Context) bool {
			return true
		}
	}
	return one.Auth
}

func (one *OneHook) cond() func(map[string]interface{}, struct{ name string }) map[string]interface{} {
	if one.Cond == nil {
		return func(cond map[string]interface{}, info struct{ name string }) map[string]interface{} {
			return cond
		}
	}
	return one.Cond
}

func (list *ListHook) pre() gin.HandlerFunc {
	if list.Pre == nil {
		return func(c *gin.Context) {
		}
	}
	return list.Pre
}

func (list *ListHook) post() gin.HandlerFunc {
	if list.Post == nil {
		return func(c *gin.Context) {
		}
	}
	return list.Post
}

func (list *ListHook) auth() func(c *gin.Context) bool {
	if list.Auth == nil {
		return func(c *gin.Context) bool {
			return true
		}
	}
	return list.Auth
}

func (list *ListHook) cond() func(map[string]interface{}, struct{ name string }) map[string]interface{} {
	if list.Cond == nil {
		return func(cond map[string]interface{}, info struct{ name string }) map[string]interface{} {
			return cond
		}
	}
	return list.Cond
}

func (create *CreateHook) pre() gin.HandlerFunc {
	if create.Pre == nil {
		return func(c *gin.Context) {
		}
	}
	return create.Pre
}

func (create *CreateHook) post() gin.HandlerFunc {
	if create.Post == nil {
		return func(c *gin.Context) {
		}
	}
	return create.Post
}

func (create *CreateHook) auth() func(c *gin.Context) bool {
	if create.Auth == nil {
		return func(c *gin.Context) bool {
			return true
		}
	}
	return create.Auth
}

func (create *CreateHook) form() func(form form) form {
	if create.Form == nil {
		return func(form form) form {
			form.Docs = funk.Map(form.Docs, func(i map[string]interface{}) map[string]interface{} {
				i["updatedAt"] = time.Now()
				return i
			}).([]map[string]interface{})
			return form
		}
	}
	return create.Form
}

func (update *UpdateHook) pre() gin.HandlerFunc {
	if update.Pre == nil {
		return func(c *gin.Context) {
		}
	}
	return update.Pre
}

func (update *UpdateHook) post() gin.HandlerFunc {
	if update.Post == nil {
		return func(c *gin.Context) {
		}
	}
	return update.Post
}

func (update *UpdateHook) auth() func(c *gin.Context) bool {
	if update.Auth == nil {
		return func(c *gin.Context) bool {
			return true
		}
	}
	return update.Auth
}

func (update *UpdateHook) form() func(form form) form {
	if update.Form == nil {
		return func(form form) form {
			return form
		}
	}
	return update.Form
}

func (delete *DeleteHook) form() func(form) form {
	if delete.Form == nil {
		return func(form form) form {
			form.Docs = funk.Map(form.Docs, func(doc map[string]interface{}) map[string]interface{} {
				return doc
			}).([]map[string]interface{})
			return form
		}
	}
	return delete.Form
}

func (delete *DeleteHook) pre() gin.HandlerFunc {
	if delete.Pre == nil {
		return func(c *gin.Context) {
		}
	}
	return delete.Pre
}

func (delete *DeleteHook) post() gin.HandlerFunc {
	if delete.Post == nil {
		return func(c *gin.Context) {
		}
	}
	return delete.Post
}

func (delete *DeleteHook) auth() func(c *gin.Context) bool {
	if delete.Auth == nil {
		return func(c *gin.Context) bool {
			return true
		}
	}
	return delete.Auth
}

func (prefixs *RoutePrefixs) one() func(string) string {
	if prefixs.One != nil {
		return prefixs.One
	}
	return func(name string) string {
		return prefixs.opts.Prefix + prefixs.opts.FeaturePrefix + "/" + name + "/:id"
	}
}

func (prefixs *RoutePrefixs) list() func(string) string {
	if prefixs.List != nil {
		return prefixs.List
	}
	return func(name string) string {
		return prefixs.opts.Prefix + prefixs.opts.FeaturePrefix + "/" + name
	}
}

func (prefixs *RoutePrefixs) create() func(string) string {
	if prefixs.Create != nil {
		return prefixs.Create
	}
	return func(name string) string {
		return prefixs.opts.Prefix + prefixs.opts.FeaturePrefix + "/" + name
	}
}

func (prefixs *RoutePrefixs) update() func(string) string {
	if prefixs.Update != nil {
		return prefixs.Update
	}
	return func(name string) string {
		return prefixs.opts.Prefix + prefixs.opts.FeaturePrefix + "/" + name
	}
}

func (prefixs *RoutePrefixs) delete() func(string) string {
	if prefixs.Delete != nil {
		return prefixs.Delete
	}
	return func(name string) string {
		return prefixs.opts.Prefix + prefixs.opts.FeaturePrefix + "/" + name
	}
}

func (route *RouteHooks) one() *OneHook {
	if route.One != nil {
		return route.One
	}
	return &OneHook{}
}

func (route *RouteHooks) list() *ListHook {
	if route.List != nil {
		return route.List
	}
	return &ListHook{}
}

func (route *RouteHooks) create() *CreateHook {
	if route.Create != nil {
		return route.Create
	}
	return &CreateHook{}
}

func (route *RouteHooks) update() *UpdateHook {
	if route.Update != nil {
		return route.Update
	}
	return &UpdateHook{}
}

func (route *RouteHooks) delete() *DeleteHook {
	if route.Delete != nil {
		return route.Delete
	}
	return &DeleteHook{}
}

func (opts *Opts) prefix() string {
	if opts.Prefix == "" {
		return "/gorm"
	}
	return opts.Prefix
}

func (opts *Opts) withDefault() *Opts {
	opts.Prefix = opts.prefix()
	opts.FeaturePrefix = opts.featurePrefix()

	opts.RoutePrefixs = opts.routePrefixs()
	opts.RoutePrefixs.opts = opts
	opts.RoutePrefixs.One = opts.RoutePrefixs.one()
	opts.RoutePrefixs.List = opts.RoutePrefixs.list()
	opts.RoutePrefixs.Create = opts.RoutePrefixs.create()
	opts.RoutePrefixs.Update = opts.RoutePrefixs.update()
	opts.RoutePrefixs.Delete = opts.RoutePrefixs.delete()

	opts.RouteHooks = opts.routeHooks()
	opts.RouteHooks.One = opts.RouteHooks.one()
	opts.RouteHooks.One.Auth = opts.RouteHooks.One.auth()
	opts.RouteHooks.One.Post = opts.RouteHooks.One.post()
	opts.RouteHooks.One.Cond = opts.RouteHooks.One.cond()
	opts.RouteHooks.One.Pre = opts.RouteHooks.One.pre()

	opts.RouteHooks.List = opts.RouteHooks.list()
	opts.RouteHooks.List.Auth = opts.RouteHooks.List.auth()
	opts.RouteHooks.List.Post = opts.RouteHooks.List.post()
	opts.RouteHooks.List.Cond = opts.RouteHooks.List.cond()
	opts.RouteHooks.List.Pre = opts.RouteHooks.List.pre()

	opts.RouteHooks.Update = opts.RouteHooks.update()
	opts.RouteHooks.Update.Auth = opts.RouteHooks.Update.auth()
	opts.RouteHooks.Update.Post = opts.RouteHooks.Update.post()
	opts.RouteHooks.Update.Pre = opts.RouteHooks.Update.pre()
	opts.RouteHooks.Update.Form = opts.RouteHooks.Update.form()

	opts.RouteHooks.Create = opts.RouteHooks.create()
	opts.RouteHooks.Create.Auth = opts.RouteHooks.Create.auth()
	opts.RouteHooks.Create.Post = opts.RouteHooks.Create.post()
	opts.RouteHooks.Create.Pre = opts.RouteHooks.Create.pre()
	opts.RouteHooks.Create.Form = opts.RouteHooks.Create.form()

	opts.RouteHooks.Delete = opts.RouteHooks.delete()
	opts.RouteHooks.Delete.Auth = opts.RouteHooks.Delete.auth()
	opts.RouteHooks.Delete.Post = opts.RouteHooks.Delete.post()
	opts.RouteHooks.Delete.Pre = opts.RouteHooks.Delete.pre()
	opts.RouteHooks.Delete.Form = opts.RouteHooks.Delete.form()
	return opts
}

func (opts *Opts) mergeOpts(upOpts *Opts) *Opts {
	newOpts := &Opts{}
	newOpts = newOpts.withDefault()

	newOpts.Prefix = opts.Prefix
	if upOpts.Prefix != "" {
		newOpts.Prefix = upOpts.Prefix
	}

	newOpts.FeaturePrefix = opts.FeaturePrefix
	if upOpts.FeaturePrefix != "" {
		newOpts.FeaturePrefix = upOpts.FeaturePrefix
	}

	newOpts.RoutePrefixs = opts.RoutePrefixs
	if upOpts.RoutePrefixs != nil {
		if upOpts.RoutePrefixs.opts != nil {
			newOpts.RoutePrefixs.opts = upOpts.RoutePrefixs.opts
		}
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

	if upOpts.RouteHooks != nil {
		if upOpts.RouteHooks.One != nil {
			if upOpts.RouteHooks.One.Pre != nil {
				newOpts.RouteHooks.One.Pre = upOpts.RouteHooks.One.Pre
			}
			if upOpts.RouteHooks.One.Post != nil {
				newOpts.RouteHooks.One.Post = upOpts.RouteHooks.One.Post
			}
			if upOpts.RouteHooks.One.Cond != nil {
				newOpts.RouteHooks.One.Cond = upOpts.RouteHooks.One.Cond
			}
			if upOpts.RouteHooks.One.Auth != nil {
				newOpts.RouteHooks.One.Auth = upOpts.RouteHooks.One.Auth
			}
		}
		if upOpts.RouteHooks.List != nil {
			if upOpts.RouteHooks.List.Pre != nil {
				newOpts.RouteHooks.List.Pre = upOpts.RouteHooks.List.Pre
			}
			if upOpts.RouteHooks.List.Post != nil {
				newOpts.RouteHooks.List.Post = upOpts.RouteHooks.List.Post
			}
			if upOpts.RouteHooks.List.Cond != nil {
				newOpts.RouteHooks.List.Cond = upOpts.RouteHooks.List.Cond
			}
			if upOpts.RouteHooks.List.Auth != nil {
				newOpts.RouteHooks.List.Auth = upOpts.RouteHooks.List.Auth
			}
		}
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
			if upOpts.RouteHooks.Update.Form != nil {
				newOpts.RouteHooks.Update.Form = upOpts.RouteHooks.Update.Form
			}
		}
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
			if upOpts.RouteHooks.Create.Form != nil {
				newOpts.RouteHooks.Create.Form = upOpts.RouteHooks.Create.Form
			}
		}
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
			if upOpts.RouteHooks.Delete.Form != nil {
				newOpts.RouteHooks.Delete.Form = upOpts.RouteHooks.Delete.Form
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
		return &RouteHooks{}
	}
	return opts.RouteHooks
}

func (opts *Opts) routePrefixs() *RoutePrefixs {
	if opts.RoutePrefixs != nil {
		return opts.RoutePrefixs
	}
	return &RoutePrefixs{}
}

// Feature defined feature api
func (ai *API) Feature(name string) *API {
	feature := &API{
		gorm: ai.gorm,
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
	feature := &API{gorm: ai.gorm, Opts: opts}
	return feature
}

// One hook auto generate api
func (ai *API) One(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	profile := ai.gorm.Profile(name)
	opts := ai.Opts.withDefault()
	if profile == nil {
		panic(fmt.Errorf("manifest %s not found", name))
	}
	if profile.Opts != nil {
		opts = opts.mergeOpts(profile.Opts)
	}
	routePrefixs := opts.RoutePrefixs
	routeHooks := opts.RouteHooks

	handler := func(c *gin.Context) {
		one(name, c, ai.gorm, opts)
	}
	h := createHooks(ai.gorm, handler)
	h.Pre(routeHooks.List.Pre)
	h.Post(routeHooks.List.Post)
	h.Auth(routeHooks.List.Auth)
	handlers = append(handlers, h.R)
	r.GET(routePrefixs.One(name), handlers...)
	*profile.docs = append(*profile.docs, *GenDoc(profile, routePrefixs, "one")...)
	return h
}

// List hook auto generate api
func (ai *API) List(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	profile := ai.gorm.Profile(name)
	opts := ai.Opts.withDefault()
	if profile == nil {
		panic(fmt.Errorf("manifest %s not found", name))
	}
	if profile.Opts != nil {
		opts = opts.mergeOpts(profile.Opts)
	}
	routePrefixs := opts.RoutePrefixs
	routeHooks := opts.RouteHooks

	handler := func(c *gin.Context) {
		list(name, c, ai.gorm, opts)
	}
	h := createHooks(ai.gorm, handler)
	h.Pre(routeHooks.List.Pre)
	h.Post(routeHooks.List.Post)
	h.Auth(routeHooks.List.Auth)
	handlers = append(handlers, h.R)
	r.GET(routePrefixs.List(name), handlers...)
	*profile.docs = append(*profile.docs, *GenDoc(profile, routePrefixs, "list")...)
	return h
}

// Create hook auto generate api
func (ai *API) Create(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	profile := ai.gorm.Profile(name)
	opts := ai.Opts.withDefault()
	if profile == nil {
		panic(fmt.Errorf("manifest %s not found", name))
	}
	if profile.Opts != nil {
		opts = opts.mergeOpts(profile.Opts)
	}
	routePrefixs := opts.RoutePrefixs
	routeHooks := opts.RouteHooks

	handler := func(c *gin.Context) {
		create(name, c, ai.gorm, opts)
	}
	h := createHooks(ai.gorm, handler)
	h.Pre(routeHooks.List.Pre)
	h.Post(routeHooks.List.Post)
	h.Auth(routeHooks.List.Auth)
	handlers = append(handlers, h.R)
	r.POST(routePrefixs.Create(name), handlers...)
	*profile.docs = append(*profile.docs, *GenDoc(profile, routePrefixs, "create")...)
	return h
}

// Update hook auto generate api
func (ai *API) Update(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	profile := ai.gorm.Profile(name)
	opts := ai.Opts.withDefault()
	if profile == nil {
		panic(fmt.Errorf("manifest %s not found", name))
	}
	if profile.Opts != nil {
		opts = opts.mergeOpts(profile.Opts)
	}
	routePrefixs := opts.RoutePrefixs
	routeHooks := opts.RouteHooks

	handler := func(c *gin.Context) {
		update(name, c, ai.gorm, opts)
	}
	h := createHooks(ai.gorm, handler)
	h.Pre(routeHooks.List.Pre)
	h.Post(routeHooks.List.Post)
	h.Auth(routeHooks.List.Auth)
	handlers = append(handlers, h.R)
	r.PUT(routePrefixs.Update(name), handlers...)
	*profile.docs = append(*profile.docs, *GenDoc(profile, routePrefixs, "update")...)
	return h
}

// Delete hook auto generate api
func (ai *API) Delete(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	profile := ai.gorm.Profile(name)
	opts := ai.Opts.withDefault()
	if profile == nil {
		panic(fmt.Errorf("manifest %s not found", name))
	}
	if profile.Opts != nil {
		opts = opts.mergeOpts(profile.Opts)
	}
	routePrefixs := opts.RoutePrefixs
	routeHooks := opts.RouteHooks

	handler := func(c *gin.Context) {
		remove(name, c, ai.gorm, opts)
	}
	h := createHooks(ai.gorm, handler)
	h.Pre(routeHooks.List.Pre)
	h.Post(routeHooks.List.Post)
	h.Auth(routeHooks.List.Auth)
	handlers = append(handlers, h.R)
	r.DELETE(routePrefixs.Delete(name), handlers...)
	*profile.docs = append(*profile.docs, *GenDoc(profile, routePrefixs, "update")...)
	return h
}

// ALL hook auto generate api
func (ai *API) ALL(r *gin.RouterGroup, name string, handlers ...gin.HandlerFunc) *Hook {
	profile := ai.gorm.Profile(name)
	opts := ai.Opts.withDefault()
	if profile == nil {
		panic(fmt.Errorf("manifest %s not found", name))
	}
	if profile.Opts != nil {
		opts = opts.mergeOpts(profile.Opts)
	}
	routePrefixs := opts.RoutePrefixs
	routeHooks := opts.RouteHooks

	h := createHooks(ai.gorm, nil)
	h.Pre(routeHooks.List.Pre)
	h.Post(routeHooks.List.Post)
	h.Auth(routeHooks.List.Auth)
	r.GET(routePrefixs.One(name), utils.Append(func(c *gin.Context) {
		handler := func(c *gin.Context) {
			one(name, c, ai.gorm, opts)
		}
		h1 := createHooks(ai.gorm, handler)
		h1.Pre(h.pre)
		h1.Post(h.post)
		h1.Auth(h.auth)
		h1.R(c)
	}, handlers).([]gin.HandlerFunc)...)
	r.GET(routePrefixs.List(name), utils.Append(func(c *gin.Context) {
		handler := func(c *gin.Context) {
			list(name, c, ai.gorm, opts)
		}
		h1 := createHooks(ai.gorm, handler)
		h1.Pre(h.pre)
		h1.Post(h.post)
		h1.Auth(h.auth)
		h1.R(c)
	}, handlers).([]gin.HandlerFunc)...)
	r.POST(routePrefixs.Create(name), utils.Append(func(c *gin.Context) {
		handler := func(c *gin.Context) {
			create(name, c, ai.gorm, opts)
		}
		h1 := createHooks(ai.gorm, handler)
		h1.Pre(h.pre)
		h1.Post(h.post)
		h1.Auth(h.auth)
		h1.R(c)
	}, handlers).([]gin.HandlerFunc)...)
	r.PUT(routePrefixs.Update(name), utils.Append(func(c *gin.Context) {
		handler := func(c *gin.Context) {
			update(name, c, ai.gorm, opts)
		}
		h1 := createHooks(ai.gorm, handler)
		h1.Pre(h.pre)
		h1.Post(h.post)
		h1.Auth(h.auth)
		h1.R(c)
	}, handlers).([]gin.HandlerFunc)...)
	r.DELETE(routePrefixs.Delete(name), utils.Append(func(c *gin.Context) {
		handler := func(c *gin.Context) {
			remove(name, c, ai.gorm, opts)
		}
		h1 := createHooks(ai.gorm, handler)
		h1.Pre(h.pre)
		h1.Post(h.post)
		h1.Auth(h.auth)
		h1.R(c)
	}, handlers).([]gin.HandlerFunc)...)
	*profile.docs = append(*profile.docs, *GenDoc(profile, routePrefixs, "one", "list", "create", "update", "delete")...)
	return h
}
