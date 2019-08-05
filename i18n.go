// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package addition

import (
	"fmt"
	"net/http"

	utils "github.com/2637309949/bulrush-utils"
	"github.com/2637309949/bulrush-utils/funcs"
	"github.com/gin-gonic/gin"
	"github.com/kataras/go-events"
)

const (
	// EventSysBulrushPluginI18NInit defined running event
	EventSysBulrushPluginI18NInit = "EventSysBulrushPluginI18NInit"
)

// M defiend map
type M map[string]interface{}

// I18N defined I18N plugin
type I18N struct {
	Prefix    string
	locale    string
	locales   M
	ctxLocale func(c *gin.Context) string
	initFunc  func()
}

const localeKey = "locale"

// NewI18N defined NewI18N plugin
func NewI18N() *I18N {
	return &I18N{
		Prefix: "/i18n",
		locale: "zh_CN",
		locales: M{
			"zh_CN": M{
				"sys_hello": "你好",
			},
			"en_US": M{
				"sys_hello": "hello",
			},
			"zh_TW": M{
				"sys_hello": "妳好",
			},
		},
		ctxLocale: func(c *gin.Context) string {
			// Query > PostForm > Header > Cookie
			return funcs.Until(
				c.Query(localeKey),
				c.PostForm(localeKey),
				c.Request.Header.Get(localeKey),
				func() interface{} {
					value, _ := c.Cookie(localeKey)
					return value
				},
			).(string)
		},
	}
}

// Init i18n
func (i18n *I18N) Init(init func(*I18N) func()) *I18N {
	i18n.initFunc = init(i18n)
	return i18n
}

// InitLocal defined InitLocal func
func (i18n *I18N) InitLocal(init func(*I18N)) *I18N {
	init(i18n)
	return i18n
}

// Pre defined Pre func
func (i18n *I18N) Pre() {
	if i18n.initFunc != nil {
		i18n.initFunc()
	}
}

// AddLocale defined AddLocale func
func (i18n *I18N) AddLocale(locale string, kv map[string]interface{}) *I18N {
	i18n.locales[locale] = kv
	return i18n
}

// SetCtxLocal defined SetCtxLocal func
func (i18n *I18N) SetCtxLocal(ctxLocal func(c *gin.Context) string) *I18N {
	i18n.ctxLocale = ctxLocal
	return i18n
}

// AddLocales defined AddLocales func
func (i18n *I18N) AddLocales(locales M) *I18N {
	i18n.locales = locales
	return i18n
}

// I18NLocale defined I18NLocale func
func (i18n *I18N) I18NLocale(locale string, init string) string {
	return utils.Some(i18n.locales[i18n.locale].(M)[locale], init).(string)
}

// BuildI18nKey defined BuildI18nKey func
func (i18n *I18N) BuildI18nKey(mod string, str string) string {
	return fmt.Sprintf("%s_%s", mod, str)
}

// GetLocale defined GetLocale func
func (i18n *I18N) GetLocale(locale string) interface{} {
	return i18n.locales[locale]
}

// Plugin defined Plugin func
func (i18n *I18N) Plugin(event events.EventEmmiter, httpProxy *gin.Engine, router *gin.RouterGroup) *I18N {
	httpProxy.Use(func(c *gin.Context) {
		if i18n.ctxLocale(c) != "" {
			i18n.locale = i18n.ctxLocale(c)
		}
		c.Next()
	})
	router.Use(func(c *gin.Context) {
		if i18n.ctxLocale(c) != "" {
			i18n.locale = i18n.ctxLocale(c)
		}
		c.Next()
	})
	router.GET(i18n.Prefix, func(c *gin.Context) {
		c.JSON(http.StatusOK, i18n.locales)
	})
	router.GET(i18n.Prefix+"/:locale", func(c *gin.Context) {
		locale := c.Param("locale")
		c.JSON(http.StatusOK, i18n.locales[locale])
	})
	event.On(EventSysBulrushPluginI18NInit, func(message ...interface{}) {
		if i18n.initFunc != nil {
			i18n.initFunc()
		}
	})
	return i18n
}
