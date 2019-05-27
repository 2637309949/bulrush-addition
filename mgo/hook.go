/**
 * @author [double]
 * @email [2637309949@qq.com]
 * @create date 2019-03-13 17:25:16
 * @modify date 2019-03-13 17:25:16
 * @desc [description]
 */

package mgo

import (
	"github.com/gin-gonic/gin"
)

type (
	// Hook for API
	Hook struct {
		handler func(c *gin.Context)
		mgo     *Mongo
		pre     func(c *gin.Context)
		post    func(c *gin.Context)
		R       func(c *gin.Context)
	}
)

// route return gin middle
func (h *Hook) route() func(c *gin.Context) {
	return func(c *gin.Context) {
		if h.pre != nil {
			h.pre(c)
		}
		if h.handler != nil {
			h.handler(c)
		}
		if h.post != nil {
			h.post(c)
		}
	}
}

// Pre for pre handler
func (h *Hook) Pre(pre func(*gin.Context)) *Hook {
	h.pre = pre
	return h
}

// Post for pre handler
func (h *Hook) Post(post func(*gin.Context)) *Hook {
	h.post = post
	return h
}

// createHooks alloc new hook
func createHooks(mgo *Mongo, handler func(c *gin.Context)) *Hook {
	h := &Hook{
		mgo:     mgo,
		handler: handler,
	}
	h.R = h.route()
	return h
}
