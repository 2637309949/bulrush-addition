// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgoext

import (
	"encoding/json"
	"net/url"
	"reflect"
	"strings"

	"github.com/globalsign/mgo/bson"
)

// Query defined query std
type Query struct {
	Query struct {
		Cond     string `form:"cond" json:"cond" xml:"cond"`
		Project  string `form:"project" json:"project" xml:"project"`
		UProject string `form:"uproject" json:"uproject" xml:"uproject"`
		Preload  string `form:"preload" json:"preload" xml:"preload"`
		Order    string `form:"order" json:"order" xml:"order"`
		Page     int    `form:"page" json:"page" xml:"page"`
		Size     int    `form:"size" json:"size" xml:"size"`
		Range    string `form:"range" json:"range" xml:"range"`
	}
	name     string
	model    interface{}
	Cond     map[string]interface{}
	Pipe     []map[string]interface{}
	Project  map[string]interface{}
	UProject map[string]interface{}
	Preload  string
	Order    map[string]interface{}
	Page     int
	Size     int
	Range    string
}

// PreloadInfo defined preload unit
type PreloadInfo struct {
	BsonName string
	IsArray  bool
	Coll     string
	Local    string
	Foreign  string
	Up       string
}

// NewQuery defined return query with default
func NewQuery() *Query {
	return &Query{
		Query: struct {
			Cond     string `form:"cond" json:"cond" xml:"cond"`
			Project  string `form:"project" json:"project" xml:"project"`
			UProject string `form:"uproject" json:"uproject" xml:"uproject"`
			Preload  string `form:"preload" json:"preload" xml:"preload"`
			Order    string `form:"order" json:"order" xml:"order"`
			Page     int    `form:"page" json:"page" xml:"page"`
			Size     int    `form:"size" json:"size" xml:"size"`
			Range    string `form:"range" json:"range" xml:"range"`
		}{
			Cond:  "%7B%7D",
			Size:  20,
			Page:  1,
			Range: "PAGE",
			Order: "-_created",
		},
		Page:  1,
		Size:  20,
		Range: "PAGE",
		Order: map[string]interface{}{
			"_created": -1,
		},
		Pipe: []map[string]interface{}{},
	}
}

func jsonDot(str1 string, str2 string) string {
	if str1 == "" {
		return str2
	} else if str2 == "" {
		return str1
	}
	return strings.Join([]string{str1, str2}, ".")
}

func replaceOid(target interface{}) interface{} {
	targetMap, ok := target.(map[string]interface{})
	if ok {
		if v, ok := targetMap["$oid"]; ok {
			return bson.ObjectIdHex(v.(string))
		}
		for k, v := range targetMap {
			targetMap[k] = replaceOid(v)
		}
		return targetMap
	}
	targetSlice, ok := target.([]interface{})
	if ok {
		for k, v := range targetSlice {
			targetSlice[k] = replaceOid(v)
		}
		return targetSlice
	}
	return target
}

func mapKey(mType reflect.Type, path string) string {
	if mType.Kind() == reflect.Ptr || mType.Kind() == reflect.Slice {
		mType = mType.Elem()
	}
	if path == "" {
		return ""
	}
	first := strings.Split(path, ".")[0]
	remain := strings.Join(strings.Split(path, ".")[1:], ".")
	if field := findFieldStruct(mType, first); field != nil {
		bsonName := strings.Split(field.Tag.Get("bson"), ",")[0]
		return jsonDot(bsonName, mapKey(field.Type, remain))
	}
	return jsonDot(first, mapKey(mType, remain))
}

func replaceKey(model interface{}, target interface{}, path string) interface{} {
	targetMap, ok := target.(map[string]interface{})
	if ok {
		for k, v := range targetMap {
			if !strings.Contains(k, "$") {
				delete(targetMap, k)
				k = mapKey(reflect.TypeOf(model), k)
			}
			targetMap[k] = replaceKey(model, v, jsonDot(path, k))
		}
		return targetMap
	}
	targetSlice, ok := target.([]interface{})
	if ok {
		for k, v := range targetSlice {
			targetSlice[k] = replaceKey(model, v, jsonDot(path, ""))
		}
		return targetSlice
	}
	return target
}

// Build defined all build
func (q *Query) Build(cond map[string]interface{}) error {
	q.Page = q.Query.Page
	q.Size = q.Query.Size
	q.Range = q.Query.Range
	q.Preload = q.Query.Preload

	err := q.BuildCond(cond)
	err = q.BuildPipe()
	return err
}

// BuildCond defined select sql
func (q *Query) BuildCond(cond map[string]interface{}) error {
	q.Cond = cond
	unescapeWhere, err := url.QueryUnescape(q.Query.Cond)
	if err != nil {
		return err
	}
	if err = json.Unmarshal([]byte(unescapeWhere), &q.Cond); err != nil {
		return err
	}
	q.Cond = replaceOid(q.Cond).(map[string]interface{})
	q.Cond = replaceKey(q.model, q.Cond, "").(map[string]interface{})
	return nil
}

// BuildPipe defined pipe array
func (q *Query) BuildPipe() error {
	q.Pipe = append(q.Pipe, map[string]interface{}{
		"$match": q.Cond,
	})
	related := q.BuildRelated()
	if related != nil {
		q.Pipe = append(q.Pipe, related...)
	}
	q.BuildOrder()
	if len(q.Order) > 0 {
		q.Pipe = append(q.Pipe, map[string]interface{}{
			"$sort": q.Order,
		})
	}
	q.BuildProject()
	if len(q.Project) > 0 {
		q.Pipe = append(q.Pipe, map[string]interface{}{
			"$project": q.Project,
		})
	}
	q.BuildUProject()
	if len(q.UProject) > 0 {
		q.Pipe = append(q.Pipe, map[string]interface{}{
			"$project": q.UProject,
		})
	}
	q.Pipe = append(q.Pipe, map[string]interface{}{
		"$skip": (q.Page - 1) * q.Size,
	})
	q.Pipe = append(q.Pipe, map[string]interface{}{
		"$limit": q.Size,
	})
	return nil
}

// BuildOrder defined order sql
func (q *Query) BuildOrder() {
	if q.Query.Order != "" {
		orderMap := map[string]interface{}{}
		orders := strings.Split(q.Query.Order, ",")
		for _, item := range orders {
			if strings.HasPrefix(item, "-") {
				orderMap[item] = -1
			} else {
				orderMap[item] = 1
			}
		}
		q.Order = orderMap
	}
}

// BuildRelated defined related sql for preLoad
func (q *Query) BuildRelated() []map[string]interface{} {
	if q.Query.Preload != "" {
		relatedArr := []map[string]interface{}{}
		related := strings.Split(q.Query.Preload, ",")
		for _, item := range related {
			preInfo := preloadInfo(q.model, item)
			if preInfo != nil {
				if preInfo.Up != "" {
					q.Query.UProject = strings.Join([]string{q.Query.UProject, preInfo.Up}, ",")
				}
				if !preInfo.IsArray {
					relatedArr = append(relatedArr, []map[string]interface{}{
						map[string]interface{}{
							"$lookup": map[string]interface{}{
								"from":         preInfo.Coll,
								"localField":   preInfo.Local,
								"foreignField": preInfo.Foreign,
								"as":           preInfo.BsonName,
							},
						},
						map[string]interface{}{
							"$unwind": map[string]interface{}{
								"path":                       "$" + preInfo.BsonName,
								"preserveNullAndEmptyArrays": true,
							},
						},
					}...)
				} else {
					group := map[string]interface{}{"_id": "$_id", "root": map[string]interface{}{"$first": "$$ROOT"}}
					group[preInfo.BsonName] = map[string]interface{}{"$push": "$" + preInfo.BsonName}
					addFields := map[string]interface{}{}
					addFields["root."+preInfo.BsonName] = "$" + preInfo.BsonName
					unwindLocal := map[string]interface{}{
						"path":                       "$" + preInfo.Local,
						"preserveNullAndEmptyArrays": true,
					}
					lookup := map[string]interface{}{
						"from":         preInfo.Coll,
						"localField":   preInfo.Local,
						"foreignField": preInfo.Foreign,
						"as":           preInfo.BsonName,
					}
					unwind := map[string]interface{}{
						"path":                       "$" + preInfo.BsonName,
						"preserveNullAndEmptyArrays": true,
					}
					replaceRoot := map[string]interface{}{
						"newRoot": "$root",
					}
					relatedArr = append(relatedArr, []map[string]interface{}{
						map[string]interface{}{
							"$unwind": unwindLocal,
						},
						map[string]interface{}{
							"$lookup": lookup,
						},
						map[string]interface{}{
							"$unwind": unwind,
						},
						map[string]interface{}{
							"$group": group,
						},
						map[string]interface{}{
							"$addFields": addFields,
						},
						map[string]interface{}{
							"$replaceRoot": replaceRoot,
						},
					}...)
				}
			}
		}
		return relatedArr
	}
	return nil
}

// BuildProject dfined build select fields
func (q *Query) BuildProject() {
	Project := map[string]interface{}{}
	if q.Query.Project != "" {
		projects := strings.Split(q.Query.Project, ",")
		for _, item := range projects {
			if item != "" {
				Project[item] = 1
			}
		}
		q.Project = Project
	}
}

// BuildUProject dfined build select fields
func (q *Query) BuildUProject() {
	UProject := map[string]interface{}{}
	if q.Query.UProject != "" {
		projects := strings.Split(q.Query.UProject, ",")
		for _, item := range projects {
			if item != "" {
				UProject[item] = 0
			}
		}
		q.UProject = UProject
	}
}
