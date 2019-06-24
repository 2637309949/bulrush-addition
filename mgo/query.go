// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mgo

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/thoas/go-funk"

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

// NewQuery defined return query with default
func NewQuery() *Query {
	return &Query{
		Page:  1,
		Size:  20,
		Range: "PAGE",
		Order: "_created",
	}
}

func replaceOid(target interface{}) interface{} {
	targetOid, ok := target.(map[string]interface{})
	if ok {
		v, ok := targetOid["$oid"]
		if ok {
			return bson.ObjectIdHex(v.(string))
		}
	}

	targetSlice, ok := target.([]map[string]interface{})
	if ok {
		for k, v := range targetSlice {
			targetSlice[k] = replaceOid(v).(map[string]interface{})
		}
	}

	targetMap, ok := target.(map[string]interface{})
	if ok {
		for k, v := range targetMap {
			targetMap[k] = replaceOid(v)
		}
	}
	return targetMap
}

// BuildCond defined select sql
func (q *Query) BuildCond(id string) (map[string]interface{}, error) {
	cond := map[string]interface{}{}
	if id != "" {
		cond["_id"] = map[string]interface{}{"$oid": id}
	}
	var cloneCond map[string]interface{}
	if q.Cond == "" {
		q.Cond = "%7B%7D"
	}
	unescapeWhere, err := url.QueryUnescape(q.Cond)
	if err != nil {
		return map[string]interface{}{}, err
	}
	err = json.Unmarshal([]byte(unescapeWhere), &cond)
	if err != nil {
		return map[string]interface{}{}, err
	}
	err = addition.CopyMap(cond, &cloneCond)
	if err != nil {
		return map[string]interface{}{}, err
	}
	cond = replaceOid(cond).(map[string]interface{})
	q.CondMap = cloneCond
	return cond, nil
}

// BuildPipe defined pipe array
func (q *Query) BuildPipe(id string) ([]map[string]interface{}, error) {
	pipe := []map[string]interface{}{}
	match, err := q.BuildCond(id)
	if err != nil {
		return []map[string]interface{}{}, err
	}
	if q.Range != "ALL" {
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
	}
	return pipe, nil
}

// BuildPipeWithPage defined pipe array
func (q *Query) BuildPipeWithPage(id string) ([]map[string]interface{}, error) {
	pipe, err := q.BuildPipe(id)
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
			tag := fieldTag(q.model, item, "ref")
			tagArr := strings.Split(tag, ";")
			refModel := tagArr[0]
			refUProject := strings.Join(findStringSubmatch(`up\((.*?)\)`, tag), ",")
			if refUProject != "" {
				refUProjectArr := strings.Split(refUProject, ",")
				refUProjectArr = funk.Map(refUProjectArr, func(x string) string {
					return fmt.Sprintf("%s.%s", item, x)
				}).([]string)
				refUProject = strings.Join(refUProjectArr, ",")
				q.UProject = refUProject
			}
			if tag != "" {
				uProject := q.BuildUProject()
				relatedArr = append(relatedArr, []map[string]interface{}{
					map[string]interface{}{
						"$lookup": map[string]interface{}{
							"from":         refModel,
							"localField":   item,
							"foreignField": "_id",
							"as":           item,
						},
					},
					map[string]interface{}{
						"$unwind": map[string]interface{}{
							"path":                       "$" + item,
							"preserveNullAndEmptyArrays": true,
						},
					},
				}...)
				if uProject != nil {
					relatedArr = append(relatedArr, map[string]interface{}{
						"$project": uProject,
					})
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
