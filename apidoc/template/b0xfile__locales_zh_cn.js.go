
// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package template
import (
"log"
"os"
"path"
)
var zh_cn_js = []byte("\x64\x65\x66\x69\x6e\x65\x28\x7b\x0a\x20\x20\x20\x20\x27\x7a\x68\x5f\x63\x6e\x27\x3a\x20\x7b\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x41\x6c\x6c\x6f\x77\x65\x64\x20\x76\x61\x6c\x75\x65\x73\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xe5\x85\x81\xe8\xae\xb8\xe5\x80\xbc\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x43\x6f\x6d\x70\x61\x72\x65\x20\x61\x6c\x6c\x20\x77\x69\x74\x68\x20\x70\x72\x65\x64\x65\x63\x65\x73\x73\x6f\x72\x27\x3a\x20\x27\xe4\xb8\x8e\xe6\x89\x80\xe6\x9c\x89\xe8\xbe\x83\xe6\x97\xa9\xe7\x9a\x84\xe6\xaf\x94\xe8\xbe\x83\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x63\x6f\x6d\x70\x61\x72\x65\x20\x63\x68\x61\x6e\x67\x65\x73\x20\x74\x6f\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xe5\xb0\x86\xe5\xbd\x93\xe5\x89\x8d\xe7\x89\x88\xe6\x9c\xac\xe4\xb8\x8e\xe6\x8c\x87\xe5\xae\x9a\xe7\x89\x88\xe6\x9c\xac\xe6\xaf\x94\xe8\xbe\x83\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x63\x6f\x6d\x70\x61\x72\x65\x64\x20\x74\x6f\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xe7\x9b\xb8\xe6\xaf\x94\xe4\xba\x8e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x44\x65\x66\x61\x75\x6c\x74\x20\x76\x61\x6c\x75\x65\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xe9\xbb\x98\xe8\xae\xa4\xe5\x80\xbc\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x44\x65\x73\x63\x72\x69\x70\x74\x69\x6f\x6e\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xe6\x8f\x8f\xe8\xbf\xb0\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x46\x69\x65\x6c\x64\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xe5\xad\x97\xe6\xae\xb5\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x47\x65\x6e\x65\x72\x61\x6c\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xe6\xa6\x82\xe8\xa6\x81\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x47\x65\x6e\x65\x72\x61\x74\x65\x64\x20\x77\x69\x74\x68\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xe5\x9f\xba\xe4\xba\x8e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x4e\x61\x6d\x65\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xe5\x90\x8d\xe7\xa7\xb0\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x4e\x6f\x20\x72\x65\x73\x70\x6f\x6e\x73\x65\x20\x76\x61\x6c\x75\x65\x73\x2e\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xe6\x97\xa0\xe8\xbf\x94\xe5\x9b\x9e\xe5\x80\xbc\x2e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x6f\x70\x74\x69\x6f\x6e\x61\x6c\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xe5\x8f\xaf\xe9\x80\x89\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x50\x61\x72\x61\x6d\x65\x74\x65\x72\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xe5\x8f\x82\xe6\x95\xb0\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x50\x65\x72\x6d\x69\x73\x73\x69\x6f\x6e\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xe6\x9d\x83\xe9\x99\x90\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x52\x65\x73\x70\x6f\x6e\x73\x65\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xe8\xbf\x94\xe5\x9b\x9e\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x53\x65\x6e\x64\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xe5\x8f\x91\xe9\x80\x81\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x53\x65\x6e\x64\x20\x61\x20\x53\x61\x6d\x70\x6c\x65\x20\x52\x65\x71\x75\x65\x73\x74\x27\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xe5\x8f\x91\xe9\x80\x81\xe7\xa4\xba\xe4\xbe\x8b\xe8\xaf\xb7\xe6\xb1\x82\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x73\x68\x6f\x77\x20\x75\x70\x20\x74\x6f\x20\x76\x65\x72\x73\x69\x6f\x6e\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xe6\x98\xbe\xe7\xa4\xba\xe5\x88\xb0\xe6\x8c\x87\xe5\xae\x9a\xe7\x89\x88\xe6\x9c\xac\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x53\x69\x7a\x65\x20\x72\x61\x6e\x67\x65\x3a\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xe5\x8f\x96\xe5\x80\xbc\xe8\x8c\x83\xe5\x9b\xb4\x3a\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x54\x79\x70\x65\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xe7\xb1\xbb\xe5\x9e\x8b\x27\x2c\x0a\x20\x20\x20\x20\x20\x20\x20\x20\x27\x75\x72\x6c\x27\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x20\x3a\x20\x27\xe7\xbd\x91\xe5\x9d\x80\x27\x0a\x20\x20\x20\x20\x7d\x0a\x7d\x29\x3b\x0a")
func init() {
filepath := "/locales/zh_cn.js"
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
_, err = f.Write(zh_cn_js)
if err != nil {
	log.Fatal(err)
}
err = f.Close()
if err != nil {
	log.Fatal(err)
}
}