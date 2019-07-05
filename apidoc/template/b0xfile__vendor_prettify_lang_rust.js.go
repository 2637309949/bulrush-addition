
// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package template
import (
"log"
"os"
"path"
)
var lang_rust_js = []byte("\x2f\x2a\x0a\x0a\x20\x43\x6f\x70\x79\x72\x69\x67\x68\x74\x20\x28\x43\x29\x20\x32\x30\x31\x35\x20\x43\x68\x72\x69\x73\x20\x4d\x6f\x72\x67\x61\x6e\x0a\x0a\x20\x4c\x69\x63\x65\x6e\x73\x65\x64\x20\x75\x6e\x64\x65\x72\x20\x74\x68\x65\x20\x41\x70\x61\x63\x68\x65\x20\x4c\x69\x63\x65\x6e\x73\x65\x2c\x20\x56\x65\x72\x73\x69\x6f\x6e\x20\x32\x2e\x30\x20\x28\x74\x68\x65\x20\x22\x4c\x69\x63\x65\x6e\x73\x65\x22\x29\x3b\x0a\x20\x79\x6f\x75\x20\x6d\x61\x79\x20\x6e\x6f\x74\x20\x75\x73\x65\x20\x74\x68\x69\x73\x20\x66\x69\x6c\x65\x20\x65\x78\x63\x65\x70\x74\x20\x69\x6e\x20\x63\x6f\x6d\x70\x6c\x69\x61\x6e\x63\x65\x20\x77\x69\x74\x68\x20\x74\x68\x65\x20\x4c\x69\x63\x65\x6e\x73\x65\x2e\x0a\x20\x59\x6f\x75\x20\x6d\x61\x79\x20\x6f\x62\x74\x61\x69\x6e\x20\x61\x20\x63\x6f\x70\x79\x20\x6f\x66\x20\x74\x68\x65\x20\x4c\x69\x63\x65\x6e\x73\x65\x20\x61\x74\x0a\x0a\x20\x20\x20\x20\x68\x74\x74\x70\x3a\x2f\x2f\x77\x77\x77\x2e\x61\x70\x61\x63\x68\x65\x2e\x6f\x72\x67\x2f\x6c\x69\x63\x65\x6e\x73\x65\x73\x2f\x4c\x49\x43\x45\x4e\x53\x45\x2d\x32\x2e\x30\x0a\x0a\x20\x55\x6e\x6c\x65\x73\x73\x20\x72\x65\x71\x75\x69\x72\x65\x64\x20\x62\x79\x20\x61\x70\x70\x6c\x69\x63\x61\x62\x6c\x65\x20\x6c\x61\x77\x20\x6f\x72\x20\x61\x67\x72\x65\x65\x64\x20\x74\x6f\x20\x69\x6e\x20\x77\x72\x69\x74\x69\x6e\x67\x2c\x20\x73\x6f\x66\x74\x77\x61\x72\x65\x0a\x20\x64\x69\x73\x74\x72\x69\x62\x75\x74\x65\x64\x20\x75\x6e\x64\x65\x72\x20\x74\x68\x65\x20\x4c\x69\x63\x65\x6e\x73\x65\x20\x69\x73\x20\x64\x69\x73\x74\x72\x69\x62\x75\x74\x65\x64\x20\x6f\x6e\x20\x61\x6e\x20\x22\x41\x53\x20\x49\x53\x22\x20\x42\x41\x53\x49\x53\x2c\x0a\x20\x57\x49\x54\x48\x4f\x55\x54\x20\x57\x41\x52\x52\x41\x4e\x54\x49\x45\x53\x20\x4f\x52\x20\x43\x4f\x4e\x44\x49\x54\x49\x4f\x4e\x53\x20\x4f\x46\x20\x41\x4e\x59\x20\x4b\x49\x4e\x44\x2c\x20\x65\x69\x74\x68\x65\x72\x20\x65\x78\x70\x72\x65\x73\x73\x20\x6f\x72\x20\x69\x6d\x70\x6c\x69\x65\x64\x2e\x0a\x20\x53\x65\x65\x20\x74\x68\x65\x20\x4c\x69\x63\x65\x6e\x73\x65\x20\x66\x6f\x72\x20\x74\x68\x65\x20\x73\x70\x65\x63\x69\x66\x69\x63\x20\x6c\x61\x6e\x67\x75\x61\x67\x65\x20\x67\x6f\x76\x65\x72\x6e\x69\x6e\x67\x20\x70\x65\x72\x6d\x69\x73\x73\x69\x6f\x6e\x73\x20\x61\x6e\x64\x0a\x20\x6c\x69\x6d\x69\x74\x61\x74\x69\x6f\x6e\x73\x20\x75\x6e\x64\x65\x72\x20\x74\x68\x65\x20\x4c\x69\x63\x65\x6e\x73\x65\x2e\x0a\x2a\x2f\x0a\x50\x52\x2e\x72\x65\x67\x69\x73\x74\x65\x72\x4c\x61\x6e\x67\x48\x61\x6e\x64\x6c\x65\x72\x28\x50\x52\x2e\x63\x72\x65\x61\x74\x65\x53\x69\x6d\x70\x6c\x65\x4c\x65\x78\x65\x72\x28\x5b\x5d\x2c\x5b\x5b\x22\x70\x6c\x6e\x22\x2c\x2f\x5e\x5b\x5c\x74\x5c\x6e\x5c\x72\x20\x5c\x78\x41\x30\x5d\x2b\x2f\x5d\x2c\x5b\x22\x63\x6f\x6d\x22\x2c\x2f\x5e\x5c\x2f\x5c\x2f\x2e\x2a\x2f\x5d\x2c\x5b\x22\x63\x6f\x6d\x22\x2c\x2f\x5e\x5c\x2f\x5c\x2a\x5b\x5c\x73\x5c\x53\x5d\x2a\x3f\x28\x3f\x3a\x5c\x2a\x5c\x2f\x7c\x24\x29\x2f\x5d\x2c\x5b\x22\x73\x74\x72\x22\x2c\x2f\x5e\x62\x22\x28\x3f\x3a\x5b\x5e\x5c\x5c\x5d\x7c\x5c\x5c\x28\x3f\x3a\x2e\x7c\x78\x5b\x5c\x64\x61\x2d\x66\x41\x2d\x46\x5d\x7b\x32\x7d\x29\x29\x2a\x3f\x22\x2f\x5d\x2c\x5b\x22\x73\x74\x72\x22\x2c\x2f\x5e\x22\x28\x3f\x3a\x5b\x5e\x5c\x5c\x5d\x7c\x5c\x5c\x28\x3f\x3a\x2e\x7c\x78\x5b\x5c\x64\x61\x2d\x66\x41\x2d\x46\x5d\x7b\x32\x7d\x7c\x75\x5c\x7b\x5c\x5b\x5c\x64\x61\x2d\x66\x41\x2d\x46\x5d\x7b\x31\x2c\x36\x7d\x5c\x7d\x29\x29\x2a\x3f\x22\x2f\x5d\x2c\x5b\x22\x73\x74\x72\x22\x2c\x2f\x5e\x62\x3f\x72\x28\x23\x2a\x29\x5c\x22\x5b\x5c\x73\x5c\x53\x5d\x2a\x3f\x5c\x22\x5c\x31\x2f\x5d\x2c\x5b\x22\x73\x74\x72\x22\x2c\x2f\x5e\x62\x27\x28\x5b\x5e\x5c\x5c\x5d\x7c\x5c\x5c\x28\x2e\x7c\x78\x5b\x5c\x64\x61\x2d\x66\x41\x2d\x46\x5d\x7b\x32\x7d\x29\x29\x27\x2f\x5d\x2c\x5b\x22\x73\x74\x72\x22\x2c\x2f\x5e\x27\x28\x5b\x5e\x5c\x5c\x5d\x7c\x5c\x5c\x28\x2e\x7c\x78\x5b\x5c\x64\x61\x2d\x66\x41\x2d\x46\x5d\x7b\x32\x7d\x7c\x75\x5c\x7b\x5b\x5c\x64\x61\x2d\x66\x41\x2d\x46\x5d\x7b\x31\x2c\x36\x7d\x5c\x7d\x29\x29\x27\x2f\x5d\x2c\x5b\x22\x74\x61\x67\x22\x2c\x2f\x5e\x27\x5c\x77\x2b\x3f\x5c\x62\x2f\x5d\x2c\x5b\x22\x6b\x77\x64\x22\x2c\x2f\x5e\x28\x3f\x3a\x6d\x61\x74\x63\x68\x7c\x69\x66\x7c\x65\x6c\x73\x65\x7c\x61\x73\x7c\x62\x72\x65\x61\x6b\x7c\x62\x6f\x78\x7c\x63\x6f\x6e\x74\x69\x6e\x75\x65\x7c\x65\x78\x74\x65\x72\x6e\x7c\x66\x6e\x7c\x66\x6f\x72\x7c\x69\x6e\x7c\x69\x66\x7c\x69\x6d\x70\x6c\x7c\x6c\x65\x74\x7c\x6c\x6f\x6f\x70\x7c\x70\x75\x62\x7c\x72\x65\x74\x75\x72\x6e\x7c\x73\x75\x70\x65\x72\x7c\x75\x6e\x73\x61\x66\x65\x7c\x77\x68\x65\x72\x65\x7c\x77\x68\x69\x6c\x65\x7c\x75\x73\x65\x7c\x6d\x6f\x64\x7c\x74\x72\x61\x69\x74\x7c\x73\x74\x72\x75\x63\x74\x7c\x65\x6e\x75\x6d\x7c\x74\x79\x70\x65\x7c\x6d\x6f\x76\x65\x7c\x6d\x75\x74\x7c\x72\x65\x66\x7c\x73\x74\x61\x74\x69\x63\x7c\x63\x6f\x6e\x73\x74\x7c\x63\x72\x61\x74\x65\x29\x5c\x62\x2f\x5d\x2c\x0a\x5b\x22\x6b\x77\x64\x22\x2c\x2f\x5e\x28\x3f\x3a\x61\x6c\x69\x67\x6e\x6f\x66\x7c\x62\x65\x63\x6f\x6d\x65\x7c\x64\x6f\x7c\x6f\x66\x66\x73\x65\x74\x6f\x66\x7c\x70\x72\x69\x76\x7c\x70\x75\x72\x65\x7c\x73\x69\x7a\x65\x6f\x66\x7c\x74\x79\x70\x65\x6f\x66\x7c\x75\x6e\x73\x69\x7a\x65\x64\x7c\x79\x69\x65\x6c\x64\x7c\x61\x62\x73\x74\x72\x61\x63\x74\x7c\x76\x69\x72\x74\x75\x61\x6c\x7c\x66\x69\x6e\x61\x6c\x7c\x6f\x76\x65\x72\x72\x69\x64\x65\x7c\x6d\x61\x63\x72\x6f\x29\x5c\x62\x2f\x5d\x2c\x5b\x22\x74\x79\x70\x22\x2c\x2f\x5e\x28\x3f\x3a\x5b\x69\x75\x5d\x28\x38\x7c\x31\x36\x7c\x33\x32\x7c\x36\x34\x7c\x73\x69\x7a\x65\x29\x7c\x63\x68\x61\x72\x7c\x62\x6f\x6f\x6c\x7c\x66\x33\x32\x7c\x66\x36\x34\x7c\x73\x74\x72\x7c\x53\x65\x6c\x66\x29\x5c\x62\x2f\x5d\x2c\x5b\x22\x74\x79\x70\x22\x2c\x2f\x5e\x28\x3f\x3a\x43\x6f\x70\x79\x7c\x53\x65\x6e\x64\x7c\x53\x69\x7a\x65\x64\x7c\x53\x79\x6e\x63\x7c\x44\x72\x6f\x70\x7c\x46\x6e\x7c\x46\x6e\x4d\x75\x74\x7c\x46\x6e\x4f\x6e\x63\x65\x7c\x42\x6f\x78\x7c\x54\x6f\x4f\x77\x6e\x65\x64\x7c\x43\x6c\x6f\x6e\x65\x7c\x50\x61\x72\x74\x69\x61\x6c\x45\x71\x7c\x50\x61\x72\x74\x69\x61\x6c\x4f\x72\x64\x7c\x45\x71\x7c\x4f\x72\x64\x7c\x41\x73\x52\x65\x66\x7c\x41\x73\x4d\x75\x74\x7c\x49\x6e\x74\x6f\x7c\x46\x72\x6f\x6d\x7c\x44\x65\x66\x61\x75\x6c\x74\x7c\x49\x74\x65\x72\x61\x74\x6f\x72\x7c\x45\x78\x74\x65\x6e\x64\x7c\x49\x6e\x74\x6f\x49\x74\x65\x72\x61\x74\x6f\x72\x7c\x44\x6f\x75\x62\x6c\x65\x45\x6e\x64\x65\x64\x49\x74\x65\x72\x61\x74\x6f\x72\x7c\x45\x78\x61\x63\x74\x53\x69\x7a\x65\x49\x74\x65\x72\x61\x74\x6f\x72\x7c\x4f\x70\x74\x69\x6f\x6e\x7c\x53\x6f\x6d\x65\x7c\x4e\x6f\x6e\x65\x7c\x52\x65\x73\x75\x6c\x74\x7c\x4f\x6b\x7c\x45\x72\x72\x7c\x53\x6c\x69\x63\x65\x43\x6f\x6e\x63\x61\x74\x45\x78\x74\x7c\x53\x74\x72\x69\x6e\x67\x7c\x54\x6f\x53\x74\x72\x69\x6e\x67\x7c\x56\x65\x63\x29\x5c\x62\x2f\x5d\x2c\x5b\x22\x6c\x69\x74\x22\x2c\x2f\x5e\x28\x73\x65\x6c\x66\x7c\x74\x72\x75\x65\x7c\x66\x61\x6c\x73\x65\x7c\x6e\x75\x6c\x6c\x29\x5c\x62\x2f\x5d\x2c\x0a\x5b\x22\x6c\x69\x74\x22\x2c\x2f\x5e\x5c\x64\x5b\x30\x2d\x39\x5f\x5d\x2a\x28\x3f\x3a\x5b\x69\x75\x5d\x28\x3f\x3a\x73\x69\x7a\x65\x7c\x38\x7c\x31\x36\x7c\x33\x32\x7c\x36\x34\x29\x29\x3f\x2f\x5d\x2c\x5b\x22\x6c\x69\x74\x22\x2c\x2f\x5e\x30\x78\x5b\x61\x2d\x66\x41\x2d\x46\x30\x2d\x39\x5f\x5d\x2b\x28\x3f\x3a\x5b\x69\x75\x5d\x28\x3f\x3a\x73\x69\x7a\x65\x7c\x38\x7c\x31\x36\x7c\x33\x32\x7c\x36\x34\x29\x29\x3f\x2f\x5d\x2c\x5b\x22\x6c\x69\x74\x22\x2c\x2f\x5e\x30\x6f\x5b\x30\x2d\x37\x5f\x5d\x2b\x28\x3f\x3a\x5b\x69\x75\x5d\x28\x3f\x3a\x73\x69\x7a\x65\x7c\x38\x7c\x31\x36\x7c\x33\x32\x7c\x36\x34\x29\x29\x3f\x2f\x5d\x2c\x5b\x22\x6c\x69\x74\x22\x2c\x2f\x5e\x30\x62\x5b\x30\x31\x5f\x5d\x2b\x28\x3f\x3a\x5b\x69\x75\x5d\x28\x3f\x3a\x73\x69\x7a\x65\x7c\x38\x7c\x31\x36\x7c\x33\x32\x7c\x36\x34\x29\x29\x3f\x2f\x5d\x2c\x5b\x22\x6c\x69\x74\x22\x2c\x2f\x5e\x5c\x64\x5b\x30\x2d\x39\x5f\x5d\x2a\x5c\x2e\x28\x3f\x21\x5b\x5e\x5c\x73\x5c\x64\x2e\x5d\x29\x2f\x5d\x2c\x5b\x22\x6c\x69\x74\x22\x2c\x2f\x5e\x5c\x64\x5b\x30\x2d\x39\x5f\x5d\x2a\x28\x3f\x3a\x5c\x2e\x5c\x64\x5b\x30\x2d\x39\x5f\x5d\x2a\x29\x28\x3f\x3a\x5b\x65\x45\x5d\x5b\x2b\x2d\x5d\x3f\x5b\x30\x2d\x39\x5f\x5d\x2b\x29\x3f\x28\x3f\x3a\x66\x33\x32\x7c\x66\x36\x34\x29\x3f\x2f\x5d\x2c\x5b\x22\x6c\x69\x74\x22\x2c\x2f\x5e\x5c\x64\x5b\x30\x2d\x39\x5f\x5d\x2a\x28\x3f\x3a\x5c\x2e\x5c\x64\x5b\x30\x2d\x39\x5f\x5d\x2a\x29\x3f\x28\x3f\x3a\x5b\x65\x45\x5d\x5b\x2b\x2d\x5d\x3f\x5b\x30\x2d\x39\x5f\x5d\x2b\x29\x28\x3f\x3a\x66\x33\x32\x7c\x66\x36\x34\x29\x3f\x2f\x5d\x2c\x5b\x22\x6c\x69\x74\x22\x2c\x2f\x5e\x5c\x64\x5b\x30\x2d\x39\x5f\x5d\x2a\x28\x3f\x3a\x5c\x2e\x5c\x64\x5b\x30\x2d\x39\x5f\x5d\x2a\x29\x3f\x28\x3f\x3a\x5b\x65\x45\x5d\x5b\x2b\x2d\x5d\x3f\x5b\x30\x2d\x39\x5f\x5d\x2b\x29\x3f\x28\x3f\x3a\x66\x33\x32\x7c\x66\x36\x34\x29\x2f\x5d\x2c\x0a\x5b\x22\x61\x74\x6e\x22\x2c\x2f\x5e\x5b\x61\x2d\x7a\x5f\x5d\x5c\x77\x2a\x21\x2f\x69\x5d\x2c\x5b\x22\x70\x6c\x6e\x22\x2c\x2f\x5e\x5b\x61\x2d\x7a\x5f\x5d\x5c\x77\x2a\x2f\x69\x5d\x2c\x5b\x22\x61\x74\x76\x22\x2c\x2f\x5e\x23\x21\x3f\x5c\x5b\x5b\x5c\x73\x5c\x53\x5d\x2a\x3f\x5c\x5d\x2f\x5d\x2c\x5b\x22\x70\x75\x6e\x22\x2c\x2f\x5e\x5b\x2b\x5c\x2d\x2f\x2a\x3d\x5e\x26\x7c\x21\x3c\x3e\x25\x5b\x5c\x5d\x28\x29\x7b\x7d\x3f\x3a\x2e\x2c\x3b\x5d\x2f\x5d\x2c\x5b\x22\x70\x6c\x6e\x22\x2c\x2f\x2e\x2f\x5d\x5d\x29\x2c\x5b\x22\x72\x75\x73\x74\x22\x5d\x29\x3b\x0a")
func init() {
filepath := "/vendor/prettify/lang-rust.js"
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
_, err = f.Write(lang_rust_js)
if err != nil {
	log.Fatal(err)
}
err = f.Close()
if err != nil {
	log.Fatal(err)
}
}