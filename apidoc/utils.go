// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package apidoc

import "regexp"

// APIData defined APIData
var APIData = "api_data.js"

// APIDataSys defined APIData
var APIDataSys = "api_data_sys.js"

// APIProject defined APIProject
var APIProject = "api_project.js"

func findStringSubmatch(matcher string, s string) []string {
	var rgx = regexp.MustCompile(matcher)
	rs := rgx.FindStringSubmatch(s)
	if rs != nil {
		return rs[1:]
	}
	return []string{}
}
