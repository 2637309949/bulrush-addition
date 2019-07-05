// Copyright (c) 2018-2020 Double All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package template

import (
	"log"
	"os"
	"path"
)

var lang_lasso_js = []byte("\x2f\x2a\x0a\x0a\x20\x43\x6f\x70\x79\x72\x69\x67\x68\x74\x20\x28\x43\x29\x20\x32\x30\x31\x33\x20\x45\x72\x69\x63\x20\x4b\x6e\x69\x62\x62\x65\x0a\x0a\x20\x4c\x69\x63\x65\x6e\x73\x65\x64\x20\x75\x6e\x64\x65\x72\x20\x74\x68\x65\x20\x41\x70\x61\x63\x68\x65\x20\x4c\x69\x63\x65\x6e\x73\x65\x2c\x20\x56\x65\x72\x73\x69\x6f\x6e\x20\x32\x2e\x30\x20\x28\x74\x68\x65\x20\x22\x4c\x69\x63\x65\x6e\x73\x65\x22\x29\x3b\x0a\x20\x79\x6f\x75\x20\x6d\x61\x79\x20\x6e\x6f\x74\x20\x75\x73\x65\x20\x74\x68\x69\x73\x20\x66\x69\x6c\x65\x20\x65\x78\x63\x65\x70\x74\x20\x69\x6e\x20\x63\x6f\x6d\x70\x6c\x69\x61\x6e\x63\x65\x20\x77\x69\x74\x68\x20\x74\x68\x65\x20\x4c\x69\x63\x65\x6e\x73\x65\x2e\x0a\x20\x59\x6f\x75\x20\x6d\x61\x79\x20\x6f\x62\x74\x61\x69\x6e\x20\x61\x20\x63\x6f\x70\x79\x20\x6f\x66\x20\x74\x68\x65\x20\x4c\x69\x63\x65\x6e\x73\x65\x20\x61\x74\x0a\x0a\x20\x20\x20\x20\x20\x20\x68\x74\x74\x70\x3a\x2f\x2f\x77\x77\x77\x2e\x61\x70\x61\x63\x68\x65\x2e\x6f\x72\x67\x2f\x6c\x69\x63\x65\x6e\x73\x65\x73\x2f\x4c\x49\x43\x45\x4e\x53\x45\x2d\x32\x2e\x30\x0a\x0a\x20\x55\x6e\x6c\x65\x73\x73\x20\x72\x65\x71\x75\x69\x72\x65\x64\x20\x62\x79\x20\x61\x70\x70\x6c\x69\x63\x61\x62\x6c\x65\x20\x6c\x61\x77\x20\x6f\x72\x20\x61\x67\x72\x65\x65\x64\x20\x74\x6f\x20\x69\x6e\x20\x77\x72\x69\x74\x69\x6e\x67\x2c\x20\x73\x6f\x66\x74\x77\x61\x72\x65\x0a\x20\x64\x69\x73\x74\x72\x69\x62\x75\x74\x65\x64\x20\x75\x6e\x64\x65\x72\x20\x74\x68\x65\x20\x4c\x69\x63\x65\x6e\x73\x65\x20\x69\x73\x20\x64\x69\x73\x74\x72\x69\x62\x75\x74\x65\x64\x20\x6f\x6e\x20\x61\x6e\x20\x22\x41\x53\x20\x49\x53\x22\x20\x42\x41\x53\x49\x53\x2c\x0a\x20\x57\x49\x54\x48\x4f\x55\x54\x20\x57\x41\x52\x52\x41\x4e\x54\x49\x45\x53\x20\x4f\x52\x20\x43\x4f\x4e\x44\x49\x54\x49\x4f\x4e\x53\x20\x4f\x46\x20\x41\x4e\x59\x20\x4b\x49\x4e\x44\x2c\x20\x65\x69\x74\x68\x65\x72\x20\x65\x78\x70\x72\x65\x73\x73\x20\x6f\x72\x20\x69\x6d\x70\x6c\x69\x65\x64\x2e\x0a\x20\x53\x65\x65\x20\x74\x68\x65\x20\x4c\x69\x63\x65\x6e\x73\x65\x20\x66\x6f\x72\x20\x74\x68\x65\x20\x73\x70\x65\x63\x69\x66\x69\x63\x20\x6c\x61\x6e\x67\x75\x61\x67\x65\x20\x67\x6f\x76\x65\x72\x6e\x69\x6e\x67\x20\x70\x65\x72\x6d\x69\x73\x73\x69\x6f\x6e\x73\x20\x61\x6e\x64\x0a\x20\x6c\x69\x6d\x69\x74\x61\x74\x69\x6f\x6e\x73\x20\x75\x6e\x64\x65\x72\x20\x74\x68\x65\x20\x4c\x69\x63\x65\x6e\x73\x65\x2e\x0a\x2a\x2f\x0a\x50\x52\x2e\x72\x65\x67\x69\x73\x74\x65\x72\x4c\x61\x6e\x67\x48\x61\x6e\x64\x6c\x65\x72\x28\x50\x52\x2e\x63\x72\x65\x61\x74\x65\x53\x69\x6d\x70\x6c\x65\x4c\x65\x78\x65\x72\x28\x5b\x5b\x22\x70\x6c\x6e\x22\x2c\x2f\x5e\x5b\x5c\x74\x5c\x6e\x5c\x72\x20\x5c\x78\x41\x30\x5d\x2b\x2f\x2c\x6e\x75\x6c\x6c\x2c\x22\x5c\x74\x5c\x6e\x5c\x72\x20\x5c\x75\x30\x30\x61\x30\x22\x5d\x2c\x5b\x22\x73\x74\x72\x22\x2c\x2f\x5e\x5c\x27\x28\x3f\x3a\x5b\x5e\x5c\x27\x5c\x5c\x5d\x7c\x5c\x5c\x5b\x5c\x73\x5c\x53\x5d\x29\x2a\x28\x3f\x3a\x5c\x27\x7c\x24\x29\x2f\x2c\x6e\x75\x6c\x6c\x2c\x22\x27\x22\x5d\x2c\x5b\x22\x73\x74\x72\x22\x2c\x2f\x5e\x5c\x22\x28\x3f\x3a\x5b\x5e\x5c\x22\x5c\x5c\x5d\x7c\x5c\x5c\x5b\x5c\x73\x5c\x53\x5d\x29\x2a\x28\x3f\x3a\x5c\x22\x7c\x24\x29\x2f\x2c\x6e\x75\x6c\x6c\x2c\x27\x22\x27\x5d\x2c\x5b\x22\x73\x74\x72\x22\x2c\x2f\x5e\x5c\x60\x5b\x5e\x5c\x60\x5d\x2a\x28\x3f\x3a\x5c\x60\x7c\x24\x29\x2f\x2c\x6e\x75\x6c\x6c\x2c\x22\x60\x22\x5d\x2c\x5b\x22\x6c\x69\x74\x22\x2c\x2f\x5e\x30\x78\x5b\x5c\x64\x61\x2d\x66\x5d\x2b\x7c\x5c\x64\x2b\x2f\x69\x2c\x6e\x75\x6c\x6c\x2c\x22\x30\x31\x32\x33\x34\x35\x36\x37\x38\x39\x22\x5d\x2c\x5b\x22\x61\x74\x6e\x22\x2c\x2f\x5e\x23\x5c\x64\x2b\x7c\x5b\x23\x24\x5d\x5b\x61\x2d\x7a\x5f\x5d\x5b\x5c\x77\x2e\x5d\x2a\x7c\x23\x21\x5b\x20\x5c\x53\x5d\x2b\x6c\x61\x73\x73\x6f\x39\x5c\x62\x2f\x69\x2c\x6e\x75\x6c\x6c\x2c\x22\x23\x24\x22\x5d\x5d\x2c\x5b\x5b\x22\x74\x61\x67\x22\x2c\x2f\x5e\x5b\x5b\x5c\x5d\x5d\x7c\x3c\x5c\x3f\x28\x3f\x3a\x6c\x61\x73\x73\x6f\x28\x3f\x3a\x73\x63\x72\x69\x70\x74\x29\x3f\x7c\x3d\x29\x7c\x5c\x3f\x3e\x7c\x6e\x6f\x70\x72\x6f\x63\x65\x73\x73\x5c\x62\x7c\x6e\x6f\x5f\x73\x71\x75\x61\x72\x65\x5f\x62\x72\x61\x63\x6b\x65\x74\x73\x5c\x62\x2f\x69\x5d\x2c\x5b\x22\x63\x6f\x6d\x22\x2c\x2f\x5e\x5c\x2f\x5c\x2f\x5b\x5e\x5c\x72\x5c\x6e\x5d\x2a\x7c\x5c\x2f\x5c\x2a\x5b\x5c\x73\x5c\x53\x5d\x2a\x3f\x5c\x2a\x5c\x2f\x2f\x5d\x2c\x0a\x5b\x22\x61\x74\x6e\x22\x2c\x2f\x5e\x2d\x28\x3f\x21\x69\x6e\x66\x69\x6e\x69\x74\x79\x29\x5b\x61\x2d\x7a\x5f\x5d\x5b\x5c\x77\x2e\x5d\x2a\x7c\x5c\x2e\x5c\x73\x2a\x27\x5b\x61\x2d\x7a\x5f\x5d\x5b\x5c\x77\x2e\x5d\x2a\x27\x2f\x69\x5d\x2c\x5b\x22\x6c\x69\x74\x22\x2c\x2f\x5e\x5c\x64\x2a\x5c\x2e\x5c\x64\x2b\x28\x3f\x3a\x65\x5b\x2d\x2b\x5d\x3f\x5c\x64\x2b\x29\x3f\x7c\x69\x6e\x66\x69\x6e\x69\x74\x79\x5c\x62\x7c\x4e\x61\x4e\x5c\x62\x2f\x69\x5d\x2c\x5b\x22\x61\x74\x76\x22\x2c\x2f\x5e\x3a\x3a\x5c\x73\x2a\x5b\x61\x2d\x7a\x5f\x5d\x5b\x5c\x77\x2e\x5d\x2a\x2f\x69\x5d\x2c\x5b\x22\x6c\x69\x74\x22\x2c\x2f\x5e\x28\x3f\x3a\x74\x72\x75\x65\x7c\x66\x61\x6c\x73\x65\x7c\x6e\x6f\x6e\x65\x7c\x6d\x69\x6e\x69\x6d\x61\x6c\x7c\x66\x75\x6c\x6c\x7c\x61\x6c\x6c\x7c\x76\x6f\x69\x64\x7c\x61\x6e\x64\x7c\x6f\x72\x7c\x6e\x6f\x74\x7c\x62\x77\x7c\x6e\x62\x77\x7c\x65\x77\x7c\x6e\x65\x77\x7c\x63\x6e\x7c\x6e\x63\x6e\x7c\x6c\x74\x7c\x6c\x74\x65\x7c\x67\x74\x7c\x67\x74\x65\x7c\x65\x71\x7c\x6e\x65\x71\x7c\x72\x78\x7c\x6e\x72\x78\x7c\x66\x74\x29\x5c\x62\x2f\x69\x5d\x2c\x5b\x22\x6b\x77\x64\x22\x2c\x2f\x5e\x28\x3f\x3a\x65\x72\x72\x6f\x72\x5f\x63\x6f\x64\x65\x7c\x65\x72\x72\x6f\x72\x5f\x6d\x73\x67\x7c\x65\x72\x72\x6f\x72\x5f\x70\x6f\x70\x7c\x65\x72\x72\x6f\x72\x5f\x70\x75\x73\x68\x7c\x65\x72\x72\x6f\x72\x5f\x72\x65\x73\x65\x74\x7c\x63\x61\x63\x68\x65\x7c\x64\x61\x74\x61\x62\x61\x73\x65\x5f\x6e\x61\x6d\x65\x73\x7c\x64\x61\x74\x61\x62\x61\x73\x65\x5f\x73\x63\x68\x65\x6d\x61\x6e\x61\x6d\x65\x73\x7c\x64\x61\x74\x61\x62\x61\x73\x65\x5f\x74\x61\x62\x6c\x65\x6e\x61\x6d\x65\x73\x7c\x64\x65\x66\x69\x6e\x65\x5f\x74\x61\x67\x7c\x64\x65\x66\x69\x6e\x65\x5f\x74\x79\x70\x65\x7c\x65\x6d\x61\x69\x6c\x5f\x62\x61\x74\x63\x68\x7c\x65\x6e\x63\x6f\x64\x65\x5f\x73\x65\x74\x7c\x68\x74\x6d\x6c\x5f\x63\x6f\x6d\x6d\x65\x6e\x74\x7c\x68\x61\x6e\x64\x6c\x65\x7c\x68\x61\x6e\x64\x6c\x65\x5f\x65\x72\x72\x6f\x72\x7c\x68\x65\x61\x64\x65\x72\x7c\x69\x66\x7c\x69\x6e\x6c\x69\x6e\x65\x7c\x69\x74\x65\x72\x61\x74\x65\x7c\x6c\x6a\x61\x78\x5f\x74\x61\x72\x67\x65\x74\x7c\x6c\x69\x6e\x6b\x7c\x6c\x69\x6e\x6b\x5f\x63\x75\x72\x72\x65\x6e\x74\x61\x63\x74\x69\x6f\x6e\x7c\x6c\x69\x6e\x6b\x5f\x63\x75\x72\x72\x65\x6e\x74\x67\x72\x6f\x75\x70\x7c\x6c\x69\x6e\x6b\x5f\x63\x75\x72\x72\x65\x6e\x74\x72\x65\x63\x6f\x72\x64\x7c\x6c\x69\x6e\x6b\x5f\x64\x65\x74\x61\x69\x6c\x7c\x6c\x69\x6e\x6b\x5f\x66\x69\x72\x73\x74\x67\x72\x6f\x75\x70\x7c\x6c\x69\x6e\x6b\x5f\x66\x69\x72\x73\x74\x72\x65\x63\x6f\x72\x64\x7c\x6c\x69\x6e\x6b\x5f\x6c\x61\x73\x74\x67\x72\x6f\x75\x70\x7c\x6c\x69\x6e\x6b\x5f\x6c\x61\x73\x74\x72\x65\x63\x6f\x72\x64\x7c\x6c\x69\x6e\x6b\x5f\x6e\x65\x78\x74\x67\x72\x6f\x75\x70\x7c\x6c\x69\x6e\x6b\x5f\x6e\x65\x78\x74\x72\x65\x63\x6f\x72\x64\x7c\x6c\x69\x6e\x6b\x5f\x70\x72\x65\x76\x67\x72\x6f\x75\x70\x7c\x6c\x69\x6e\x6b\x5f\x70\x72\x65\x76\x72\x65\x63\x6f\x72\x64\x7c\x6c\x6f\x67\x7c\x6c\x6f\x6f\x70\x7c\x6e\x61\x6d\x65\x73\x70\x61\x63\x65\x5f\x75\x73\x69\x6e\x67\x7c\x6f\x75\x74\x70\x75\x74\x5f\x6e\x6f\x6e\x65\x7c\x70\x6f\x72\x74\x61\x6c\x7c\x70\x72\x69\x76\x61\x74\x65\x7c\x70\x72\x6f\x74\x65\x63\x74\x7c\x72\x65\x63\x6f\x72\x64\x73\x7c\x72\x65\x66\x65\x72\x65\x72\x7c\x72\x65\x66\x65\x72\x72\x65\x72\x7c\x72\x65\x70\x65\x61\x74\x69\x6e\x67\x7c\x72\x65\x73\x75\x6c\x74\x73\x65\x74\x7c\x72\x6f\x77\x73\x7c\x73\x65\x61\x72\x63\x68\x5f\x61\x72\x67\x73\x7c\x73\x65\x61\x72\x63\x68\x5f\x61\x72\x67\x75\x6d\x65\x6e\x74\x73\x7c\x73\x65\x6c\x65\x63\x74\x7c\x73\x6f\x72\x74\x5f\x61\x72\x67\x73\x7c\x73\x6f\x72\x74\x5f\x61\x72\x67\x75\x6d\x65\x6e\x74\x73\x7c\x74\x68\x72\x65\x61\x64\x5f\x61\x74\x6f\x6d\x69\x63\x7c\x76\x61\x6c\x75\x65\x5f\x6c\x69\x73\x74\x7c\x77\x68\x69\x6c\x65\x7c\x61\x62\x6f\x72\x74\x7c\x63\x61\x73\x65\x7c\x65\x6c\x73\x65\x7c\x69\x66\x5f\x65\x6d\x70\x74\x79\x7c\x69\x66\x5f\x66\x61\x6c\x73\x65\x7c\x69\x66\x5f\x6e\x75\x6c\x6c\x7c\x69\x66\x5f\x74\x72\x75\x65\x7c\x6c\x6f\x6f\x70\x5f\x61\x62\x6f\x72\x74\x7c\x6c\x6f\x6f\x70\x5f\x63\x6f\x6e\x74\x69\x6e\x75\x65\x7c\x6c\x6f\x6f\x70\x5f\x63\x6f\x75\x6e\x74\x7c\x70\x61\x72\x61\x6d\x73\x7c\x70\x61\x72\x61\x6d\x73\x5f\x75\x70\x7c\x72\x65\x74\x75\x72\x6e\x7c\x72\x65\x74\x75\x72\x6e\x5f\x76\x61\x6c\x75\x65\x7c\x72\x75\x6e\x5f\x63\x68\x69\x6c\x64\x72\x65\x6e\x7c\x73\x6f\x61\x70\x5f\x64\x65\x66\x69\x6e\x65\x74\x61\x67\x7c\x73\x6f\x61\x70\x5f\x6c\x61\x73\x74\x72\x65\x71\x75\x65\x73\x74\x7c\x73\x6f\x61\x70\x5f\x6c\x61\x73\x74\x72\x65\x73\x70\x6f\x6e\x73\x65\x7c\x74\x61\x67\x5f\x6e\x61\x6d\x65\x7c\x61\x73\x63\x65\x6e\x64\x69\x6e\x67\x7c\x61\x76\x65\x72\x61\x67\x65\x7c\x62\x79\x7c\x64\x65\x66\x69\x6e\x65\x7c\x64\x65\x73\x63\x65\x6e\x64\x69\x6e\x67\x7c\x64\x6f\x7c\x65\x71\x75\x61\x6c\x73\x7c\x66\x72\x6f\x7a\x65\x6e\x7c\x67\x72\x6f\x75\x70\x7c\x68\x61\x6e\x64\x6c\x65\x5f\x66\x61\x69\x6c\x75\x72\x65\x7c\x69\x6d\x70\x6f\x72\x74\x7c\x69\x6e\x7c\x69\x6e\x74\x6f\x7c\x6a\x6f\x69\x6e\x7c\x6c\x65\x74\x7c\x6d\x61\x74\x63\x68\x7c\x6d\x61\x78\x7c\x6d\x69\x6e\x7c\x6f\x6e\x7c\x6f\x72\x64\x65\x72\x7c\x70\x61\x72\x65\x6e\x74\x7c\x70\x72\x6f\x74\x65\x63\x74\x65\x64\x7c\x70\x72\x6f\x76\x69\x64\x65\x7c\x70\x75\x62\x6c\x69\x63\x7c\x72\x65\x71\x75\x69\x72\x65\x7c\x72\x65\x74\x75\x72\x6e\x68\x6f\x6d\x65\x7c\x73\x6b\x69\x70\x7c\x73\x70\x6c\x69\x74\x5f\x74\x68\x72\x65\x61\x64\x7c\x73\x75\x6d\x7c\x74\x61\x6b\x65\x7c\x74\x68\x72\x65\x61\x64\x7c\x74\x6f\x7c\x74\x72\x61\x69\x74\x7c\x74\x79\x70\x65\x7c\x77\x68\x65\x72\x65\x7c\x77\x69\x74\x68\x7c\x79\x69\x65\x6c\x64\x7c\x79\x69\x65\x6c\x64\x68\x6f\x6d\x65\x29\x5c\x62\x2f\x69\x5d\x2c\x0a\x5b\x22\x74\x79\x70\x22\x2c\x2f\x5e\x28\x3f\x3a\x61\x72\x72\x61\x79\x7c\x64\x61\x74\x65\x7c\x64\x65\x63\x69\x6d\x61\x6c\x7c\x64\x75\x72\x61\x74\x69\x6f\x6e\x7c\x69\x6e\x74\x65\x67\x65\x72\x7c\x6d\x61\x70\x7c\x70\x61\x69\x72\x7c\x73\x74\x72\x69\x6e\x67\x7c\x74\x61\x67\x7c\x78\x6d\x6c\x7c\x6e\x75\x6c\x6c\x7c\x62\x6f\x6f\x6c\x65\x61\x6e\x7c\x62\x79\x74\x65\x73\x7c\x6b\x65\x79\x77\x6f\x72\x64\x7c\x6c\x69\x73\x74\x7c\x6c\x6f\x63\x61\x6c\x65\x7c\x71\x75\x65\x75\x65\x7c\x73\x65\x74\x7c\x73\x74\x61\x63\x6b\x7c\x73\x74\x61\x74\x69\x63\x61\x72\x72\x61\x79\x7c\x6c\x6f\x63\x61\x6c\x7c\x76\x61\x72\x7c\x76\x61\x72\x69\x61\x62\x6c\x65\x7c\x67\x6c\x6f\x62\x61\x6c\x7c\x64\x61\x74\x61\x7c\x73\x65\x6c\x66\x7c\x69\x6e\x68\x65\x72\x69\x74\x65\x64\x7c\x63\x75\x72\x72\x65\x6e\x74\x63\x61\x70\x74\x75\x72\x65\x7c\x67\x69\x76\x65\x6e\x62\x6c\x6f\x63\x6b\x29\x5c\x62\x7c\x5e\x5c\x2e\x5c\x2e\x3f\x2f\x69\x5d\x2c\x5b\x22\x70\x6c\x6e\x22\x2c\x2f\x5e\x5b\x61\x2d\x7a\x5f\x5d\x5b\x5c\x77\x2e\x5d\x2a\x28\x3f\x3a\x3d\x5c\x73\x2a\x28\x3f\x3d\x5c\x28\x29\x29\x3f\x2f\x69\x5d\x2c\x5b\x22\x70\x75\x6e\x22\x2c\x2f\x5e\x3a\x3d\x7c\x5b\x2d\x2b\x2a\x5c\x2f\x25\x3d\x3c\x3e\x26\x7c\x21\x3f\x5c\x5c\x5d\x2f\x5d\x5d\x29\x2c\x5b\x22\x6c\x61\x73\x73\x6f\x22\x2c\x22\x6c\x73\x22\x2c\x22\x6c\x61\x73\x73\x6f\x73\x63\x72\x69\x70\x74\x22\x5d\x29\x3b\x0a")

func init() {
	filepath := "/vendor/prettify/lang-lasso.js"
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
	_, err = f.Write(lang_lasso_js)
	if err != nil {
		log.Fatal(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}
