// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package template

import (
	"log"
	"os"
	"path"
)

var pt_br_js = []byte("\x64\x65\x66\x69\x6e\x65\x28\x7b\x0a\x20\x20\x20\x20\x27\x70\x74\x5f\x62\x72\x27\x3a\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x41\x6c\x6c\x6f\x77\x65\x64\x20\x76\x61\x6c\x75\x65\x73\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x56\x61\x6c\x6f\x72\x65\x73\x20\x70\x65\x72\x6d\x69\x74\x69\x64\x6f\x73\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x43\x6f\x6d\x70\x61\x72\x65\x20\x61\x6c\x6c\x20\x77\x69\x74\x68\x20\x70\x72\x65\x64\x65\x63\x65\x73\x73\x6f\x72\x27\x3a\x20\x27\x43\x6f\x6d\x70\x61\x72\x65\x20\x74\x6f\x64\x6f\x73\x20\x63\x6f\x6d\x20\x61\x6e\x74\x65\x63\x65\x73\x73\x6f\x72\x65\x73\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x63\x6f\x6d\x70\x61\x72\x65\x20\x63\x68\x61\x6e\x67\x65\x73\x20\x74\x6f\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x63\x6f\x6d\x70\x61\x72\x61\x72\x20\x61\x6c\x74\x65\x72\x61\xc3\xa7\xc3\xb5\x65\x73\x20\x63\x6f\x6d\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x63\x6f\x6d\x70\x61\x72\x65\x64\x20\x74\x6f\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x63\x6f\x6d\x70\x61\x72\x61\x64\x6f\x20\x63\x6f\x6d\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x44\x65\x66\x61\x75\x6c\x74\x20\x76\x61\x6c\x75\x65\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x56\x61\x6c\x6f\x72\x20\x70\x61\x64\x72\xc3\xa3\x6f\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x44\x65\x73\x63\x72\x69\x70\x74\x69\x6f\x6e\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x44\x65\x73\x63\x72\x69\xc3\xa7\xc3\xa3\x6f\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x46\x69\x65\x6c\x64\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x43\x61\x6d\x70\x6f\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x47\x65\x6e\x65\x72\x61\x6c\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x47\x65\x72\x61\x6c\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x47\x65\x6e\x65\x72\x61\x74\x65\x64\x20\x77\x69\x74\x68\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x47\x65\x72\x61\x64\x6f\x20\x63\x6f\x6d\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x4e\x61\x6d\x65\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x4e\x6f\x6d\x65\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x4e\x6f\x20\x72\x65\x73\x70\x6f\x6e\x73\x65\x20\x76\x61\x6c\x75\x65\x73\x2e\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x53\x65\x6d\x20\x76\x61\x6c\x6f\x72\x65\x73\x20\x64\x65\x20\x72\x65\x73\x70\x6f\x73\x74\x61\x2e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x6f\x70\x74\x69\x6f\x6e\x61\x6c\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x6f\x70\x63\x69\x6f\x6e\x61\x6c\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x50\x61\x72\x61\x6d\x65\x74\x65\x72\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x50\x61\x72\xc3\xa2\x6d\x65\x74\x72\x6f\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x50\x65\x72\x6d\x69\x73\x73\x69\x6f\x6e\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x50\x65\x72\x6d\x69\x73\x73\xc3\xa3\x6f\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x52\x65\x73\x70\x6f\x6e\x73\x65\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x52\x65\x73\x70\x6f\x73\x74\x61\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x53\x65\x6e\x64\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x45\x6e\x76\x69\x61\x72\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x53\x65\x6e\x64\x20\x61\x20\x53\x61\x6d\x70\x6c\x65\x20\x52\x65\x71\x75\x65\x73\x74\x27\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x45\x6e\x76\x69\x61\x72\x20\x75\x6d\x20\x45\x78\x65\x6d\x70\x6c\x6f\x20\x64\x65\x20\x50\x65\x64\x69\x64\x6f\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x73\x68\x6f\x77\x20\x75\x70\x20\x74\x6f\x20\x76\x65\x72\x73\x69\x6f\x6e\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x61\x70\x61\x72\x65\x63\x65\x72\x20\x70\x61\x72\x61\x20\x61\x20\x76\x65\x72\x73\xc3\xa3\x6f\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x53\x69\x7a\x65\x20\x72\x61\x6e\x67\x65\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x46\x61\x69\x78\x61\x20\x64\x65\x20\x74\x61\x6d\x61\x6e\x68\x6f\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x54\x79\x70\x65\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x54\x69\x70\x6f\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x75\x72\x6c\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x75\x72\x6c\x27\x0a\x20\x20\x20\x20\x7d\x0a\x7d\x29\x3b\x0a")

func init() {
	filepath := "/locales/pt_br.js"
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
	_, err = f.Write(pt_br_js)
	if err != nil {
		log.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
