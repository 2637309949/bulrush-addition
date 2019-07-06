// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgoext

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/thoas/go-funk"
)

type (
	// Doc defined model api doc
	Doc struct {
		Type       string    `json:"type" yaml:"type"`
		URL        string    `json:"url" yaml:"url"`
		Title      string    `json:"title" yaml:"title"`
		Name       string    `json:"name" yaml:"name"`
		Group      string    `json:"group" yaml:"group"`
		Desc       string    `json:"description" yaml:"description"`
		Version    string    `json:"version" yaml:"version"`
		GroupTitle string    `json:"groupTitle" yaml:"groupTitle"`
		Parameter  Parameter `json:"parameter" yaml:"parameter"`
		Success    Success   `json:"success" yaml:"success"`
	}
	// Parameter defined model api doc
	Parameter struct {
		Fields Fields `json:"fields" yaml:"fields"`
	}
	// Fields defined model api doc
	Fields struct {
		FieldsParameter []FieldsParameter `json:"Parameter" yaml:"Parameter"`
	}
	// FieldsSuccess defined model api doc
	FieldsSuccess struct {
		FieldsSuccess []FieldsParameter `json:"Success 200" yaml:"Success 200"`
	}
	// FieldsParameter defined model api doc
	FieldsParameter struct {
		Tags     reflect.StructTag `json:"-" yaml:"-"`
		Group    string            `json:"group" yaml:"group"`
		Type     string            `json:"type" yaml:"type"`
		Optional bool              `json:"optional" yaml:"optional"`
		Field    string            `json:"field" yaml:"field"`
		Desc     string            `json:"description" yaml:"description"`
	}
	// Success defined model api doc
	Success struct {
		FieldsSuccess FieldsSuccess `json:"fields" yaml:"fields"`
	}
)

func fieldScope(structType reflect.Type, items *[]FieldsParameter) {
	itemType := structType
	if itemType.Kind() == reflect.Ptr {
		itemType = itemType.Elem()
	}
	if itemType.Kind() == reflect.Struct {
		count := itemType.NumField()
		for index := 0; index < count; index++ {
			field := itemType.Field(index)
			fieldType := field.Type
			if fieldType.Kind() == reflect.Ptr {
				fieldType = fieldType.Elem()
			}
			// extra first level component struct
			if index == 0 && fieldType.Kind() == reflect.Struct {
				fieldScope(fieldType, items)
			} else {
				optional := strings.Contains(field.Tag.Get("bson"), "not null")
				descArr := findStringSubmatch(`comment:(.*?),`, field.Tag.Get("bson"))
				enumArr := findStringSubmatch(`enum:(.*?),`, field.Tag.Get("bson"))
				desc := ""
				if len(descArr) > 0 {
					desc = descArr[0]
				}
				if len(enumArr) > 0 {
					desc = desc + "( " + enumArr[0] + " )"
				}
				*items = append(*items, FieldsParameter{
					Tags:     field.Tag,
					Group:    "Parameter",
					Type:     fmt.Sprintf("%v", fieldType),
					Optional: optional,
					Field:    field.Name,
					Desc:     desc,
				})
			}
		}
	}
}

// GenDoc defined doc
func GenDoc(profile *Profile, routePrefixs *RoutePrefixs, apis ...string) *[]Doc {
	// fetch info before gen
	modelType := reflect.TypeOf(profile.Reflector)
	fieldsParameter := []FieldsParameter{}
	fieldScope(modelType, &fieldsParameter)

	// start gen
	docs := []Doc{}
	funk.ForEach(apis, func(api string) {
		if api == "one" {
			fieldsSuccess := append([]FieldsParameter{FieldsParameter{Group: "Success 200", Type: "Object", Field: "", Desc: "one of " + profile.Name}}, funk.Map(fieldsParameter, func(p FieldsParameter) FieldsParameter {
				p.Group = "Success 200"
				p.Field = "." + p.Field
				return p
			}).([]FieldsParameter)...)
			docs = append(docs, Doc{
				Type:       "get",
				URL:        routePrefixs.One(profile.Name),
				Title:      fmt.Sprintf("%s one", profile.Name),
				Name:       fmt.Sprintf("%s one", profile.Name),
				Group:      "NoSql Default",
				GroupTitle: "NoSql Default",
				Version:    "0.0.0",
				Parameter: Parameter{
					Fields: Fields{
						FieldsParameter: fieldsParameter,
					},
				},
				Success: Success{
					FieldsSuccess: FieldsSuccess{
						FieldsSuccess: fieldsSuccess,
					},
				},
			})
		} else if api == "list" {
			fieldsSuccess := append([]FieldsParameter{FieldsParameter{Group: "Success 200", Type: "Object[]", Field: "", Desc: "list of " + profile.Name}}, funk.Map(fieldsParameter, func(p FieldsParameter) FieldsParameter {
				p.Group = "Success 200"
				p.Field = "." + p.Field
				return p
			}).([]FieldsParameter)...)
			docs = append(docs, Doc{
				Type:       "get",
				URL:        routePrefixs.List(profile.Name),
				Title:      fmt.Sprintf("%s list", profile.Name),
				Name:       fmt.Sprintf("%s list", profile.Name),
				Group:      "NoSql Default",
				GroupTitle: "NoSql Default",
				Version:    "0.0.0",
				Parameter: Parameter{
					Fields: Fields{
						FieldsParameter: fieldsParameter,
					},
				},
				Success: Success{
					FieldsSuccess: FieldsSuccess{
						FieldsSuccess: fieldsSuccess,
					},
				},
			})
		} else if api == "update" {
			docs = append(docs, Doc{
				Type:       "put",
				URL:        routePrefixs.Update(profile.Name),
				Title:      fmt.Sprintf("%s update", profile.Name),
				Name:       fmt.Sprintf("%s update", profile.Name),
				Group:      "NoSql Default",
				GroupTitle: "NoSql Default",
				Version:    "0.0.0",
				Success: Success{
					FieldsSuccess: FieldsSuccess{
						FieldsSuccess: []FieldsParameter{},
					},
				},
				Parameter: Parameter{
					Fields: Fields{
						FieldsParameter: fieldsParameter,
					},
				},
			})
		} else if api == "create" {
			docs = append(docs, Doc{
				Type:       "post",
				URL:        routePrefixs.Create(profile.Name),
				Title:      fmt.Sprintf("%s create", profile.Name),
				Name:       fmt.Sprintf("%s create", profile.Name),
				Group:      "NoSql Default",
				GroupTitle: "NoSql Default",
				Version:    "0.0.0",
				Success: Success{
					FieldsSuccess: FieldsSuccess{
						FieldsSuccess: []FieldsParameter{},
					},
				},
				Parameter: Parameter{
					Fields: Fields{
						FieldsParameter: fieldsParameter,
					},
				},
			})
		} else if api == "delete" {
			docs = append(docs, Doc{
				Type:       "put",
				URL:        routePrefixs.Delete(profile.Name),
				Title:      fmt.Sprintf("%s delete", profile.Name),
				Name:       fmt.Sprintf("%s delete", profile.Name),
				Group:      "NoSql Default",
				GroupTitle: "NoSql Default",
				Version:    "0.0.0",
				Success: Success{
					FieldsSuccess: FieldsSuccess{
						FieldsSuccess: []FieldsParameter{},
					},
				},
				Parameter: Parameter{
					Fields: Fields{
						FieldsParameter: fieldsParameter,
					},
				},
			})
		}
	})

	docs = funk.Map(docs, func(doc Doc) Doc {
		subs := funk.Map(strings.Split(strings.ReplaceAll(doc.URL, ":", ""), "/"), func(sub string) string {
			return strings.Title(sub)
		}).([]string)
		doc.Name = strings.Title(doc.Type) + strings.Join(subs, "")
		return doc
	}).([]Doc)
	return &docs
}
