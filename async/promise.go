/**
 * @author [Double]
 * @email [2637309949@qq.com.com]
 * @create date 2019-01-12 22:46:31
 * @modify date 2019-01-12 22:46:31
 * @desc [bulrush async plugin]
 */

package async

import (
	"reflect"
)

// STATES validStates type
type STATES int

const (
	// PENDING validStates
	PENDING STATES = iota + 1
	// FULFILLED validStates
	FULFILLED
	// REJECTED validStates
	REJECTED
)

type (
	// promise implement Nodejs p+
	promise struct {
		value    interface{}
		state    STATES
		queue    []*promise
		handlers struct {
			fulfill interface{}
			reject  interface{}
		}
	}
)

func (p *promise) reject(reason interface{}) {
	p.transition(REJECTED, reason)
}

func (p *promise) fulfill(value interface{}) {
	p.transition(FULFILLED, value)
}

func (p *promise) transition(state STATES, value interface{}) {
	if p.state == state || p.state != FULFILLED {
		return
	}
	p.value = value
	p.state = state
	p.process()
}

func (p *promise) process() {
	if p.state == PENDING {
		return
	}
	fulfillFallBack := func(value interface{}) interface{} {
		return value
	}
	rejectFallBack := func(reason interface{}) interface{} {
		return reason
	}
	runAsync(func() {
		for len(p.queue) > 0 {
			var queuedPromise *promise
			var handler interface{}
			queuedPromise, p.queue = p.queue[0], p.queue[1:]
			if p.state == FULFILLED {
				handler = queuedPromise.handlers.fulfill
				if handler != nil {
					handler = fulfillFallBack
				}
			} else if p.state == REJECTED {
				handler = queuedPromise.handlers.reject
				if handler != nil {
					handler = rejectFallBack
				}
			}
			value := reflect.ValueOf(handler).Call([]reflect.Value{reflect.ValueOf(p.value)})[0].Interface()
			resolve(queuedPromise, value)
		}
	})
}

func resolve(p *promise, x interface{}) {
}

func runAsync(funk func()) {
	go funk()
}

func gPromise(funk interface{}) *promise {
	if reflect.TypeOf(funk).Kind() != reflect.Func {
		panic("Argument must be function")
	}
	p := &promise{}
	funcValue := reflect.ValueOf(funk)
	funcType := reflect.TypeOf(funk)
	params := []reflect.Value{}
	resolve := func(value interface{}) {
		resolve(p, value)
	}
	reject := func(reason interface{}) {
		p.reject(reason)
	}
	numIn := funcType.NumIn()
	numRes := 1
	numResRej := 2
	if numIn <= numResRej {
		if numIn == numResRej {
			params = append(params, reflect.ValueOf(resolve))
			params = append(params, reflect.ValueOf(reject))
		} else if numIn == numRes {
			params = append(params, reflect.ValueOf(resolve))
		}
	} else {
		panic("Argument Func arguments must lte 2")
	}
	funcValue.Call(params)
	return p
}

// Deferred export a Promise
func Deferred() {
}

// Rejected export a Promise
func Rejected(reason interface{}) {
}

// Resolved export a Promise
func Resolved(value interface{}) {
}
