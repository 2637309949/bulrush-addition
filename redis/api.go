// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package redis

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

// API defined api type
type API struct {
	Client *redis.Client
}

// SetJSON store json data
func (h *API) SetJSON(key string, value interface{}, expiration time.Duration) (*redis.StatusCmd, error) {
	value, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	ret := h.Client.Set(key, value, expiration)
	return ret, nil
}

// GetJSON get json data
func (h *API) GetJSON(key string) (map[string]interface{}, error) {
	var imapGet map[string]interface{}
	value, err := h.Client.Get(key).Result()
	if err != nil {
		return nil, err
	}
	json.Unmarshal([]byte(value), &imapGet)
	return imapGet, nil
}
