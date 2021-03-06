// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package template

import (
	"log"
	"os"
	"path"
)

var locale_js = []byte("\x64\x65\x66\x69\x6e\x65\x28\x5b\x0a\x20\x20\x20\x20\x27\x2e\x2f\x6c\x6f\x63\x61\x6c\x65\x73\x2f\x63\x61\x2e\x6a\x73\x27\x2c\x0a\x20\x20\x20\x20\x27\x2e\x2f\x6c\x6f\x63\x61\x6c\x65\x73\x2f\x64\x65\x2e\x6a\x73\x27\x2c\x0a\x20\x20\x20\x20\x27\x2e\x2f\x6c\x6f\x63\x61\x6c\x65\x73\x2f\x65\x73\x2e\x6a\x73\x27\x2c\x0a\x20\x20\x20\x20\x27\x2e\x2f\x6c\x6f\x63\x61\x6c\x65\x73\x2f\x66\x72\x2e\x6a\x73\x27\x2c\x0a\x20\x20\x20\x20\x27\x2e\x2f\x6c\x6f\x63\x61\x6c\x65\x73\x2f\x69\x74\x2e\x6a\x73\x27\x2c\x0a\x20\x20\x20\x20\x27\x2e\x2f\x6c\x6f\x63\x61\x6c\x65\x73\x2f\x6e\x6c\x2e\x6a\x73\x27\x2c\x0a\x20\x20\x20\x20\x27\x2e\x2f\x6c\x6f\x63\x61\x6c\x65\x73\x2f\x70\x6c\x2e\x6a\x73\x27\x2c\x0a\x20\x20\x20\x20\x27\x2e\x2f\x6c\x6f\x63\x61\x6c\x65\x73\x2f\x70\x74\x5f\x62\x72\x2e\x6a\x73\x27\x2c\x0a\x20\x20\x20\x20\x27\x2e\x2f\x6c\x6f\x63\x61\x6c\x65\x73\x2f\x72\x6f\x2e\x6a\x73\x27\x2c\x0a\x20\x20\x20\x20\x27\x2e\x2f\x6c\x6f\x63\x61\x6c\x65\x73\x2f\x72\x75\x2e\x6a\x73\x27\x2c\x0a\x20\x20\x20\x20\x27\x2e\x2f\x6c\x6f\x63\x61\x6c\x65\x73\x2f\x74\x72\x2e\x6a\x73\x27\x2c\x0a\x20\x20\x20\x20\x27\x2e\x2f\x6c\x6f\x63\x61\x6c\x65\x73\x2f\x76\x69\x2e\x6a\x73\x27\x2c\x0a\x20\x20\x20\x20\x27\x2e\x2f\x6c\x6f\x63\x61\x6c\x65\x73\x2f\x7a\x68\x2e\x6a\x73\x27\x2c\x0a\x20\x20\x20\x20\x27\x2e\x2f\x6c\x6f\x63\x61\x6c\x65\x73\x2f\x7a\x68\x5f\x63\x6e\x2e\x6a\x73\x27\x0a\x5d\x2c\x20\x66\x75\x6e\x63\x74\x69\x6f\x6e\x28\x29\x20\x7b\x0a\x20\x20\x20\x20\x76\x61\x72\x20\x6c\x61\x6e\x67\x49\x64\x20\x3d\x20\x28\x6e\x61\x76\x69\x67\x61\x74\x6f\x72\x2e\x6c\x61\x6e\x67\x75\x61\x67\x65\x20\x7c\x7c\x20\x6e\x61\x76\x69\x67\x61\x74\x6f\x72\x2e\x75\x73\x65\x72\x4c\x61\x6e\x67\x75\x61\x67\x65\x29\x2e\x74\x6f\x4c\x6f\x77\x65\x72\x43\x61\x73\x65\x28\x29\x2e\x72\x65\x70\x6c\x61\x63\x65\x28\x27\x2d\x27\x2c\x20\x27\x5f\x27\x29\x3b\x0a\x20\x20\x20\x20\x76\x61\x72\x20\x6c\x61\x6e\x67\x75\x61\x67\x65\x20\x3d\x20\x6c\x61\x6e\x67\x49\x64\x2e\x73\x75\x62\x73\x74\x72\x28\x30\x2c\x20\x32\x29\x3b\x0a\x20\x20\x20\x20\x76\x61\x72\x20\x6c\x6f\x63\x61\x6c\x65\x73\x20\x3d\x20\x7b\x7d\x3b\x0a\x0a\x20\x20\x20\x20\x66\x6f\x72\x20\x28\x69\x6e\x64\x65\x78\x20\x69\x6e\x20\x61\x72\x67\x75\x6d\x65\x6e\x74\x73\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x66\x6f\x72\x20\x28\x70\x72\x6f\x70\x65\x72\x74\x79\x20\x69\x6e\x20\x61\x72\x67\x75\x6d\x65\x6e\x74\x73\x5b\x69\x6e\x64\x65\x78\x5d\x29\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x6c\x6f\x63\x61\x6c\x65\x73\x5b\x70\x72\x6f\x70\x65\x72\x74\x79\x5d\x20\x3d\x20\x61\x72\x67\x75\x6d\x65\x6e\x74\x73\x5b\x69\x6e\x64\x65\x78\x5d\x5b\x70\x72\x6f\x70\x65\x72\x74\x79\x5d\x3b\x0a\x20\x20\x20\x20\x7d\x0a\x20\x20\x20\x20\x69\x66\x20\x28\x20\x21\x20\x6c\x6f\x63\x61\x6c\x65\x73\x5b\x27\x65\x6e\x27\x5d\x29\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x6c\x6f\x63\x61\x6c\x65\x73\x5b\x27\x65\x6e\x27\x5d\x20\x3d\x20\x7b\x7d\x3b\x0a\x0a\x20\x20\x20\x20\x69\x66\x20\x28\x20\x21\x20\x6c\x6f\x63\x61\x6c\x65\x73\x5b\x6c\x61\x6e\x67\x49\x64\x5d\x20\x26\x26\x20\x21\x20\x6c\x6f\x63\x61\x6c\x65\x73\x5b\x6c\x61\x6e\x67\x75\x61\x67\x65\x5d\x29\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x6c\x61\x6e\x67\x75\x61\x67\x65\x20\x3d\x20\x27\x65\x6e\x27\x3b\x0a\x0a\x20\x20\x20\x20\x76\x61\x72\x20\x6c\x6f\x63\x61\x6c\x65\x20\x3d\x20\x28\x6c\x6f\x63\x61\x6c\x65\x73\x5b\x6c\x61\x6e\x67\x49\x64\x5d\x20\x3f\x20\x6c\x6f\x63\x61\x6c\x65\x73\x5b\x6c\x61\x6e\x67\x49\x64\x5d\x20\x3a\x20\x6c\x6f\x63\x61\x6c\x65\x73\x5b\x6c\x61\x6e\x67\x75\x61\x67\x65\x5d\x29\x3b\x0a\x0a\x20\x20\x20\x20\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x5f\x5f\x28\x74\x65\x78\x74\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x76\x61\x72\x20\x69\x6e\x64\x65\x78\x20\x3d\x20\x6c\x6f\x63\x61\x6c\x65\x5b\x74\x65\x78\x74\x5d\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x69\x66\x20\x28\x69\x6e\x64\x65\x78\x20\x3d\x3d\x3d\x20\x75\x6e\x64\x65\x66\x69\x6e\x65\x64\x29\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x20\x74\x65\x78\x74\x3b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x20\x69\x6e\x64\x65\x78\x3b\x0a\x20\x20\x20\x20\x7d\x3b\x0a\x0a\x20\x20\x20\x20\x66\x75\x6e\x63\x74\x69\x6f\x6e\x20\x73\x65\x74\x4c\x61\x6e\x67\x75\x61\x67\x65\x28\x6c\x61\x6e\x67\x75\x61\x67\x65\x29\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x6c\x6f\x63\x61\x6c\x65\x20\x3d\x20\x6c\x6f\x63\x61\x6c\x65\x73\x5b\x6c\x61\x6e\x67\x75\x61\x67\x65\x5d\x3b\x0a\x20\x20\x20\x20\x7d\x0a\x0a\x20\x20\x20\x20\x72\x65\x74\x75\x72\x6e\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x5f\x5f\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x5f\x5f\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x6c\x6f\x63\x61\x6c\x65\x73\x20\x20\x20\x20\x3a\x20\x6c\x6f\x63\x61\x6c\x65\x73\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x6c\x6f\x63\x61\x6c\x65\x20\x20\x20\x20\x20\x3a\x20\x6c\x6f\x63\x61\x6c\x65\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x73\x65\x74\x4c\x61\x6e\x67\x75\x61\x67\x65\x3a\x20\x73\x65\x74\x4c\x61\x6e\x67\x75\x61\x67\x65\x0a\x20\x20\x20\x20\x7d\x3b\x0a\x7d\x29\x3b\x0a")

func init() {
	filepath := "/locales/locale.js"
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
	_, err = f.Write(locale_js)
	if err != nil {
		log.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
