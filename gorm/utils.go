// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gormext

import (
	"bytes"
	"reflect"
	"strings"

	utils "github.com/2637309949/bulrush-utils"
)

var smap = utils.NewSafeMap()
var commonInitialisms = []string{"API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SSH", "TLS", "TTL", "UID", "UI", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XSRF", "XSS"}
var commonInitialismsReplacer = columnReplacer(commonInitialisms)

func findFieldStruct(vType reflect.Type, name string) *reflect.StructField {
	if vType.Kind() == reflect.Ptr {
		vType = vType.Elem()
	}
	if vType.Kind() == reflect.Struct {
		numField := vType.NumField()
		if numField > 0 {
			field, ok := vType.FieldByName(name)
			if ok {
				return &field
			}
			field = vType.Field(0)
			return findFieldStruct(field.Type, name)
		}
	}
	return nil
}

func columnReplacer(sms []string) *strings.Replacer {
	var list []string
	for _, initialism := range sms {
		list = append(list, initialism, strings.Title(strings.ToLower(initialism)))
	}
	replacer := strings.NewReplacer(list...)
	return replacer
}

// columnNamer copy from gorm
func columnNamer(name string) string {
	const (
		lower = false
		upper = true
	)
	if v := smap.Get(name); v != "" {
		return v
	}
	if name == "" {
		return ""
	}
	var (
		value                                    = commonInitialismsReplacer.Replace(name)
		buf                                      = bytes.NewBufferString("")
		lastCase, currCase, nextCase, nextNumber bool
	)
	for i, v := range value[:len(value)-1] {
		nextCase = bool(value[i+1] >= 'A' && value[i+1] <= 'Z')
		nextNumber = bool(value[i+1] >= '0' && value[i+1] <= '9')
		if i > 0 {
			if currCase == upper {
				if lastCase == upper && (nextCase == upper || nextNumber == upper) {
					buf.WriteRune(v)
				} else {
					if value[i-1] != '_' && value[i+1] != '_' {
						buf.WriteRune('_')
					}
					buf.WriteRune(v)
				}
			} else {
				buf.WriteRune(v)
				if i == len(value)-2 && (nextCase == upper && nextNumber == lower) {
					buf.WriteRune('_')
				}
			}
		} else {
			currCase = upper
			buf.WriteRune(v)
		}
		lastCase = currCase
		currCase = nextCase
	}
	buf.WriteByte(value[len(value)-1])
	s := strings.ToLower(buf.String())
	smap.Set(name, s)
	return s
}

func toColumnName(m *map[string]interface{}) {
	for k, v := range *m {
		c := columnNamer(k)
		(*m)[c] = v
		delete(*m, k)
	}
}
