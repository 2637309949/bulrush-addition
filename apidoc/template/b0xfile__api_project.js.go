// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package template

import (
	"log"
	"os"
	"path"
)

var api_project_js = []byte("\x64\x65\x66\x69\x6e\x65\x28\x7b\x0a\x20\x20\x22\x6e\x61\x6d\x65\x22\x3a\x20\x22\x62\x75\x6c\x72\x75\x73\x68\x2d\x74\x65\x6d\x70\x6c\x61\x74\x65\x22\x2c\x0a\x20\x20\x22\x76\x65\x72\x73\x69\x6f\x6e\x22\x3a\x20\x22\x31\x2e\x30\x22\x2c\x0a\x20\x20\x22\x64\x65\x73\x63\x72\x69\x70\x74\x69\x6f\x6e\x22\x3a\x20\x22\x61\x70\x69\x20\x64\x6f\x63\x75\x6d\x65\x6e\x74\x61\x74\x69\x6f\x6e\x22\x2c\x0a\x20\x20\x22\x74\x69\x74\x6c\x65\x22\x3a\x20\x22\x62\x75\x6c\x72\x75\x73\x68\x2d\x74\x65\x6d\x70\x6c\x61\x74\x65\x20\x2d\x20\x61\x70\x69\x20\x64\x6f\x63\x75\x6d\x65\x6e\x74\x61\x74\x69\x6f\x6e\x22\x2c\x0a\x20\x20\x22\x75\x72\x6c\x22\x3a\x20\x22\x68\x74\x74\x70\x3a\x2f\x2f\x31\x32\x37\x2e\x30\x2e\x30\x2e\x31\x3a\x38\x30\x38\x30\x2f\x61\x70\x69\x2f\x76\x31\x22\x2c\x0a\x20\x20\x22\x73\x61\x6d\x70\x6c\x65\x55\x72\x6c\x22\x3a\x20\x22\x68\x74\x74\x70\x3a\x2f\x2f\x31\x32\x37\x2e\x30\x2e\x30\x2e\x31\x3a\x38\x30\x38\x30\x2f\x61\x70\x69\x2f\x76\x31\x22\x2c\x0a\x20\x20\x22\x74\x65\x6d\x70\x6c\x61\x74\x65\x22\x3a\x20\x7b\x0a\x20\x20\x20\x20\x22\x66\x6f\x72\x63\x65\x4c\x61\x6e\x67\x75\x61\x67\x65\x22\x3a\x20\x22\x7a\x68\x5f\x63\x6e\x22\x2c\x0a\x20\x20\x20\x20\x22\x77\x69\x74\x68\x43\x6f\x6d\x70\x61\x72\x65\x22\x3a\x20\x66\x61\x6c\x73\x65\x2c\x0a\x20\x20\x20\x20\x22\x77\x69\x74\x68\x47\x65\x6e\x65\x72\x61\x74\x6f\x72\x22\x3a\x20\x66\x61\x6c\x73\x65\x0a\x20\x20\x7d\x2c\x0a\x20\x20\x22\x64\x65\x66\x61\x75\x6c\x74\x56\x65\x72\x73\x69\x6f\x6e\x22\x3a\x20\x22\x30\x2e\x30\x2e\x30\x22\x2c\x0a\x20\x20\x22\x61\x70\x69\x64\x6f\x63\x22\x3a\x20\x22\x30\x2e\x33\x2e\x30\x22\x2c\x0a\x20\x20\x22\x67\x65\x6e\x65\x72\x61\x74\x6f\x72\x22\x3a\x20\x7b\x0a\x20\x20\x20\x20\x22\x6e\x61\x6d\x65\x22\x3a\x20\x22\x61\x70\x69\x64\x6f\x63\x22\x2c\x0a\x20\x20\x20\x20\x22\x74\x69\x6d\x65\x22\x3a\x20\x22\x32\x30\x31\x39\x2d\x30\x36\x2d\x31\x33\x54\x30\x36\x3a\x34\x34\x3a\x30\x34\x2e\x30\x39\x38\x5a\x22\x2c\x0a\x20\x20\x20\x20\x22\x75\x72\x6c\x22\x3a\x20\x22\x68\x74\x74\x70\x3a\x2f\x2f\x61\x70\x69\x64\x6f\x63\x6a\x73\x2e\x63\x6f\x6d\x22\x2c\x0a\x20\x20\x20\x20\x22\x76\x65\x72\x73\x69\x6f\x6e\x22\x3a\x20\x22\x30\x2e\x31\x37\x2e\x36\x22\x0a\x20\x20\x7d\x0a\x7d\x29\x3b\x0a")

func init() {
	filepath := "/api_project.js"
	if _, err := FS.Stat(CTX, path.Dir(filepath)); os.IsNotExist(err) {
		err = FS.Mkdir(CTX, path.Dir(filepath), 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err := FS.OpenFile(CTX, filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write(api_project_js)
	if err != nil {
		log.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
