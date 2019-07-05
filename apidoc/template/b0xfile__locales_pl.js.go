// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package template

import (
	"log"
	"os"
	"path"
)

var pl_js = []byte("\x64\x65\x66\x69\x6e\x65\x28\x7b\x0a\x20\x20\x20\x20\x70\x6c\x3a\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x41\x6c\x6c\x6f\x77\x65\x64\x20\x76\x61\x6c\x75\x65\x73\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x44\x6f\x7a\x77\x6f\x6c\x6f\x6e\x65\x20\x77\x61\x72\x74\x6f\xc5\x9b\x63\x69\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x43\x6f\x6d\x70\x61\x72\x65\x20\x61\x6c\x6c\x20\x77\x69\x74\x68\x20\x70\x72\x65\x64\x65\x63\x65\x73\x73\x6f\x72\x27\x3a\x20\x27\x50\x6f\x72\xc3\xb3\x77\x6e\x61\x6a\x20\x7a\x20\x70\x6f\x70\x72\x7a\x65\x64\x6e\x69\x6d\x69\x20\x77\x65\x72\x73\x6a\x61\x6d\x69\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x63\x6f\x6d\x70\x61\x72\x65\x20\x63\x68\x61\x6e\x67\x65\x73\x20\x74\x6f\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x70\x6f\x72\xc3\xb3\x77\x6e\x61\x6a\x20\x7a\x6d\x69\x61\x6e\x79\x20\x64\x6f\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x63\x6f\x6d\x70\x61\x72\x65\x64\x20\x74\x6f\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x70\x6f\x72\xc3\xb3\x77\x6e\x61\x6a\x20\x64\x6f\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x44\x65\x66\x61\x75\x6c\x74\x20\x76\x61\x6c\x75\x65\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x57\x61\x72\x74\x6f\xc5\x9b\xc4\x87\x20\x64\x6f\x6d\x79\xc5\x9b\x6c\x6e\x61\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x44\x65\x73\x63\x72\x69\x70\x74\x69\x6f\x6e\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x4f\x70\x69\x73\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x46\x69\x65\x6c\x64\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x50\x6f\x6c\x65\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x47\x65\x6e\x65\x72\x61\x6c\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x47\x65\x6e\x65\x72\x61\x6c\x6e\x69\x65\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x47\x65\x6e\x65\x72\x61\x74\x65\x64\x20\x77\x69\x74\x68\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x57\x79\x67\x65\x6e\x65\x72\x6f\x77\x61\x6e\x6f\x20\x7a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x4e\x61\x6d\x65\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x4e\x61\x7a\x77\x61\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x4e\x6f\x20\x72\x65\x73\x70\x6f\x6e\x73\x65\x20\x76\x61\x6c\x75\x65\x73\x2e\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x42\x72\x61\x6b\x20\x6f\x64\x70\x6f\x77\x69\x65\x64\x7a\x69\x2e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x6f\x70\x74\x69\x6f\x6e\x61\x6c\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x6f\x70\x63\x6a\x6f\x6e\x61\x6c\x6e\x79\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x50\x61\x72\x61\x6d\x65\x74\x65\x72\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x50\x61\x72\x61\x6d\x65\x74\x72\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x50\x65\x72\x6d\x69\x73\x73\x69\x6f\x6e\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x55\x70\x72\x61\x77\x6e\x69\x65\x6e\x69\x61\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x52\x65\x73\x70\x6f\x6e\x73\x65\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x4f\x64\x70\x6f\x77\x69\x65\x64\xc5\xba\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x53\x65\x6e\x64\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x57\x79\xc5\x9b\x6c\x69\x6a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x53\x65\x6e\x64\x20\x61\x20\x53\x61\x6d\x70\x6c\x65\x20\x52\x65\x71\x75\x65\x73\x74\x27\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x57\x79\xc5\x9b\x6c\x69\x6a\x20\x70\x72\x7a\x79\x6b\xc5\x82\x61\x64\x6f\x77\x65\x20\xc5\xbc\xc4\x85\x64\x61\x6e\x69\x65\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x73\x68\x6f\x77\x20\x75\x70\x20\x74\x6f\x20\x76\x65\x72\x73\x69\x6f\x6e\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x70\x6f\x6b\x61\xc5\xbc\x20\x64\x6f\x20\x77\x65\x72\x73\x6a\x69\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x53\x69\x7a\x65\x20\x72\x61\x6e\x67\x65\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x5a\x61\x6b\x72\x65\x73\x20\x72\x6f\x7a\x6d\x69\x61\x72\x75\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x54\x79\x70\x65\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x54\x79\x70\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x75\x72\x6c\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x75\x72\x6c\x27\x0a\x20\x20\x20\x20\x7d\x0a\x7d\x29\x3b\x0a")

func init() {
	filepath := "/locales/pl.js"
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
	_, err = f.Write(pl_js)
	if err != nil {
		log.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}