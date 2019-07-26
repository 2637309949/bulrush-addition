// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gormext

import (
	"fmt"
	"strings"
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
		Pre  func(*gin.Context)
		Post func(*gin.Context)
		Auth func(c *gin.Context) bool
		Cond func(map[string]interface{}, *gin.Context, struct{ Name string }) map[string]interface{}
	}
	// ListHook defined list hook opts
	ListHook struct {
		Pre  func(*gin.Context)
		Post func(*gin.Context)
		Auth func(*gin.Context) bool
		Cond func(map[string]interface{}, *gin.Context, struct{ Name string }) map[string]interface{}
	}
	// CreateHook defined create hook opts
	CreateHook struct {
		Pre  func(*gin.Context)
		Post func(*gin.Context)
		Auth func(*gin.Context) bool
		Form func(form) form
	}
	// UpdateHook defined create hook opts
	UpdateHook struct {
		Pre  func(*gin.Context)
		Post func(*gin.Context)
		Auth func(*gin.Context) bool
		Form func(form) form
		Cond func(map[string]interface{}, *gin.Context, struct{ Name string }) map[string]interface{}
	}
	// DeleteHook defined delete hook opts
	DeleteHook struct {
		Pre  func(*gin.Context)
		Post func(*gin.Context)
		Auth func(*gin.Context) bool
		Form func(form) form
		Cond func(map[string]interface{}, *gin.Context, struct{ Name string }) map[string]interface{}
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

func (one *OneHook) cond() func(map[string]interface{}, *gin.Context, struct{ Name string }) map[string]interface{} {
	if one.Cond == nil {
		return func(cond map[string]interface{}, c *gin.Context, info struct{ Name string }) map[string]interface{} {
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

func (list *ListHook) cond() func(map[string]interface{}, *gin.Context, struct{ Name string }) map[string]interface{} {
	if list.Cond == nil {
		return func(cond map[string]interface{}, c *gin.Context, info struct{ Name string }) map[string]interface{} {
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

func (update *UpdateHook) cond() func(map[string]interface{}, *gin.Context, struct{ Name string }) map[string]interface{} {
	if update.Cond == nil {
		return func(cond map[string]interface{}, c *gin.Context, info struct{ Name string }) map[string]interface{} {
			return cond
		}
	}
	return update.Cond
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

func (delete *DeleteHook) cond() func(map[string]interface{}, *gin.Context, struct{ Name string }) map[string]interface{} {
	if delete.Cond == nil {
		return func(cond map[string]interface{}, c *gin.Context, info struct{ Name string }) map[string]interface{} {
			return cond
		}
	}
	return delete.Cond
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
		return prefixs.opts.Prefix + prefixs.opts.FeaturePrefix + "/" + strings.ToLower(name) + "/:id"
	}
}

func (prefixs *RoutePrefixs) list() func(string) string {
	if prefixs.List != nil {
		return prefixs.List
	}
	return func(name string) string {
		return prefixs.opts.Prefix + prefixs.opts.FeaturePrefix + "/" + strings.ToLower(name)
	}
}

func (prefixs *RoutePrefixs) create() func(string) string {
	if prefixs.Create != nil {
		return prefixs.Create
	}
	return func(name string) string {
		return prefixs.opts.Prefix + prefixs.opts.FeaturePrefix + "/" + strings.ToLower(name)
	}
}

func (prefixs *RoutePrefixs) update() func(string) string {
	if prefixs.Update != nil {
		return prefixs.Update
	}
	return func(name string) string {
		return prefixs.opts.Prefix + prefixs.opts.FeaturePrefix + "/" + strings.ToLower(name)
	}
}

func (prefixs *RoutePrefixs) delete() func(string) string {
	if prefixs.Delete != nil {
		return prefixs.Delete
	}
	return func(name string) string {
		return prefixs.opts.Prefix + prefixs.opts.FeaturePrefix + "/" + strings.ToLower(name)
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
	opts.RouteHooks.Update.Cond = opts.RouteHooks.Update.cond()

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
	opts.RouteHooks.Delete.Cond = opts.RouteHooks.Delete.cond()
	return opts
}

func (opts *Opts) mergeOpts(upOpts *Opts) *Opts {
	newOpts := &Opts{}
	newOpts = newOpts.withDefault()

	newOpts.Prefix = utils.Some(upOpts.Prefix, opts.Prefix).(string)
	newOpts.FeaturePrefix = utils.Some(upOpts.FeaturePrefix, opts.FeaturePrefix).(string)

	upOpts.RoutePrefixs = utils.Some(upOpts.RoutePrefixs, opts.RoutePrefixs).(*RoutePrefixs)
	newOpts.RoutePrefixs.opts = utils.Some(upOpts.RoutePrefixs.opts, opts.RoutePrefixs.opts).(*Opts)
	newOpts.RoutePrefixs.One = utils.Some(upOpts.RoutePrefixs.One, opts.RoutePrefixs.One).(func(string) string)
	newOpts.RoutePrefixs.List = utils.Some(upOpts.RoutePrefixs.List, opts.RoutePrefixs.List).(func(string) string)
	newOpts.RoutePrefixs.Create = utils.Some(upOpts.RoutePrefixs.Create, opts.RoutePrefixs.Create).(func(string) string)
	newOpts.RoutePrefixs.Update = utils.Some(upOpts.RoutePrefixs.Update, opts.RoutePrefixs.Update).(func(string) string)
	newOpts.RoutePrefixs.Delete = utils.Some(upOpts.RoutePrefixs.Delete, opts.RoutePrefixs.Delete).(func(string) string)

	upOpts.RouteHooks = utils.Some(upOpts.RouteHooks, opts.RouteHooks).(*RouteHooks)
	upOpts.RouteHooks.One = utils.Some(upOpts.RouteHooks.One, opts.RouteHooks.One).(*OneHook)
	newOpts.RouteHooks.One.Pre = utils.Some(upOpts.RouteHooks.One.Pre, opts.RouteHooks.One.Pre).(func(*gin.Context))
	newOpts.RouteHooks.One.Post = utils.Some(upOpts.RouteHooks.One.Post, opts.RouteHooks.One.Post).(func(*gin.Context))
	newOpts.RouteHooks.One.Cond = utils.Some(upOpts.RouteHooks.One.Cond, opts.RouteHooks.One.Cond).(func(map[string]interface{}, *gin.Context, struct{ Name string }) map[string]interface{})
	newOpts.RouteHooks.One.Auth = utils.Some(upOpts.RouteHooks.One.Auth, opts.RouteHooks.One.Auth).(func(c *gin.Context) bool)

	upOpts.RouteHooks.List = utils.Some(upOpts.RouteHooks.List, opts.RouteHooks.List).(*ListHook)
	newOpts.RouteHooks.List.Pre = utils.Some(upOpts.RouteHooks.List.Pre, opts.RouteHooks.List.Pre).(func(*gin.Context))
	newOpts.RouteHooks.List.Post = utils.Some(upOpts.RouteHooks.List.Post, opts.RouteHooks.List.Post).(func(*gin.Context))
	newOpts.RouteHooks.List.Cond = utils.Some(upOpts.RouteHooks.List.Cond, opts.RouteHooks.List.Cond).(func(map[string]interface{}, *gin.Context, struct{ Name string }) map[string]interface{})
	newOpts.RouteHooks.List.Auth = utils.Some(upOpts.RouteHooks.List.Auth, opts.RouteHooks.List.Auth).(func(c *gin.Context) bool)

	upOpts.RouteHooks.Update = utils.Some(upOpts.RouteHooks.Update, opts.RouteHooks.Update).(*UpdateHook)
	newOpts.RouteHooks.Update.Pre = utils.Some(upOpts.RouteHooks.Update.Pre, opts.RouteHooks.Update.Pre).(func(*gin.Context))
	newOpts.RouteHooks.Update.Post = utils.Some(upOpts.RouteHooks.Update.Post, opts.RouteHooks.Update.Post).(func(*gin.Context))
	newOpts.RouteHooks.Update.Cond = utils.Some(upOpts.RouteHooks.Update.Cond, opts.RouteHooks.Update.Cond).(func(map[string]interface{}, *gin.Context, struct{ Name string }) map[string]interface{})
	newOpts.RouteHooks.Update.Form = utils.Some(upOpts.RouteHooks.Update.Form, opts.RouteHooks.Update.Form).(func(form) form)
	newOpts.RouteHooks.Update.Auth = utils.Some(upOpts.RouteHooks.Update.Auth, opts.RouteHooks.Update.Auth).(func(c *gin.Context) bool)

	upOpts.RouteHooks.Create = utils.Some(upOpts.RouteHooks.Create, opts.RouteHooks.Create).(*CreateHook)
	newOpts.RouteHooks.Create.Pre = utils.Some(upOpts.RouteHooks.Create.Pre, opts.RouteHooks.Create.Pre).(func(*gin.Context))
	newOpts.RouteHooks.Create.Post = utils.Some(upOpts.RouteHooks.Create.Post, opts.RouteHooks.Create.Post).(func(*gin.Context))
	newOpts.RouteHooks.Create.Form = utils.Some(upOpts.RouteHooks.Create.Form, opts.RouteHooks.Create.Form).(func(form) form)
	newOpts.RouteHooks.Create.Auth = utils.Some(upOpts.RouteHooks.Create.Auth, opts.RouteHooks.Create.Auth).(func(c *gin.Context) bool)

	upOpts.RouteHooks.Delete = utils.Some(upOpts.RouteHooks.Delete, opts.RouteHooks.Delete).(*DeleteHook)
	newOpts.RouteHooks.Delete.Pre = utils.Some(upOpts.RouteHooks.Delete.Pre, opts.RouteHooks.Delete.Pre).(func(*gin.Context))
	newOpts.RouteHooks.Delete.Post = utils.Some(upOpts.RouteHooks.Delete.Post, opts.RouteHooks.Delete.Post).(func(*gin.Context))
	newOpts.RouteHooks.Delete.Cond = utils.Some(upOpts.RouteHooks.Delete.Cond, opts.RouteHooks.Delete.Cond).(func(map[string]interface{}, *gin.Context, struct{ Name string }) map[string]interface{})
	newOpts.RouteHooks.Delete.Form = utils.Some(upOpts.RouteHooks.Delete.Form, opts.RouteHooks.Delete.Form).(func(form) form)
	newOpts.RouteHooks.Delete.Auth = utils.Some(upOpts.RouteHooks.Delete.Auth, opts.RouteHooks.Delete.Auth).(func(c *gin.Context) bool)
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
	h.Pre(routeHooks.One.Pre)
	h.Post(routeHooks.One.Post)
	h.Auth(routeHooks.One.Auth)
	h.routeHooks = func(hooks *RouteHooks) {
		newOpts := &Opts{}
		newOpts.RouteHooks = hooks
		opts.RouteHooks = opts.mergeOpts(newOpts).RouteHooks
	}

	handlers = append(handlers, h.r)
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
	h.routeHooks = func(hooks *RouteHooks) {
		newOpts := &Opts{}
		newOpts.RouteHooks = hooks
		opts.RouteHooks = opts.mergeOpts(newOpts).RouteHooks
	}

	handlers = append(handlers, h.r)
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
	h.Pre(routeHooks.Create.Pre)
	h.Post(routeHooks.Create.Post)
	h.Auth(routeHooks.Create.Auth)
	h.routeHooks = func(hooks *RouteHooks) {
		newOpts := &Opts{}
		newOpts.RouteHooks = hooks
		opts.RouteHooks = opts.mergeOpts(newOpts).RouteHooks
	}

	handlers = append(handlers, h.r)
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
	h.Pre(routeHooks.Update.Pre)
	h.Post(routeHooks.Update.Post)
	h.Auth(routeHooks.Update.Auth)
	h.routeHooks = func(hooks *RouteHooks) {
		newOpts := &Opts{}
		newOpts.RouteHooks = hooks
		opts.RouteHooks = opts.mergeOpts(newOpts).RouteHooks
	}

	handlers = append(handlers, h.r)
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
	h.Pre(routeHooks.Delete.Pre)
	h.Post(routeHooks.Delete.Post)
	h.Auth(routeHooks.Delete.Auth)
	h.routeHooks = func(hooks *RouteHooks) {
		newOpts := &Opts{}
		newOpts.RouteHooks = hooks
		opts.RouteHooks = opts.mergeOpts(newOpts).RouteHooks
	}

	handlers = append(handlers, h.r)
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

	h := createHooks(ai.gorm, nil)
	h.routeHooks = func(hooks *RouteHooks) {
		newOpts := &Opts{}
		newOpts.RouteHooks = hooks
		opts.RouteHooks = opts.mergeOpts(newOpts).RouteHooks
	}
	r.GET(routePrefixs.One(name), utils.Append(func(c *gin.Context) {
		handler := func(c *gin.Context) {
			one(name, c, ai.gorm, opts)
		}
		h1 := createHooks(ai.gorm, handler)
		h1.Pre(h.pre)
		h1.Post(h.post)
		h1.Auth(h.auth)
		h1.routeHooks = func(hooks *RouteHooks) {
			newOpts := &Opts{}
			newOpts.RouteHooks = hooks
			opts.RouteHooks = opts.mergeOpts(newOpts).RouteHooks
		}
		h1.r(c)
	}, handlers).([]gin.HandlerFunc)...)
	r.GET(routePrefixs.List(name), utils.Append(func(c *gin.Context) {
		handler := func(c *gin.Context) {
			list(name, c, ai.gorm, opts)
		}
		h1 := createHooks(ai.gorm, handler)
		h1.Pre(h.pre)
		h1.Post(h.post)
		h1.Auth(h.auth)
		h1.routeHooks = func(hooks *RouteHooks) {
			newOpts := &Opts{}
			newOpts.RouteHooks = hooks
			opts.RouteHooks = opts.mergeOpts(newOpts).RouteHooks
		}
		h1.r(c)
	}, handlers).([]gin.HandlerFunc)...)
	r.POST(routePrefixs.Create(name), utils.Append(func(c *gin.Context) {
		handler := func(c *gin.Context) {
			create(name, c, ai.gorm, opts)
		}
		h1 := createHooks(ai.gorm, handler)
		h1.Pre(h.pre)
		h1.Post(h.post)
		h1.Auth(h.auth)
		h1.routeHooks = func(hooks *RouteHooks) {
			newOpts := &Opts{}
			newOpts.RouteHooks = hooks
			opts.RouteHooks = opts.mergeOpts(newOpts).RouteHooks
		}
		h1.r(c)
	}, handlers).([]gin.HandlerFunc)...)
	r.PUT(routePrefixs.Update(name), utils.Append(func(c *gin.Context) {
		handler := func(c *gin.Context) {
			update(name, c, ai.gorm, opts)
		}
		h1 := createHooks(ai.gorm, handler)
		h1.Pre(h.pre)
		h1.Post(h.post)
		h1.Auth(h.auth)
		h1.routeHooks = func(hooks *RouteHooks) {
			newOpts := &Opts{}
			newOpts.RouteHooks = hooks
			opts.RouteHooks = opts.mergeOpts(newOpts).RouteHooks
		}
		h1.r(c)
	}, handlers).([]gin.HandlerFunc)...)
	r.DELETE(routePrefixs.Delete(name), utils.Append(func(c *gin.Context) {
		handler := func(c *gin.Context) {
			remove(name, c, ai.gorm, opts)
		}
		h1 := createHooks(ai.gorm, handler)
		h1.Pre(h.pre)
		h1.Post(h.post)
		h1.Auth(h.auth)
		h1.routeHooks = func(hooks *RouteHooks) {
			newOpts := &Opts{}
			newOpts.RouteHooks = hooks
			opts.RouteHooks = opts.mergeOpts(newOpts).RouteHooks
		}
		h1.r(c)
	}, handlers).([]gin.HandlerFunc)...)
	*profile.docs = append(*profile.docs, *GenDoc(profile, routePrefixs, "one", "list", "create", "update", "delete")...)
	return h
}
