// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package asset

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 1, 31, 8, 20, 53, 756344200, time.UTC),
		},
		"/_themes": &vfsgen۰DirInfo{
			name:    "_themes",
			modTime: time.Date(2019, 1, 31, 8, 7, 54, 657958300, time.UTC),
		},
		"/_themes/github": &vfsgen۰DirInfo{
			name:    "github",
			modTime: time.Date(2019, 1, 31, 8, 56, 56, 7876200, time.UTC),
		},
		"/_themes/github/files": &vfsgen۰DirInfo{
			name:    "files",
			modTime: time.Date(2019, 2, 4, 5, 49, 57, 400782900, time.UTC),
		},
		"/_themes/github/files/github.css": &vfsgen۰CompressedFileInfo{
			name:             "github.css",
			modTime:          time.Date(2019, 1, 31, 8, 29, 50, 533766900, time.UTC),
			uncompressedSize: 13510,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x3a\x69\xb3\xa2\x48\xb6\x9f\xab\x7e\x05\xef\x76\x4c\x4c\xd5\x78\xbd\x02\xe2\x5a\xd1\x1d\x93\x28\xee\xe8\xf5\xba\x62\x47\xbf\x88\x04\x12\x48\x05\x12\x01\x45\x9d\xe8\xff\xfe\x02\x5c\xae\x0b\xa8\xd5\x33\x31\xcf\x0f\x86\x24\xe7\x9c\x3c\x5b\x9e\x2d\xfd\xa7\x46\x6c\x3f\xad\x41\x05\x51\xff\xfa\x4a\x51\x87\x27\x0b\x9b\xdb\x32\x45\x14\x1f\x2b\xc4\xf6\xd2\x26\xb6\x17\x3f\xbe\x52\x94\xe7\x2a\x65\x6a\xe5\x9a\xdf\x54\xe8\xc3\x72\x08\x9b\x09\x88\xa6\xfd\x50\x0c\xe8\x7a\xc8\xff\x75\xe5\x6b\xe9\xe2\x0f\x19\x7a\x28\xcf\xbd\xaa\x74\xa9\xfe\xa1\x03\x1e\x44\x9f\x59\x70\xfc\x55\xa9\xf5\xc1\xfd\x0f\x2f\x8c\x68\xb3\x01\x00\xa8\xc3\xe8\x59\x0f\xbf\x9a\xd1\xcf\x91\x3a\x1c\x8f\xc2\x9f\x33\x25\xa2\x15\xbd\x22\x00\xf4\x01\x18\x32\xe2\x5a\x0c\x9f\xb7\x21\x7d\xbe\x15\xbe\x91\x6a\x92\x30\xca\x7e\xcc\xe5\x49\x2d\x00\x00\x54\x23\x24\x61\x14\x62\x02\xd0\x1a\x1a\x6b\xab\xcb\xaa\x95\x70\x71\xb8\x88\x76\x8e\x36\xd9\xf3\x37\xb3\x79\x5b\x0e\x7f\x16\x42\xa2\x95\x6d\xf8\xba\x32\x12\x53\xc5\xa6\x61\x4b\xd3\x6e\x48\xaf\x3e\x3c\x47\xe2\x09\x68\xb2\xaa\x87\x26\x12\x00\xa0\xe6\x85\x6f\xde\xf7\xa2\x2b\xc2\xb2\xe4\xc3\xfa\xd8\x98\x85\xcf\x5e\xc4\x34\x1d\x7e\x75\x75\x83\x83\x25\x86\x84\xef\x42\xfe\x22\x56\x78\x12\x7e\x2f\x2a\xa0\x58\xfd\xa8\x1b\xbe\xda\x08\xe1\x3b\xc5\x70\xb1\x1a\x6d\x15\xd4\x01\xa8\x68\x72\xbd\x34\x97\x42\xfe\x3c\x70\xd2\x4f\x05\xf0\x18\x0c\x79\x5f\x9a\x1a\x21\x7f\x95\x65\x44\x4f\x3f\x28\xb1\x08\x06\x56\xce\x90\x27\xa1\xfc\xc3\x70\x13\xbe\x1f\xbe\x32\xa7\xab\x42\xd6\xeb\x29\xf5\xd2\x4e\x0d\x17\x71\x88\x0a\x50\xf8\x25\x0a\xd9\xbe\xd3\xe3\x03\xc5\x1a\x87\x8b\x82\x1c\x2e\x36\x42\xf9\xf8\x0c\xac\xd5\x9d\x45\x76\x0e\x87\x52\x7e\x03\x8b\x2d\x50\x17\x27\x99\x5e\x9e\xe5\xab\x98\xf6\x5b\x1d\xa9\x8f\x6d\x65\x2a\x6c\x1d\xa9\x89\xeb\xad\xf9\x40\x6f\xd8\xb8\x9f\x5f\x59\x43\x6f\x24\x6c\x3b\x56\x8e\x1f\xe7\xbb\x55\xfe\xbd\x38\x74\x7c\x2f\x5f\xa3\xd7\xa9\x45\x86\x86\x36\x8b\x53\xd8\x6f\x54\x83\xec\x9a\x4d\x95\x52\x55\xbe\x3d\xdc\x79\xad\xae\x3d\x69\x75\x87\x7a\x43\xd8\x72\xbc\x5e\xcf\x0a\x62\xb3\x54\xad\x08\xd5\x5e\x5d\x98\xee\xaa\xa0\x3a\xca\x19\x7c\x5b\x6c\xea\xdd\xf7\xd9\x92\x54\xb3\x03\x6c\x8e\xe1\x74\x56\x11\x3e\xb2\x99\x66\x01\xf8\x1b\xa1\xd5\xf1\x77\xbb\xd5\x4c\x6b\xa6\xc6\xe3\x85\xe3\x6e\x86\xe6\x74\x60\x4c\xda\x72\x76\xc8\x23\xa5\xce\x30\x6e\x40\xba\xa6\x65\xd9\xcc\x3b\x3b\x91\x94\x96\xb2\x33\xb3\x2c\xf2\x07\x4e\xdb\xde\xe1\x4a\xc1\xec\x6f\x27\x88\xf1\xac\xf1\xfb\x36\xd3\xf1\x0b\x6d\x25\x45\xaf\x27\x52\x46\x07\x7a\xb3\x29\x2c\x41\xb7\x14\x20\xda\x09\xda\x53\x17\x61\x11\x7a\x9b\x35\x94\xab\x7d\x51\xe4\x5c\xdc\x4b\x2d\x37\x22\x4b\xf4\xa0\x5a\xef\xcd\x86\xd3\x4d\xb0\xa9\xe2\xad\xd2\x6f\x2a\x44\xaa\xf1\x9d\x79\xae\x9d\x15\x9a\x70\xa0\xf8\x60\xc9\x2e\x86\x12\x0e\x52\x5b\xcb\x50\x50\x61\x1d\x88\xa5\xf9\x60\xd9\x2b\xb6\xb6\x63\x35\xf7\xd1\x28\xe9\xdb\xa1\xcf\xa6\x5a\x99\xed\xc8\x92\xcc\xe6\x07\xed\xd1\x9c\x9d\x4f\x15\xc6\x16\x43\x76\x68\x37\x42\x1d\x01\x8e\xe6\x06\xac\x0e\x56\xd3\x46\x30\xfe\xd0\xd7\x9d\x96\xcd\xf8\xfd\xc2\x06\xaf\xc6\xeb\x0c\x51\x86\x1f\x35\x8e\xb5\xba\xfa\xac\xce\xeb\x52\x5d\x0e\x66\x3d\x1e\x03\x50\xab\xb7\xf8\xa6\x08\x00\xde\x81\x5a\xe4\x0a\x18\xd4\x9b\x60\x67\xcf\xa1\xc4\xf2\x0b\xa9\x0e\x00\x87\xed\xe2\x2e\x98\xe2\xd4\x84\x4d\x89\xf3\xca\x4e\x6c\x56\x81\x33\x08\xd6\xd3\x5d\xa5\x54\x98\x71\x4d\xbd\xd8\xcd\xf0\x1b\xa9\x3e\xd3\x15\xdd\xcc\xb1\x7c\x65\xd0\x6f\x03\x90\x9d\x57\xc6\xc5\x0a\x00\xbc\x06\x8e\x67\x49\xe0\x4f\xfb\x73\x5a\x76\x0d\x2a\xfd\x59\x1f\xf0\x4d\x71\xde\xd6\x81\x25\x81\xb6\xa0\xf3\x53\x1d\x00\xd4\x75\xe6\x52\x5d\xca\x07\x43\xcc\xeb\xb3\x09\xaf\xb3\x0b\x6b\xb4\x21\x55\xc0\x89\xef\x46\x7d\x26\x4a\x3b\x1e\x33\xa0\xb1\xd5\xc7\x1d\xa9\x3f\xaa\x40\x18\x2c\xab\x80\x7b\xaf\x18\x1b\xc3\x32\x32\xc5\x5e\xb5\x2a\x78\x6b\x10\x34\x74\xb1\x2d\x36\xab\x76\xbd\x43\x6f\x0a\x7a\xab\x5f\x01\x81\x08\x5a\x2a\x27\x46\x11\xa0\x11\xc9\xa7\x4b\x75\x08\xb8\x2a\xa8\x7f\xe8\x52\x7f\x01\x1a\xdb\xba\x58\x2b\x76\x75\xc9\x6d\x8a\xd9\x56\x13\xd4\xc7\x92\x54\x1d\xa6\x80\x30\x07\xc1\xaa\x5a\x73\x78\x0b\x94\xda\x62\x55\x08\xc4\x8a\x51\xc2\x99\x75\xb1\x51\xf4\x1a\x74\x86\x53\xfb\x0a\x83\x81\x05\x16\xa0\x03\x47\xed\x8e\xbe\xa7\x3f\x94\x4a\x9d\xaa\xd7\xd4\x85\xa6\xec\xeb\xcb\xc6\xe8\xdd\xa9\xe2\xac\xfe\x4e\xf8\xf1\xf6\x63\x68\x0d\x55\xb5\x67\x2d\x87\xd3\xa1\x21\x4c\x97\x2e\x91\x59\xbd\xcf\xd4\xe6\x81\x53\x5d\x6b\x41\x85\x57\x2d\x75\x5a\xc9\x81\x71\xbb\xb6\xca\xa2\x9c\xa8\x75\x6b\x2d\xb6\xd4\x1e\xf6\x87\x5c\xb1\x27\x97\x32\xe6\x52\x0a\x7a\xf5\xd9\x06\x8d\x90\xd9\x65\x47\xec\x47\x3e\xa5\x00\x57\xf7\x2b\x2d\x07\xae\x26\x85\x51\x9f\x5f\xda\xb5\xc5\xc8\x9b\x03\x29\xb3\xe8\x8d\x18\xe5\x3d\x55\x05\xfa\x7a\x13\xd8\x8c\x62\xcc\xaa\xc1\x48\x56\xf3\x95\x1a\xb6\xea\xd3\x60\x17\xd4\xf2\xfe\xbb\x5c\x6b\x2a\x73\xc1\x4c\xad\xd7\x96\x98\x91\xb7\x80\x2b\xa2\xbc\x3f\x71\xdb\xc0\xb5\xb8\x59\xcb\xac\xc8\xaa\xe7\x6e\x16\x5e\x87\x01\xc1\xc4\xce\x6c\xf9\x41\xab\xed\x48\xf2\xb2\x08\xa6\x10\x0e\xe5\x62\x24\x2f\x5b\x9c\x83\xa0\x57\xa1\xe9\x99\xcb\xa3\x7e\xb7\xda\xef\x4d\x7a\x99\x8c\xa7\xf2\xa1\x9a\x3f\xb0\x34\x91\x80\x20\x74\x84\x40\x1c\x0a\xdc\x6a\x47\x72\xb3\x1d\xc9\xc9\x2c\xbf\x51\xed\x5a\x4f\x01\x9d\x4d\x77\x0e\xf2\x32\xcb\x6f\x87\x5e\x50\x29\xce\xa5\x40\xa7\xc7\x66\x77\x45\x2a\xc3\x09\x10\x97\xdd\x9d\xb8\xf3\x48\x9b\x71\x05\xa3\xbb\xe4\xb7\xc2\x16\xb9\x7a\xee\x5d\x6c\x99\xd2\x6a\xbc\x42\xc2\xb0\xad\xa8\x99\x62\x69\xc5\x3b\xb6\xb3\x6e\x0a\x63\x62\xa1\x46\x87\x88\x1e\x00\x88\x69\xaa\xdc\x31\x9f\x70\x2c\x99\xf4\x87\x74\xa1\xd2\xe7\x87\xf5\x35\xdd\xe2\x0d\xa8\x2f\x0a\x8d\xfe\xae\xbd\x51\x20\xeb\xb5\x2a\x02\x63\x54\x7d\xae\x5f\x4b\x95\x5a\xbd\x01\x6d\xcb\x10\x4a\xd5\x4a\x5f\x0b\x2a\xad\x02\x58\x65\x41\x63\x9e\xea\xf4\x98\x6c\x4d\xb4\xac\xbc\x62\x16\x0a\xc5\xdc\x7a\x8d\x6c\x7a\xc1\xcf\x1b\x15\xde\xd0\x1c\x69\xd5\x85\xb9\x77\x83\x51\x68\xc4\x4e\x57\xd9\xb9\xb0\x9e\xd4\x0b\x23\xf5\xbd\xda\x99\x71\xdd\x12\x6b\xf7\xac\x94\xc0\x4f\x57\x40\x6e\x58\x4d\x71\xf0\x21\x7a\x29\x0e\x8e\x04\x95\xeb\xaa\xd9\x4a\xa3\x5a\xec\xaa\xeb\x5e\x67\xe8\x01\xb6\xde\x29\x8a\xa5\xf7\x5e\x55\x56\x3a\x29\xa3\x5a\xa8\x30\x1b\x02\x1b\xa8\xd3\x1a\x08\x90\xd0\x35\x61\xc2\x70\xca\x62\x53\x49\x0d\x47\xc5\xe1\x66\xed\x49\xf9\x29\x8d\x3a\xef\xd6\x87\xe1\x6e\xd9\xc9\x18\x93\x8e\xb3\x70\x65\xa7\xc8\x75\x3a\xfd\xf7\x7a\xb3\xa0\xe4\xbd\x1e\x1e\xed\x9c\x49\x73\x32\xc8\xd5\x77\xe6\x40\x1f\xed\x76\x1d\x7e\x80\x17\xbd\xf7\xda\xb0\x37\x5d\x9a\xdb\x82\xbb\xdc\xd0\x33\xa6\x9f\xe3\x41\x93\xcc\xf8\x41\x0d\x1b\x7d\xa9\xdf\xeb\xf1\x82\xba\xa8\xf4\xf4\xe9\xb0\xd7\x00\x74\xa1\x01\xea\xf3\xfa\x04\x37\xe7\xf0\x7d\xd6\x9d\x30\xd9\x4c\xca\xb4\xf2\x83\x52\x6d\x58\x70\x3b\x8d\x5a\x2b\xaf\xf5\xe5\x05\x18\xf6\xea\xcc\x9c\xed\xd5\xc4\x95\xd2\x6e\xb5\xbc\x4d\x73\xac\xf5\x7b\x1f\x66\xaa\xd4\xda\xaa\x30\x3f\x30\x19\x75\x24\x19\x83\x8a\xc5\xa8\xdb\x8a\xa9\x11\x54\x5d\x23\x6e\x29\x4a\x6a\x47\x90\xb5\x65\x43\xcb\xf6\x32\x40\xad\xae\x2c\x6f\x1e\xd9\xcb\xea\xea\x12\x01\x60\xd6\x97\xe6\xbc\xb5\x05\x75\xa9\x3f\xb3\x54\xa3\x53\xdc\x75\xd4\xaa\xb0\x55\xc1\x87\x46\xc0\xb2\x79\xc8\xbe\x22\xe0\x03\xd0\x06\xbc\x08\xf8\x4c\x26\x13\xe6\x39\x70\x5b\x62\x1c\xaa\x8f\x5f\x7f\xfd\x4e\x69\xc4\xb5\xa0\xff\xed\xef\x61\xed\xf2\xf7\xef\x3f\xbe\xfe\xf9\xf5\xeb\x9b\x05\xdd\x85\x4a\x02\x3b\x2d\x13\x75\x1b\x95\x42\x69\xcb\x4b\xfb\x68\xe3\xa7\x3d\xbc\x43\x69\xa8\xce\x57\x9e\x5f\xa6\x18\x9a\xfe\x5b\x58\x0a\xa5\x03\x24\x2f\xb0\x7f\x07\xc2\xc4\x36\x4a\x1b\x08\xeb\x46\xb8\xf8\x96\x0b\xd7\x14\x62\x12\xb7\x4c\xfd\xc2\x72\x6c\x89\x45\x3f\xae\x0b\xae\x34\x74\x1c\x13\xa5\xbd\xad\xe7\x23\xeb\x95\xe2\xc3\xba\x4b\x84\xca\x20\x7a\xae\x11\xdb\x7f\xa5\x5e\x06\x48\x27\x88\x1a\x35\x5f\x5e\xa9\x06\x32\xd7\xc8\xc7\x0a\x7c\xa5\x80\x8b\xa1\xf9\x4a\x79\xd0\xf6\xd2\x1e\x72\xb1\xf6\x4a\xbd\x80\x90\x18\x55\x09\xb7\xa4\x04\x8b\xcc\xf1\xcb\x19\x7a\xcc\xca\x60\x6b\xc9\xc4\x7c\x39\x71\x15\x4a\x55\xa6\x98\xbc\xb3\x49\x12\x27\x20\xae\x9a\x0e\x5c\xe8\x94\x29\xd9\x45\x70\x91\x0e\x17\xe2\xf4\xf9\xe6\x98\x69\x25\xd2\xea\x51\x03\x79\x58\xc8\x16\x92\x61\x99\xd7\xb8\x65\x2f\xfa\x5e\x5f\x10\xa2\xe9\x9c\xa2\xe4\x92\x08\xa1\x58\x3a\xc8\xbe\xe4\x45\xe3\x58\x85\x49\x22\xe1\x59\xf8\x0e\x33\x1e\x73\x41\xea\x68\xd8\x04\x6e\x6c\xff\x12\x9a\x2d\xe6\xb3\x30\x09\x7a\x71\x01\xab\x16\xb2\x90\x2b\x25\x32\x19\xcb\xa2\xa3\xc6\xaf\xef\x59\x77\x3c\x74\x10\x21\x1e\xc8\x4d\x58\xde\x5b\x48\x89\x57\xed\xe1\xb5\xe7\x3e\x78\x0d\x2f\x8d\x98\x65\xb5\x3c\x9b\x24\xdc\x3a\x9e\x94\x15\x5c\xd0\x40\xd9\x3c\x4b\x27\x2a\x48\x5e\x5d\x00\xcb\x59\x46\x65\x8b\x49\xc0\x18\x5f\x00\x6b\x50\x93\x35\x25\x74\x78\x19\x2a\x0b\xdd\x25\x2b\x5b\x4d\x3f\x47\x49\x61\x9f\xa7\x74\xdf\xc2\x0a\x5b\x2e\xcb\x48\x23\x2e\x3a\x50\xb4\x7d\x64\xfb\x65\xea\xe5\x7f\xc5\x97\x44\xaf\x38\xd9\xea\xb3\xb9\x0b\x0e\x47\x58\x26\xa6\xfa\xe3\x69\x5f\xb4\xcc\x0b\x39\x0a\xd9\x9c\x42\x6b\x89\xc0\x46\xac\xc1\x2c\xe3\x70\x08\xe2\xdf\x7a\x8f\x79\xbc\x7f\xd6\x2d\xfc\x49\xc1\xf3\xb7\x26\x2a\x53\xd8\x87\x26\x56\xe2\x22\x6f\x02\x09\xf9\x09\x45\xdd\xa7\xa0\xc6\xfa\x59\xac\xc1\x35\x0d\x21\x8d\x4e\x16\x87\x89\x0d\x16\xf1\xa4\x68\x4d\xd3\xb8\x44\x52\x4a\xec\x49\x49\x62\x4a\x56\x93\xfd\x00\x5f\x39\x74\x5e\x2b\x6a\x49\x4c\x3d\x30\x97\xea\x3e\x56\xf6\xfd\xd0\x2c\x5f\x46\x91\x5c\x31\x4f\xe7\x93\x43\xa4\x7e\x01\x5c\xca\x95\x54\x98\xc8\x9c\x42\xdc\xbd\xc7\x47\xf9\x5d\x45\x0a\x71\xa1\x8f\x89\x5d\xa6\x56\xb6\x8a\xdc\x30\x19\xfe\x78\x2e\x7e\x1d\x86\x28\x11\x31\x15\x7b\x8e\x09\xb7\x65\x0a\xdb\x51\x3a\x95\x4d\xa2\x44\x83\x95\x35\x72\xc3\x3c\x6e\xa6\xa1\x89\x75\xbb\xbc\xdf\xd5\x27\x4e\x94\x8b\xb1\x69\x96\x29\x65\xe5\xba\xc8\xf6\xa3\x6c\x1e\xb7\xcf\x5e\x13\xb7\x46\xf0\x5d\x68\x7b\x0e\x0c\x71\x63\xd1\xca\x50\xf1\xf1\xfa\x36\x58\xc3\xb2\x41\xd6\x68\x6f\x21\xb2\xf2\x23\x76\x03\xac\xfa\x46\x99\x8a\xf5\x58\xcf\x77\x89\xad\xdf\x1a\x14\xdb\x06\x72\x71\xec\xde\x49\x28\xa1\x0f\xa0\x58\x21\x0d\xe6\xec\x84\x47\xd5\x09\x8b\xac\x50\x47\x16\x74\x75\x6c\x97\x29\xfa\x2d\x5f\x40\x56\x3c\x87\xd8\xda\xef\x25\x13\x57\x45\xee\x31\x42\xd8\xc4\x8e\x3d\xcb\x0a\x51\x6f\x95\xb2\x90\xd5\x9b\x35\xc7\x8d\x99\x9c\x59\xc4\x26\x9e\x03\x15\xf4\xfa\xf9\xf3\xba\xae\x0a\x39\x8f\x11\xd1\x3d\x30\xb9\x09\xe1\xb0\xad\x97\x8f\xa1\x3e\x2d\x93\xa8\x10\x3b\xd6\x60\x74\xf8\x10\xda\x48\x33\x49\x50\xa6\xd6\xd8\xc3\xb2\x19\x2b\x0b\xb6\x9d\x95\x7f\x62\xf2\xcc\x26\x67\x8a\xbb\x8f\xf7\xd4\x36\xbf\xfb\x5b\x07\xfd\xfa\xa2\x18\x48\x59\xc8\x64\xf3\xf2\xc7\x8d\x24\x07\xd5\x1f\x04\x71\xa0\xaa\x46\xcb\xb1\x9b\xff\xe3\x2e\xf6\x23\x21\x4f\x96\x38\x93\xf5\x4c\xf9\x67\xab\x17\x75\xed\x1d\x6f\xbd\x2e\x58\xf2\x79\x35\xff\x23\x2e\x40\x24\x79\xd4\xf9\x89\xba\x17\x55\x9e\x3f\x28\x79\x3a\x56\x73\x07\x17\xba\xf0\x93\xa3\x9d\x99\x9c\xb3\xb9\xf6\x1c\x03\xab\x2a\xb2\x2f\x63\xf8\x55\xe0\x38\x9e\x9a\x03\xb5\x93\x25\x7c\x9f\x58\x65\x8a\x71\x36\x94\x47\x4c\xac\x52\xbf\xa8\x1a\x62\x51\x6c\x54\x35\xdc\x8b\xfa\xe5\x14\x0a\x7d\x18\xb9\xd3\x79\x45\x13\x5b\xcf\x84\xf8\x50\xf3\x0f\x0a\x8c\x41\x37\x11\x74\x43\x27\xf1\x8d\x27\xa8\x45\x68\x17\xe1\xc0\x81\xca\xd1\x19\x4f\x8b\x0a\x31\x4d\xe8\x78\x28\x3c\x83\xfb\x5f\xb1\xb4\x6e\x83\x82\x6f\x44\xb4\xef\x7b\xb8\x71\x5b\x77\x1b\xec\xed\x52\xf6\x76\x89\xbb\x5d\xca\xdd\x2e\xe5\x23\x1e\xf6\x96\x0f\x93\xc9\x85\x2b\x9c\x8c\x97\xc0\xda\x75\xa0\xcd\xb2\xfb\x36\xf0\x29\x07\x64\x6f\xc2\x34\xf7\x13\xd8\xd9\x1b\x6c\xfa\x27\xb0\xb9\x6b\xec\x63\x03\xfb\x14\x76\xee\x06\xfb\x67\x38\xcf\xdf\x60\xff\x84\xd6\x9c\xa7\xcc\xc5\x44\xba\xb8\xc5\x8e\x4a\x89\xe5\x8a\xf8\xe8\x8c\x4c\x82\x75\x57\xe6\x8d\xb3\x10\xf3\xdc\x61\xd3\x26\xd2\x2e\x23\xc7\x4f\xfa\x0f\x31\x29\x72\xbb\xc9\xca\x3c\xee\x63\x62\xef\x50\xa3\xa7\xc3\xa4\x51\xa6\x4c\x12\x20\x37\xed\x12\x0b\xda\xf1\x1c\xef\x91\x13\x48\xc6\xbc\x20\x49\x18\xe4\x80\x71\x8f\x0f\x68\x3a\x46\x6c\x09\xac\xaa\xe7\x46\x3a\x69\x29\xbe\x7c\xb8\x4d\x46\x2f\x83\x9a\x48\x6c\x92\xfe\x40\xfa\xca\x84\xee\xcb\x2b\x55\x21\xb6\x47\x4c\xe8\xbd\x52\x2f\x1d\x2c\xa3\x7d\x3a\xa0\x42\xa0\x97\x57\x4a\x44\xb6\x49\x42\x98\x95\x8b\x91\x7b\x51\x4a\xc4\xb8\x8f\x8b\x6e\x1c\x88\x39\xaf\x8e\x3e\x5d\x68\xbf\xfa\x5f\xe3\xec\xa2\x04\x8e\xad\x74\xf7\x9c\x25\xd5\xe2\x74\x9c\x6b\x52\xff\x83\x2d\x87\xb8\x3e\x8c\xaf\x6b\x43\x3c\x26\x06\x8f\x73\x36\x4f\x60\xb2\x31\x98\xc5\xa7\x30\xb3\x31\x98\x61\x00\x7a\x02\x95\x8b\x41\x65\x9f\xe3\x37\x17\x83\x1a\x06\xec\x27\x50\xf3\x71\x4a\xa2\x1f\xa1\xfe\x5b\xa9\xfc\x3f\x98\xc7\x7f\xfb\x47\x59\xc3\xae\xe7\xa7\x15\x03\x9b\xea\x6d\x00\xbd\x2f\xc5\x6f\xff\x28\x9b\x30\x0e\xf9\x14\xdb\x1e\x28\x10\x96\x6d\xe2\x7f\xfb\xdd\x70\x91\xf6\xc7\xf7\xf3\xfa\xf0\xac\xbe\x7c\xba\x3e\x7c\x83\xb6\x62\x90\x43\x4f\x6c\x12\xe8\x97\xa9\xd0\x1c\x67\x95\x72\xda\xdd\x27\x90\x43\x4a\xba\x88\x42\xe9\x63\x92\xbc\x9c\xd3\xde\xd9\xa8\xac\x11\x65\xe5\x9d\x37\x78\xc9\xbc\x39\x37\x61\xf4\x33\xdf\xc4\xc4\xe4\x98\xa0\x7b\xb3\xa4\xde\x2e\x45\x3e\x90\xd8\x60\x3d\x19\xd8\x1e\xd7\xc3\x6f\x6c\x6e\x8f\x7e\x5e\x9e\x7d\xa6\xcc\xe8\xcc\xd1\x09\x23\x0d\xc4\x20\x0e\x15\xaf\x0a\xe2\x87\xd9\xf8\x73\xa3\x23\xe7\xd7\x63\xf0\x53\xc9\x79\x08\x6f\x11\x8f\x8f\x8b\xea\xcf\x7d\x7e\xbb\x7f\x10\x1e\x21\x3f\x38\x07\x71\xe8\x0b\x59\xbd\x3f\xd4\x38\x49\x9d\x75\x36\x54\xee\xac\x10\x3a\xd4\x46\x4c\x9c\xc3\x1e\xbc\xf8\xa8\x1f\x8e\xe3\xd4\x5c\x3e\x6e\x44\x62\x61\x55\xdd\x47\x8c\xb8\x21\xd6\xe7\x9c\xf5\x60\xa6\xbd\x2a\xc3\x4e\xe5\x17\x25\xaf\xc8\x2a\x73\xd3\xc7\xa4\xaf\x27\x43\xa7\xf7\x2e\x54\xf1\xca\x8b\xe4\xf8\x71\xec\x49\x0d\xa8\x86\xbd\x13\xb6\x3d\xe4\x53\x34\x95\x0e\x29\xd3\xf7\xa6\x4a\xff\xfd\x6a\x9f\xbd\x8c\x13\xa7\x53\x92\x58\x0f\xdf\x5c\xf2\xb0\x09\x92\x9c\x52\x7a\xf4\xcf\xa0\x18\xb9\x1e\x01\x64\x1f\x01\x70\x8f\x00\x72\x8f\x00\xf2\x97\x00\x17\xdd\x3b\x23\x33\x1a\x9b\xbd\xef\x55\xd1\xa0\x03\x9b\xd8\xdf\x7e\x76\xc8\x71\xba\x38\x74\xf5\x87\xa0\x1a\xa3\x8b\x47\x00\xd9\x47\x00\xdc\x23\x80\xdc\x23\x80\xfc\x25\x40\xfc\x08\x22\x29\xf6\x5f\x8b\xf8\xd0\xf8\x3f\x07\x9e\xfd\x39\x70\xee\xe7\xc0\x73\x3f\x07\x7e\xad\xa8\x5b\x17\x3a\x77\x8b\x3b\xb3\x30\xe3\xb2\xf8\x3c\x05\xd2\xb7\xec\x59\xf1\x7d\x31\xc2\x4c\x9e\xa9\x20\x88\x14\x14\x7b\xd9\x62\xb0\x4f\xef\xc2\xbc\xe5\xfe\x8d\x7d\x6e\x5a\x72\xe6\x90\x44\x9f\x6a\xc0\x13\x00\x6f\x7a\x6d\xfa\xad\x58\x48\xa2\x7a\xd3\x5a\xd3\x6f\xc5\x5c\x7c\x2e\xfd\xeb\x0d\x2f\x1b\xbf\x79\xd4\x7d\xc6\xb7\x9e\xf1\xed\x65\x42\x37\xfa\x97\xc7\x31\xe6\xfe\x66\xeb\xe6\xba\x1d\x9a\x66\x3c\xf8\x6f\xb7\xc3\x84\x7d\xd8\x8f\x03\x4e\x1d\xc8\x5f\xb0\x96\x68\x5f\xd5\x7c\x62\xb4\xa5\x9a\x94\xea\xdf\xc0\xc5\x31\x74\xeb\x29\x49\x37\x78\xcf\xcc\x50\xc2\x7d\xd5\x9b\x92\x2b\x7f\x27\x0f\xde\x1b\x0d\x9e\x2a\x9b\x53\x49\x73\xb8\x01\x39\xfe\xc5\xe3\x73\x86\x0a\x57\x3e\x49\x26\x76\x18\x07\x3e\x23\xc1\x11\x21\xbe\x38\xa6\xfc\x2b\xe9\xc2\xd6\x92\x39\xd5\x25\xfb\x52\xe7\xa9\x71\xec\x81\x9c\x9b\x70\x6d\xf4\x8b\xa6\x69\x67\xd1\x62\x6f\xb0\x4f\xba\xc7\x12\x2a\x99\x6e\xd9\xf6\x8d\x7d\x41\xf9\x8d\xb5\xbf\x27\xee\x72\xb8\x3e\x4c\xbe\xad\xb1\xe0\x26\x7d\xa5\xf5\x3b\x77\x23\x49\x82\xc4\xd2\xff\x3d\xca\xf9\xbf\x46\x4d\xd5\x1f\x71\xd1\x20\x61\xd4\xf6\x89\x1a\xc2\x5d\x62\x1e\x3a\xb4\x24\xd4\xd3\x58\xe8\xd3\x3d\xdf\x58\x64\x51\xf4\x1b\x77\x75\x91\x75\x75\x32\x8a\xb9\xbf\xc5\xcb\xe7\xea\x32\xfc\xc6\x16\x5e\xb3\xcc\x6b\x36\xf7\x4a\xbf\xd1\xb9\xef\x49\x95\x6b\xf2\xcc\xe8\x2c\xb0\xd8\xc4\xb5\x60\x6c\x54\x71\x5c\xf4\x5b\x0c\xff\x4f\x70\x1d\x91\x8f\x22\xd6\x27\x7d\x8a\x0a\x0c\xec\xa3\x68\xee\x8e\xca\x21\xf1\xe7\x2f\x20\x62\x5a\x5a\x03\xeb\x86\x19\xea\x3e\xae\x7d\x49\x3a\xe9\x67\x58\x57\x4d\xe6\xd5\xa8\x2c\xa6\xf5\x8c\x93\xe9\xd1\x0e\x89\xad\xed\x49\x9b\xc7\x50\x75\x1d\x58\x6e\x2a\x73\x2e\xb1\x3f\x89\xb7\xf2\xc9\x70\x57\xdd\xda\x8f\xcb\x23\x76\xdc\xed\x8e\x79\x63\xae\x02\x93\x2e\xd1\x62\x1d\xeb\xe1\x25\xf5\x23\x53\x6b\x2b\xd3\x4c\x2b\xc4\xb2\xb0\x4f\xbd\xc9\xbe\x9d\x3e\x4e\x2f\x6c\xe2\x7f\x2b\xab\xd8\x0b\x23\x90\xfa\xfd\xec\xb6\xed\xfa\xff\x08\xe7\xf7\x3b\x0f\xfe\xa8\xf0\x97\x9b\xdc\x7d\x7f\xfb\x1f\x9c\xab\xfe\x3f\x34\xca\x2a\xa3\xe6\x54\x98\xdc\x28\xdf\x34\xd2\x3f\xd1\x28\x27\x67\x90\x72\x74\x81\x8c\xd4\xd4\x5b\x48\x8e\xa4\x4d\x28\xa3\x43\xcd\x41\x3c\xbc\xef\x55\x5c\x64\x42\x1f\xaf\x23\xa9\x76\x69\x6c\xab\x68\x13\x0d\xbb\x6e\x2d\x7b\xb8\xa7\x8d\xf1\x23\x1f\x7a\x8b\x74\x74\x01\x80\x7d\x64\xc5\x5f\x07\x24\xce\xea\x2e\x91\x53\x71\xc4\xce\x83\x46\xc2\xc9\xbc\x46\xfb\xbc\xc3\x3e\x1d\xb9\x53\x82\x88\x06\x42\x69\xe6\x2d\xbf\x8f\x3f\x49\x56\xbe\xf3\x8f\x82\x38\x13\x22\x14\x8b\xf3\xaf\xaf\x5f\xbe\x24\x5d\xbd\x7f\xf9\x62\x61\xfb\x18\x30\x58\x3a\xf2\xc3\x2f\x5f\xce\x82\x48\xa9\x78\x5a\x3b\x0a\xb1\x8f\x2b\x5f\xbe\x9c\x4e\x0b\x17\x1d\x95\x2f\x7f\x7e\xfd\xfa\x4f\x0b\xa9\x18\x52\xdf\xce\x08\x14\xf2\x05\x67\x13\xd6\x0c\x5f\x62\xd8\xfa\x8c\x95\x47\x12\x7f\x7e\xfd\xfa\x7f\x01\x00\x00\xff\xff\x16\x1d\xb6\x70\xc6\x34\x00\x00"),
		},
		"/_themes/github/templates": &vfsgen۰DirInfo{
			name:    "templates",
			modTime: time.Date(2019, 1, 31, 8, 21, 29, 575097300, time.UTC),
		},
		"/_themes/github/templates/main.html": &vfsgen۰CompressedFileInfo{
			name:             "main.html",
			modTime:          time.Date(2019, 2, 4, 4, 48, 5, 19837200, time.UTC),
			uncompressedSize: 281,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x90\x31\x4f\xc3\x30\x14\x84\x77\xff\x8a\x47\xf6\x26\x2b\x83\xeb\xa5\xa5\x62\x40\x02\xd1\x30\x30\xbe\xc6\x07\xb1\x70\x62\xe4\x3c\x40\x55\xe4\xff\x8e\x8c\x1d\x26\xfb\xee\x3e\x9f\x4e\xd6\x37\xc7\xc7\x43\xff\xfa\x74\x47\xa3\x4c\xde\x28\xbd\x1d\x60\x6b\x94\x9e\x20\x4c\xc3\xc8\x71\x81\xec\x9b\x97\xfe\xb4\xbb\x6d\x8c\xd2\xe2\xc4\xc3\xac\x6b\xdb\xe7\x4b\x4a\xba\x2b\x8e\x5a\x57\x6a\xcf\x72\xf5\x38\x8f\x80\x2c\x94\x52\xb1\x86\xe8\x3e\xff\xa5\x7b\xa3\xf6\x88\xcb\xd7\x7b\xd6\x44\x19\x78\x70\xdf\x78\x86\x0f\x6c\x0b\x5a\x49\xcc\x76\xab\xb8\x07\x5b\xc4\xac\x74\x57\xc7\x5d\x82\xbd\x1a\xa5\x39\x8a\x1b\x3c\x68\xf0\xbc\x2c\xfb\x66\xe2\xf8\x61\xc3\xcf\xbc\xcb\x71\x53\x26\x1d\xc2\x2c\x98\xa5\xbc\xae\x7c\x49\x4e\x21\x08\xe2\x9f\x5f\xeb\xba\xf2\x03\xbf\x01\x00\x00\xff\xff\xbd\x4f\xa4\x37\x19\x01\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/_themes"].(os.FileInfo),
	}
	fs["/_themes"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/_themes/github"].(os.FileInfo),
	}
	fs["/_themes/github"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/_themes/github/files"].(os.FileInfo),
		fs["/_themes/github/templates"].(os.FileInfo),
	}
	fs["/_themes/github/files"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/_themes/github/files/github.css"].(os.FileInfo),
	}
	fs["/_themes/github/templates"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/_themes/github/templates/main.html"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
