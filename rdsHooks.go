package addition

import (
	"encoding/json"
	"time"
)

// SaveToken save a token
func (hook *rdsHooks) SaveToken(token map[string]interface{}) {
	accessToken, _ := token["accessToken"]
	refreshToken, _ := token["refreshToken"]
	value, _ := json.Marshal(token)
	hook.Client.Set("TOKEN:"+accessToken.(string), value, 2*24*time.Hour)
	hook.Client.Set("TOKEN:"+refreshToken.(string), value, 5*24*time.Hour)
}

// RevokeToken revoke a token
func (hook *rdsHooks) RevokeToken(accessToken string) bool {
	status, err := hook.Client.Del("TOKEN:" + accessToken).Result()
	if err != nil {
		return false
	} else if status != 1 {
		return false
	}
	return true
}

// FindToken find a token
func (hook *rdsHooks) FindToken(accessToken string, refreshToken string) map[string]interface{} {
	var imapGet map[string]interface{}
	var token string
	if accessToken != "" {
		token = accessToken
	} else if refreshToken != "" {
		token = refreshToken
	}
	value, err := hook.Client.Get("TOKEN:" + token).Result()
	if err != nil {
		return nil
	}
	json.Unmarshal([]byte(value), &imapGet)
	return imapGet
}
