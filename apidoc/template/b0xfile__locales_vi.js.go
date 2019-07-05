
// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package template

import (
	"log"
	"os"
	"path"
)

var vi_js = []byte("\x64\x65\x66\x69\x6e\x65\x28\x7b\x0a\x20\x20\x20\x20\x76\x69\x3a\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x41\x6c\x6c\x6f\x77\x65\x64\x20\x76\x61\x6c\x75\x65\x73\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x47\x69\xc3\xa1\x20\x74\x72\xe1\xbb\x8b\x20\x63\x68\xe1\xba\xa5\x70\x20\x6e\x68\xe1\xba\xad\x6e\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x43\x6f\x6d\x70\x61\x72\x65\x20\x61\x6c\x6c\x20\x77\x69\x74\x68\x20\x70\x72\x65\x64\x65\x63\x65\x73\x73\x6f\x72\x27\x3a\x20\x27\x53\x6f\x20\x73\xc3\xa1\x6e\x68\x20\x76\xe1\xbb\x9b\x69\x20\x74\xe1\xba\xa5\x74\x20\x63\xe1\xba\xa3\x20\x70\x68\x69\xc3\xaa\x6e\x20\x62\xe1\xba\xa3\x6e\x20\x74\x72\xc6\xb0\xe1\xbb\x9b\x63\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x63\x6f\x6d\x70\x61\x72\x65\x20\x63\x68\x61\x6e\x67\x65\x73\x20\x74\x6f\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x73\x6f\x20\x73\xc3\xa1\x6e\x68\x20\x73\xe1\xbb\xb1\x20\x74\x68\x61\x79\x20\xc4\x91\xe1\xbb\x95\x69\x20\x76\xe1\xbb\x9b\x69\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x63\x6f\x6d\x70\x61\x72\x65\x64\x20\x74\x6f\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x73\x6f\x20\x73\xc3\xa1\x6e\x68\x20\x76\xe1\xbb\x9b\x69\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x44\x65\x66\x61\x75\x6c\x74\x20\x76\x61\x6c\x75\x65\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x47\x69\xc3\xa1\x20\x74\x72\xe1\xbb\x8b\x20\x6d\xe1\xba\xb7\x63\x20\xc4\x91\xe1\xbb\x8b\x6e\x68\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x44\x65\x73\x63\x72\x69\x70\x74\x69\x6f\x6e\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x43\x68\xc3\xba\x20\x74\x68\xc3\xad\x63\x68\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x46\x69\x65\x6c\x64\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x54\x72\xc6\xb0\xe1\xbb\x9d\x6e\x67\x20\x64\xe1\xbb\xaf\x20\x6c\x69\xe1\xbb\x87\x75\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x47\x65\x6e\x65\x72\x61\x6c\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x54\xe1\xbb\x95\x6e\x67\x20\x71\x75\x61\x6e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x47\x65\x6e\x65\x72\x61\x74\x65\x64\x20\x77\x69\x74\x68\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xc4\x90\xc6\xb0\xe1\xbb\xa3\x63\x20\x74\xe1\xba\xa1\x6f\x20\x62\xe1\xbb\x9f\x69\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x4e\x61\x6d\x65\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x54\xc3\xaa\x6e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x4e\x6f\x20\x72\x65\x73\x70\x6f\x6e\x73\x65\x20\x76\x61\x6c\x75\x65\x73\x2e\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x4b\x68\xc3\xb4\x6e\x67\x20\x63\xc3\xb3\x20\x6b\xe1\xba\xbf\x74\x20\x71\x75\xe1\xba\xa3\x20\x74\x72\xe1\xba\xa3\x20\x76\xe1\xbb\x81\x2e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x6f\x70\x74\x69\x6f\x6e\x61\x6c\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x54\xc3\xb9\x79\x20\x63\x68\xe1\xbb\x8d\x6e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x50\x61\x72\x61\x6d\x65\x74\x65\x72\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x54\x68\x61\x6d\x20\x73\xe1\xbb\x91\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x50\x65\x72\x6d\x69\x73\x73\x69\x6f\x6e\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x51\x75\x79\xe1\xbb\x81\x6e\x20\x68\xe1\xba\xa1\x6e\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x52\x65\x73\x70\x6f\x6e\x73\x65\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x4b\xe1\xba\xbf\x74\x20\x71\x75\xe1\xba\xa3\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x53\x65\x6e\x64\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x47\xe1\xbb\xad\x69\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x53\x65\x6e\x64\x20\x61\x20\x53\x61\x6d\x70\x6c\x65\x20\x52\x65\x71\x75\x65\x73\x74\x27\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x47\xe1\xbb\xad\x69\x20\x6d\xe1\xbb\x99\x74\x20\x79\xc3\xaa\x75\x20\x63\xe1\xba\xa7\x75\x20\x6d\xe1\xba\xab\x75\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x73\x68\x6f\x77\x20\x75\x70\x20\x74\x6f\x20\x76\x65\x72\x73\x69\x6f\x6e\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x68\x69\xe1\xbb\x83\x6e\x20\x74\x68\xe1\xbb\x8b\x20\x70\x68\x69\xc3\xaa\x6e\x20\x62\xe1\xba\xa3\x6e\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x53\x69\x7a\x65\x20\x72\x61\x6e\x67\x65\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x4b\xc3\xad\x63\x68\x20\x63\xe1\xbb\xa1\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x54\x79\x70\x65\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x4b\x69\xe1\xbb\x83\x75\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x75\x72\x6c\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\x6c\x69\xc3\xaa\x6e\x20\x6b\xe1\xba\xbf\x74\x27\x0a\x20\x20\x20\x20\x7d\x0a\x7d\x29\x3b\x0a")

func init() {
	filepath := "/locales/vi.js"
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
	_, err = f.Write(vi_js)
	if err != nil {
		log.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}