// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package template

import (
	"log"
	"os"
	"path"
)

var tr_js = []byte("\x64\x65\x66\x69\x6e\x65\x28\x7b\x0a\x20\x20\x20\x20\x74\x72\x3a\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x41\x6c\x6c\x6f\x77\x65\x64\x20\x76\x61\x6c\x75\x65\x73\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xc4\xb0\x7a\x69\x6e\x20\x76\x65\x72\x69\x6c\x65\x6e\x20\x64\x65\xc4\x9f\x65\x72\x6c\x65\x72\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x43\x6f\x6d\x70\x61\x72\x65\x20\x61\x6c\x6c\x20\x77\x69\x74\x68\x20\x70\x72\x65\x64\x65\x63\x65\x73\x73\x6f\x72\x27\x3a\x20\x27\x54\xc3\xbc\x6d\xc3\xbc\x6e\xc3\xbc\x20\xc3\xb6\x6e\x63\x65\x6b\x69\x6c\x65\x72\x20\x69\x6c\x65\x20\x6b\x61\x72\xc5\x9f\xc4\xb1\x6c\x61\xc5\x9f\x74\xc4\xb1\x72\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x63\x6f\x6d\x70\x61\x72\x65\x20\x63\x68\x61\x6e\x67\x65\x73\x20\x74\x6f\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x64\x65\xc4\x9f\x69\xc5\x9f\x69\x6b\x6c\x69\x6b\x6c\x65\x72\x69\x20\x6b\x61\x72\xc5\x9f\xc4\xb1\x6c\x61\xc5\x9f\x74\xc4\xb1\x72\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x63\x6f\x6d\x70\x61\x72\x65\x64\x20\x74\x6f\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x6b\x61\x72\xc5\x9f\xc4\xb1\x6c\x61\xc5\x9f\x74\xc4\xb1\x72\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x44\x65\x66\x61\x75\x6c\x74\x20\x76\x61\x6c\x75\x65\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x56\x61\x72\x73\x61\x79\xc4\xb1\x6c\x61\x6e\x20\x64\x65\xc4\x9f\x65\x72\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x44\x65\x73\x63\x72\x69\x70\x74\x69\x6f\x6e\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x41\xc3\xa7\xc4\xb1\x6b\x6c\x61\x6d\x61\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x46\x69\x65\x6c\x64\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x41\x6c\x61\x6e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x47\x65\x6e\x65\x72\x61\x6c\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x47\x65\x6e\x65\x6c\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x47\x65\x6e\x65\x72\x61\x74\x65\x64\x20\x77\x69\x74\x68\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x4f\x6c\x75\xc5\x9f\x74\x75\x72\x61\x6e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x4e\x61\x6d\x65\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xc4\xb0\x73\x69\x6d\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x4e\x6f\x20\x72\x65\x73\x70\x6f\x6e\x73\x65\x20\x76\x61\x6c\x75\x65\x73\x2e\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x44\xc3\xb6\x6e\xc3\xbc\xc5\x9f\x20\x76\x65\x72\x69\x73\x69\x20\x79\x6f\x6b\x2e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x6f\x70\x74\x69\x6f\x6e\x61\x6c\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x6f\x70\x73\x69\x79\x6f\x6e\x65\x6c\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x50\x61\x72\x61\x6d\x65\x74\x65\x72\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x50\x61\x72\x61\x6d\x65\x74\x72\x65\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x50\x65\x72\x6d\x69\x73\x73\x69\x6f\x6e\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xc4\xb0\x7a\x69\x6e\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x52\x65\x73\x70\x6f\x6e\x73\x65\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x44\xc3\xb6\x6e\xc3\xbc\xc5\x9f\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x53\x65\x6e\x64\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x47\xc3\xb6\x6e\x64\x65\x72\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x53\x65\x6e\x64\x20\x61\x20\x53\x61\x6d\x70\x6c\x65\x20\x52\x65\x71\x75\x65\x73\x74\x27\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xc3\x96\x72\x6e\x65\x6b\x20\x69\x73\x74\x65\x6b\x20\x67\xc3\xb6\x6e\x64\x65\x72\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x73\x68\x6f\x77\x20\x75\x70\x20\x74\x6f\x20\x76\x65\x72\x73\x69\x6f\x6e\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x62\x75\x20\x76\x65\x72\x73\x69\x79\x6f\x6e\x61\x20\x6b\x61\x64\x61\x72\x20\x67\xc3\xb6\x73\x74\x65\x72\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x53\x69\x7a\x65\x20\x72\x61\x6e\x67\x65\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x42\x6f\x79\x75\x74\x20\x61\x72\x61\x6c\xc4\xb1\xc4\x9f\xc4\xb1\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x54\x79\x70\x65\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x54\x69\x70\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x75\x72\x6c\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x75\x72\x6c\x27\x0a\x20\x20\x20\x20\x7d\x0a\x7d\x29\x3b\x0a")

func init() {
	filepath := "/locales/tr.js"
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
	_, err = f.Write(tr_js)
	if err != nil {
		log.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
