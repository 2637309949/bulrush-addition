/**
 * @author [double]
 * @email [2637309949@qq.com]
 * @create date 2019-03-13 17:25:16
 * @modify date 2019-03-13 17:25:16
 * @desc [description]
 */

package mongo

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type (
	// API type defined
	API struct {
		mgo *Mongo
	}
)

// One hook auto generate api
func (api *API) One(r *gin.RouterGroup, name string) *Hook {
	handler := func(c *gin.Context) {
		one(name, api.mgo, c)
	}
	h := createHooks(api.mgo, handler)
	r.GET(fmt.Sprintf("/%s/:id", name), h.R)
	return h
}

// List hook auto generate api
func (api *API) List(r *gin.RouterGroup, name string) *Hook {
	handler := func(c *gin.Context) {
		list(name, api.mgo, c)
	}
	h := createHooks(api.mgo, handler)
	r.GET(fmt.Sprintf("/%s", name), h.R)
	return h
}

// Create hook auto generate api
func (api *API) Create(r *gin.RouterGroup, name string) *Hook {
	handler := func(c *gin.Context) {
		create(name, api.mgo, c)
	}
	h := createHooks(api.mgo, handler)
	r.POST(fmt.Sprintf("/%s", name), h.R)
	return h
}

// Update hook auto generate api
func (api *API) Update(r *gin.RouterGroup, name string) *Hook {
	handler := func(c *gin.Context) {
		update(name, api.mgo, c)
	}
	h := createHooks(api.mgo, handler)
	r.PUT(fmt.Sprintf("/%s", name), h.R)
	return h
}

// Delete hook auto generate api
func (api *API) Delete(r *gin.RouterGroup, name string) *Hook {
	handler := func(c *gin.Context) {
		delete(name, api.mgo, c)
	}
	h := createHooks(api.mgo, handler)
	r.DELETE(fmt.Sprintf("/%s", name), h.R)
	return h
}

// ALL hook auto generate api
func (api *API) ALL(r *gin.RouterGroup, name string) {
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
