// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgoext

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"strings"

	addition "github.com/2637309949/bulrush-addition"
	"github.com/globalsign/mgo/bson"
)

// Query defined query std
type Query struct {
	name     string
	model    interface{}
	Cond     string `form:"cond" json:"cond" xml:"cond"`
	CondMap  map[string]interface{}
	Project  string `form:"project" json:"project" xml:"project"`
	UProject string `form:"uproject" json:"uproject" xml:"uproject"`
	Related  string `form:"related" json:"related" xml:"related"`
	Order    string `form:"order" json:"order" xml:"order"`
	Page     int    `form:"page" json:"page" xml:"page"`
	Size     int    `form:"size" json:"size" xml:"size"`
	Range    string `form:"range" json:"range" xml:"range"`
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
		Page:  1,
		Size:  20,
		Range: "PAGE",
		Order: "_created",
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

// BuildCond defined select sql
func (q *Query) BuildCond(preCond ...map[string]interface{}) (map[string]interface{}, error) {
	cond := map[string]interface{}{}
	if len(preCond) > 0 {
		cond = preCond[0]
	}
	var cloneCond map[string]interface{}
	if q.Cond == "" {
		q.Cond = "%7B%7D"
	}
	unescapeWhere, err := url.QueryUnescape(q.Cond)
	if err != nil {
		return map[string]interface{}{}, err
	}
	fmt.Println("===")

	err = json.Unmarshal([]byte(unescapeWhere), &cond)
	if err != nil {
		return map[string]interface{}{}, err
	}
	err = addition.CopyMap(cond, &cloneCond)
	if err != nil {
		return map[string]interface{}{}, err
	}
	cond = replaceOid(cond).(map[string]interface{})
	cond = replaceKey(q.model, cond, "").(map[string]interface{})
	fmt.Println(cond)
	q.CondMap = cloneCond
	return cond, nil
}

// BuildPipe defined pipe array
func (q *Query) BuildPipe(preCond ...map[string]interface{}) ([]map[string]interface{}, error) {
	pipe := []map[string]interface{}{}
	match, err := q.BuildCond(preCond...)
	if err != nil {
		return []map[string]interface{}{}, err
	}
	pipe = append(pipe, map[string]interface{}{
		"$match": match,
	})
	related := q.BuildRelated()
	if related != nil {
		pipe = append(pipe, related...)
	}
	order := q.BuildOrder()
	if order != nil {
		pipe = append(pipe, map[string]interface{}{
			"$sort": order,
		})
	}
	project := q.BuildProject()
	if project != nil {
		pipe = append(pipe, map[string]interface{}{
			"$project": project,
		})
	}
	uProject := q.BuildUProject()
	if uProject != nil {
		pipe = append(pipe, map[string]interface{}{
			"$project": uProject,
		})
	}
	return pipe, nil
}

// BuildPipeWithPage defined pipe array
func (q *Query) BuildPipeWithPage(preCond ...map[string]interface{}) ([]map[string]interface{}, error) {
	pipe, err := q.BuildPipe(preCond...)
	pipe = append(pipe, map[string]interface{}{
		"$skip": (q.Page - 1) * q.Size,
	})
	pipe = append(pipe, map[string]interface{}{
		"$limit": q.Size,
	})
	return pipe, err
}

// BuildOrder defined order sql
func (q *Query) BuildOrder() map[string]interface{} {
	if q.Order != "" {
		orderMap := map[string]interface{}{}
		orders := strings.Split(q.Order, ",")
		for _, item := range orders {
			if strings.HasPrefix(item, "-") {
				orderMap[item] = -1
			} else {
				orderMap[item] = 1
			}
		}
		return orderMap
	}
	return nil
}

// BuildRelated defined related sql for preLoad
func (q *Query) BuildRelated() []map[string]interface{} {
	if q.Related != "" {
		relatedArr := []map[string]interface{}{}
		related := strings.Split(q.Related, ",")
		for _, item := range related {
			preInfo := preloadInfo(q.model, item)
			if preInfo != nil {
				if preInfo.Up != "" {
					q.UProject = strings.Join([]string{q.UProject, preInfo.Up}, ",")
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
					relatedArr = append(relatedArr, []map[string]interface{}{
						map[string]interface{}{
							"$unwind": map[string]interface{}{
								"path":                       "$" + preInfo.Local,
								"preserveNullAndEmptyArrays": true,
							},
						},
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
						map[string]interface{}{
							"$group": group,
						},
						map[string]interface{}{
							"$addFields": addFields,
						},
						map[string]interface{}{
							"$replaceRoot": map[string]interface{}{
								"newRoot": "$root",
							},
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
func (q *Query) BuildProject() map[string]interface{} {
	projectMap := map[string]interface{}{}
	if q.Project != "" {
		projects := strings.Split(q.Project, ",")
		for _, item := range projects {
			if item != "" {
				projectMap[item] = 1
			}
		}
		return projectMap
	}
	return nil
}

// BuildUProject dfined build select fields
func (q *Query) BuildUProject() map[string]interface{} {
	projectMap := map[string]interface{}{}
	if q.UProject != "" {
		projects := strings.Split(q.UProject, ",")
		for _, item := range projects {
			if item != "" {
				projectMap[item] = 0
			}
		}
		return projectMap
	}
	return nil
}
