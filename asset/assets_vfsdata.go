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
			modTime: time.Date(2019, 2, 6, 5, 9, 37, 291710100, time.UTC),
		},
		"/_themes": &vfsgen۰DirInfo{
			name:    "_themes",
			modTime: time.Date(2019, 2, 6, 5, 9, 37, 291710100, time.UTC),
		},
		"/_themes/github": &vfsgen۰DirInfo{
			name:    "github",
			modTime: time.Date(2019, 2, 6, 5, 9, 37, 291710100, time.UTC),
		},
		"/_themes/github/files": &vfsgen۰DirInfo{
			name:    "files",
			modTime: time.Date(2019, 2, 6, 5, 9, 37, 291710100, time.UTC),
		},
		"/_themes/github/files/github.css": &vfsgen۰CompressedFileInfo{
			name:             "github.css",
			modTime:          time.Date(2019, 2, 6, 9, 52, 52, 988178100, time.UTC),
			uncompressedSize: 13904,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x3a\xfb\x8f\xaa\x48\xd6\x3f\xdf\xfb\x57\xf0\xf5\x64\xb3\xf7\x8e\x6d\x0b\x88\xcf\x9b\x99\x6c\xe1\xdb\x16\x6d\x5b\x6d\xc5\xc9\x7c\x49\x01\x05\x94\x02\x85\x80\xa2\x6e\xf6\x7f\xff\x02\x3e\xda\x47\xa1\xde\xd9\xcd\x7e\xdd\x89\xd1\xe2\x9c\x53\xe7\x55\xe7\x55\xfc\x43\x27\x4e\x90\xd6\xa1\x8a\x98\x7f\x7e\x65\x98\xfd\x2f\x1b\x5b\x9b\x32\x43\xd4\x00\xab\xc4\xf1\xd3\x16\x76\xe6\x3f\xbe\x32\x8c\xef\xa9\x65\x66\xe9\x59\xdf\x34\x18\xc0\x72\x04\x9b\x09\x89\xae\xff\x50\x4d\xe8\xf9\x28\xf8\x6d\x19\xe8\xe9\xe2\x0f\x05\xfa\x28\x2f\x3c\x6b\x6c\xa9\xf1\x6e\x00\x11\xc4\x7f\xd3\xf0\xf0\xad\x52\xef\x83\xdb\x7f\x62\x6d\xc4\x5a\x4d\x00\x40\x03\xc6\xbf\x8d\xe8\xa3\x15\x7f\x1d\x69\xc3\x8f\x51\xf4\x75\xaa\xc6\xb4\xe2\x47\x04\x80\x3e\x00\x43\x4e\x5a\x49\xd1\xef\x4d\x44\x5f\x6c\x47\x4f\xe4\xba\x5c\x1b\x65\xdf\x67\xca\xb8\x1e\x02\x00\xaa\x31\x52\x6d\x14\x61\x02\xd0\x1e\x9a\x2b\xbb\xcb\x6b\x95\x68\x71\x38\x8f\x77\x8e\x37\xd9\xf1\x37\x75\x44\x47\x89\xbe\x16\x22\xa2\x95\x4d\xf4\xb8\x32\x92\x52\xc5\x96\xe9\xc8\x93\x6e\x44\xaf\x31\x3c\x45\x12\x09\x68\xf1\x9a\x8f\xc6\x32\x00\xa0\xee\x47\x4f\xde\x76\xa2\xab\xb5\x45\x29\x80\x8d\x0f\x73\x1a\xfd\xf6\x63\xa6\xd9\xe8\xa3\x6b\x98\x02\x2c\x71\x24\x7a\x16\xf1\x17\xb3\x22\x92\xe8\x73\x5e\x01\xc5\xea\x7b\xc3\x0c\xb4\x66\x04\xdf\x29\x46\x8b\xd5\x78\xab\xb0\x01\x40\x45\x57\x1a\xa5\x99\x1c\xf1\xe7\x83\xa3\x7e\x2a\x40\xc4\x60\x28\x06\xf2\xc4\x8c\xf8\xab\x2c\x62\x7a\xc6\x5e\x89\x45\x30\xb0\x73\xa6\x32\x8e\xe4\x1f\x46\x9b\x88\xfd\xe8\x91\x35\x59\x16\xb2\x7e\x4f\x6d\x94\xb6\x5a\xb4\x88\x23\x54\x80\xa2\x0f\xa9\x96\xed\xbb\x3d\x31\x54\xed\x8f\x68\xb1\xa6\x44\x8b\xcd\x48\x3e\x31\x03\xeb\x0d\x77\x9e\x9d\xc1\xa1\x9c\x5f\xc3\x62\x1b\x34\xa4\x71\xa6\x97\xe7\xc5\x2a\x66\x83\x76\x47\xee\x63\x47\x9d\xd4\x36\xae\xdc\xc2\x8d\xf6\x6c\x60\x34\x1d\xdc\xcf\x2f\xed\xa1\x3f\xaa\x6d\x3a\x76\x4e\xfc\xc8\x77\xab\xe2\x5b\x71\xe8\x06\x7e\xbe\xce\xae\x52\xf3\x0c\x0b\x1d\x1e\xa7\x70\xd0\xac\x86\xd9\x15\x9f\x2a\xa5\xaa\xe2\xeb\x70\xeb\xb7\xbb\xce\xb8\xdd\x1d\x1a\xcd\xda\x46\x10\x8d\x46\xb6\x26\xb5\x4a\xd5\x4a\xad\xda\x6b\xd4\x26\xdb\x2a\xa8\x8e\x72\xa6\xf8\x2a\xb5\x8c\xee\xdb\x74\x41\xaa\xd9\x01\xb6\x3e\xe0\x64\x5a\xa9\xbd\x67\x33\xad\x02\x08\xd6\xb5\x76\x27\xd8\x6e\x97\x53\xbd\x95\xfa\xf8\x98\xbb\xde\x7a\x68\x4d\x06\xe6\xf8\x55\xc9\x0e\x45\xa4\x36\x38\xce\x0b\x49\xd7\xb2\x6d\x87\x7b\xe3\xc7\xb2\xda\x56\xb7\x56\x96\x47\xc1\xc0\x7d\x75\xb6\xb8\x52\xb0\xfa\x9b\x31\xe2\x7c\xfb\xe3\x6d\x93\xe9\x04\x85\x57\x35\xc5\xae\xc6\x72\xc6\x00\x46\xab\x55\x5b\x80\x6e\x29\x44\xac\x1b\xbe\x4e\x3c\x84\x25\xe8\xaf\x57\x50\xa9\xf6\x25\x49\xf0\x70\x2f\xb5\x58\x4b\x3c\x31\xc2\x6a\xa3\x37\x1d\x4e\xd6\xe1\xba\x8a\x37\x6a\xbf\xa5\x12\xb9\x2e\x76\x66\xb9\xd7\x6c\xad\x05\x07\x6a\x00\x16\xfc\x7c\x28\xe3\x30\xb5\xb1\x4d\x15\x15\x56\xa1\x54\x9a\x0d\x16\xbd\x62\x7b\xf3\xa1\xe5\xde\x9b\x25\x63\x33\x0c\xf8\x54\x3b\xb3\x19\xd9\xb2\xd5\x7a\x67\x7d\x56\x70\xf2\xa9\xc2\x87\xcd\x91\x2d\xda\x8e\x50\xa7\x06\x47\x33\x13\x56\x07\xcb\x49\x33\xfc\x78\x37\x56\x9d\xb6\xc3\x05\xfd\xc2\x1a\x2f\x3f\x56\x19\xa2\x0e\xdf\xeb\x02\x6f\x77\x8d\x69\x43\x34\xe4\x86\x12\x4e\x7b\x22\x06\xa0\xde\x68\x8b\x2d\x09\x00\xbc\x05\xf5\xd8\x15\x30\x68\xb4\xc0\xd6\x99\x41\x99\x17\xe7\x72\x03\x00\x01\x3b\xc5\x6d\x38\xc1\xa9\x31\x9f\x92\x66\x95\xad\xd4\xaa\x02\x77\x10\xae\x26\xdb\x4a\xa9\x30\x15\x5a\x46\xb1\x9b\x11\xd7\x72\x63\x6a\xa8\x86\x95\xe3\xc5\xca\xa0\xff\x0a\x40\x76\x56\xf9\x28\x56\x00\x10\x75\x70\x38\x4b\x35\xf1\xb8\xbf\xa0\x67\x57\xa0\xd2\x9f\xf6\x81\xd8\x92\x66\xaf\x06\xb0\x65\xf0\x5a\x33\xc4\x89\x01\x00\xea\xba\x33\xb9\x21\xe7\xc3\x21\x16\x8d\xe9\x58\x34\xf8\xb9\x3d\x5a\x93\x2a\x10\xa4\x37\xb3\x31\x95\xe4\xad\x88\x39\xd0\xdc\x18\x1f\x1d\xb9\x3f\xaa\x40\x18\x2e\xaa\x40\x78\xab\x98\x6b\xd3\x36\x33\xc5\x5e\xb5\x5a\xf3\x57\x20\x6c\x1a\xd2\xab\xd4\xaa\x3a\x8d\x0e\xbb\x2e\x18\xed\x7e\x05\x84\x12\x68\x6b\x82\x14\x47\x80\x66\x2c\x9f\x21\x37\x20\x10\xaa\xa0\xf1\x6e\xc8\xfd\x39\x68\x6e\x1a\x52\xbd\xd8\x35\x64\xaf\x25\x65\xdb\x2d\xd0\xf8\x90\xe5\xea\x30\x05\x6a\x33\x10\x2e\xab\x75\x57\xb4\x41\xe9\x55\xaa\xd6\x42\xa9\x62\x96\x70\x66\x55\x6c\x16\xfd\x26\x9b\x11\xb4\xbe\xca\x61\x60\x83\x39\xe8\xc0\xd1\x6b\xc7\xd8\xd1\x1f\xca\xa5\x4e\xd5\x6f\x19\xb5\x96\x12\x18\x8b\xe6\xe8\xcd\xad\xe2\xac\xf1\x46\xc4\x8f\xcd\xfb\xd0\x1e\x6a\x5a\xcf\x5e\x0c\x27\x43\xb3\x36\x59\x78\x44\xe1\x8d\x3e\x57\x9f\x85\x6e\x75\xa5\x87\x15\x51\xb3\xb5\x49\x25\x07\x3e\x5e\xeb\xcb\x2c\xca\x49\x7a\xb7\xde\xe6\x4b\xaf\xc3\xfe\x50\x28\xf6\x94\x52\xc6\x5a\xc8\x61\xaf\x31\x5d\xa3\x11\xb2\xba\xfc\x88\x7f\xcf\xa7\x54\xe0\x19\x41\xa5\xed\xc2\xe5\xb8\x30\xea\x8b\x0b\xa7\x3e\x1f\xf9\x33\x20\x67\xe6\xbd\x11\xa7\xbe\xa5\xaa\xc0\x58\xad\x43\x87\x53\xcd\x69\x35\x1c\x29\x5a\xbe\x52\xc7\x76\x63\x12\x6e\xc3\x7a\x3e\x78\x53\xea\x2d\x75\x56\xb3\x52\xab\x95\x2d\x65\x94\x0d\x10\x8a\x28\x1f\x8c\xbd\x57\xe0\xd9\xc2\xb4\x6d\x55\x14\xcd\xf7\xd6\x73\xbf\xc3\x81\x70\xec\x64\x36\xe2\xa0\xfd\xea\xca\xca\xa2\x08\x26\x10\x0e\x95\x62\x2c\x2f\x5f\x9c\x81\xb0\x57\x61\xd9\xa9\x27\xa2\x7e\xb7\xda\xef\x8d\x7b\x99\x8c\xaf\x89\x91\x9a\xdf\xb1\x3c\x96\x41\xad\xd6\xa9\x85\xd2\xb0\x26\x2c\xb7\x24\x37\xdd\x92\x9c\xc2\x8b\x6b\xcd\xa9\xf7\x54\xd0\x59\x77\x67\x20\xaf\xf0\xe2\x66\xe8\x87\x95\xe2\x4c\x0e\x0d\xf6\xc3\xea\x2e\x49\x65\x38\x06\xd2\xa2\xbb\x95\xb6\x3e\x79\xe5\xbc\x9a\xd9\x5d\x88\x9b\xda\x06\x79\x46\xee\x4d\x6a\x5b\xf2\xf2\x63\x89\x6a\xc3\x57\x55\xcb\x14\x4b\x4b\xd1\x75\xdc\x55\xab\xf6\x41\x6c\xd4\xec\x10\xc9\x07\x00\x71\x2d\x4d\x38\xe4\x13\x81\x27\xe3\xfe\x90\x2d\x54\xfa\xe2\xb0\xb1\x62\xdb\xa2\x09\x8d\x79\xa1\xd9\xdf\xbe\xae\x55\xc8\xfb\xed\x4a\x8d\x33\xab\x81\xd0\xaf\xa7\x4a\xed\xde\x80\x75\x14\x08\xe5\x6a\xa5\xaf\x87\x95\x76\x01\x2c\xb3\xa0\x39\x4b\x75\x7a\x5c\xb6\x2e\xd9\x76\x5e\xb5\x0a\x85\x62\x6e\xb5\x42\x0e\x3b\x17\x67\xcd\x8a\x68\xea\xae\xbc\xec\xc2\xdc\x9b\xc9\xa9\x2c\xe2\x27\xcb\xec\xac\xb6\x1a\x37\x0a\x23\xed\xad\xda\x99\x0a\xdd\x12\xef\xf4\xec\x54\x4d\x9c\x2c\x81\xd2\xb4\x5b\xd2\xe0\x5d\xf2\x53\x02\x1c\xd5\x34\xa1\xab\x65\x2b\xcd\x6a\xb1\xab\xad\x7a\x9d\xa1\x0f\xf8\x46\xa7\x28\x95\xde\x7a\x55\x45\xed\xa4\xcc\x6a\xa1\xc2\xad\x09\x6c\xa2\x4e\x7b\x50\x83\x84\xad\xd7\xc6\x9c\xa0\xce\xd7\x95\xd4\x70\x54\x1c\xae\x57\xbe\x9c\x9f\xb0\xa8\xf3\x66\xbf\x9b\xde\x86\x1f\x7f\x60\xd2\x71\xe7\x9e\xe2\x16\x85\x4e\xa7\xff\xd6\x68\x15\xd4\xbc\xdf\xc3\xa3\xad\x3b\x6e\x8d\x07\xb9\xc6\xd6\x1a\x18\xa3\xed\xb6\x23\x0e\xf0\xbc\xf7\x56\x1f\xf6\x26\x0b\x6b\x53\xf0\x16\x6b\x76\xca\xf5\x73\x22\x68\x91\xa9\x38\xa8\x63\xb3\x2f\xf7\x7b\x3d\xb1\xa6\xcd\x2b\x3d\x63\x32\xec\x35\x01\x5b\x68\x82\xc6\xac\x31\xc6\xad\x19\x7c\x9b\x76\xc7\x5c\x36\x93\xb2\xec\xfc\xa0\x54\x1f\x16\xbc\x4e\xb3\xde\xce\xeb\x7d\x65\x0e\x86\xbd\x06\x37\xe3\x7b\x75\x69\xa9\xbe\xb6\xdb\xfe\xba\xf5\xa1\xf7\x7b\xef\x56\xaa\xd4\xde\x68\x30\x3f\xb0\x38\x6d\x24\x9b\x83\x8a\xcd\x69\x9b\x8a\xa5\x13\x54\x5d\x21\x61\x21\xc9\x5a\xa7\xa6\xe8\x8b\xa6\x9e\xed\x65\x80\x56\x5d\xda\xfe\x2c\xb6\x97\xdd\x35\x64\x02\xc0\xb4\x2f\xcf\x44\x7b\x03\x1a\x72\x7f\x6a\x6b\x66\xa7\xb8\xed\x68\xd5\xda\x46\x03\xef\x3a\x01\x8b\xd6\x3e\xfb\x4a\x40\x0c\xc1\x2b\x10\x25\x20\x66\x32\x99\x28\xcf\x81\xeb\x12\x63\x5f\x7d\xfc\xf6\xdb\x77\x46\x27\x9e\x0d\x83\x6f\x7f\x8f\x6a\x97\xbf\x7f\xff\xf1\xf5\x5f\x5f\xbf\xbe\xd8\xd0\x9b\x6b\x24\x74\xd2\x0a\xd1\x36\x71\x29\x94\xb6\xfd\x74\x80\xd6\x41\xda\xc7\x5b\x94\x86\xda\x6c\xe9\x07\x65\x86\x63\xd9\xbf\x45\xa5\x50\x3a\x44\xca\x1c\x07\x37\x20\x2c\xec\xa0\xb4\x89\xb0\x61\x46\x8b\x2f\xb9\x68\x4d\x25\x16\xf1\xca\xcc\x2f\xbc\xc0\x97\x78\xf4\xe3\xb2\xe0\x4a\x43\xd7\xb5\x50\xda\xdf\xf8\x01\xb2\x9f\x19\x31\xaa\xbb\x24\xa8\x0e\xe2\xdf\x75\xe2\x04\xcf\xcc\xd3\x00\x19\x04\x31\xa3\xd6\xd3\x33\xd3\x44\xd6\x0a\x05\x58\x85\xcf\x0c\xf0\x30\xb4\x9e\x19\x1f\x3a\x7e\xda\x47\x1e\xd6\x9f\x99\x27\x10\x11\x63\x2a\xd1\x96\x4c\xcd\x26\x33\xfc\x74\x82\x4e\x59\x19\x6c\x6c\x85\x58\x4f\x47\xae\x22\xa9\xca\x0c\x97\x77\xd7\x49\xe2\x84\xc4\xd3\xd2\xa1\x07\xdd\x32\xa3\x78\x08\xce\xd3\xd1\x02\x4d\x9f\x2f\xae\x95\x56\x63\xad\x1e\x34\x90\x87\x85\x6c\x21\x19\x96\x7b\xa6\x2d\xfb\xf1\xe7\xea\x8c\x10\xcb\xe6\x54\x35\x97\x44\x08\x51\xe9\x20\xe7\x9c\x17\x5d\xe0\x55\x2e\x89\x84\x6f\xe3\x1b\xcc\xf8\xdc\x19\xa9\x83\x61\x13\xb8\x71\x82\x73\x68\xbe\x98\xcf\xc2\x24\xe8\xf9\x19\xac\x56\xc8\x42\xa1\x94\xc8\x24\x95\x45\x57\xa3\xaf\xef\x58\x77\x7d\xb4\x17\x81\x0e\xe4\x25\x2c\xef\x2c\xa4\xd2\x55\xbb\x7f\xec\x7b\x77\x1e\xc3\x73\x23\x66\x79\x3d\xcf\x27\x09\xb7\xa2\x93\xb2\xc3\x33\x1a\x28\x9b\xe7\xd9\x44\x05\x29\xcb\x33\x60\x25\xcb\x69\x7c\x31\x09\x18\xe3\x33\x60\x1d\xea\x8a\xae\x46\x0e\xaf\x40\x75\x6e\x78\x64\xe9\x68\xe9\xc7\x28\xa9\xfc\xe3\x94\x6e\x5b\x58\xe5\xcb\x65\x05\xe9\xc4\x43\x7b\x8a\x4e\x80\x9c\xa0\xcc\x3c\xfd\xaf\xf4\x94\xe8\x15\x47\x5b\x7d\x36\x77\xe1\xfe\x08\x2b\xc4\xd2\x7e\x3c\xec\x8b\xb6\x75\x26\x47\x21\x9b\x53\x59\x3d\x11\xd8\xa4\x1a\xcc\x36\xf7\x87\x80\xfe\xd4\xbf\xcf\xe3\xed\xb3\x6e\xe3\x4f\x0a\x7e\xb0\xb1\x50\x99\xc1\x01\xb4\xb0\x4a\x8b\xbc\x09\x24\x94\x07\x14\x75\x9b\x82\x46\xf5\x33\xaa\xc1\x75\x1d\x21\x9d\x4d\x16\x87\xa3\x06\x0b\x3a\x29\x56\xd7\x75\x21\x91\x94\x4a\x3d\x29\x49\x4c\x29\x5a\xb2\x1f\xe0\x0b\x87\xce\xeb\x45\x3d\x89\xa9\x3b\xe6\xd2\xbc\xfb\xca\xbe\x1d\x9a\x95\xf3\x28\x92\x2b\xe6\xd9\x7c\x72\x88\x34\xce\x80\x4b\xb9\x92\x06\x13\x99\x53\x89\xb7\xf3\xf8\x38\xbf\x6b\x48\x25\x1e\x0c\x30\x71\xca\xcc\xd2\xd1\x90\x17\x25\xc3\x1f\x8f\xc5\xaf\xfd\x10\x25\x26\xa6\x61\xdf\xb5\xe0\xa6\xcc\x60\x27\x4e\xa7\x8a\x45\xd4\x78\xb0\xb2\x42\x5e\x94\xc7\xad\x34\xb4\xb0\xe1\x94\x77\xbb\x06\xc4\x8d\x73\x31\xb6\xac\x32\xa3\x2e\x3d\x0f\x39\x41\x9c\xcd\x69\xfb\xec\x34\x71\x6d\x84\xc0\x83\x8e\xef\xc2\x08\x97\x8a\x56\x86\x6a\x80\x57\xd7\xc1\x1a\x96\x4d\xb2\x42\x3b\x0b\x91\x65\x10\xb3\x1b\x62\x2d\x30\xcb\x0c\xd5\x63\xfd\xc0\x23\x8e\x71\x6d\x50\xec\x98\xc8\xc3\xd4\xbd\x93\x50\x22\x1f\x40\x54\x21\x4d\xee\xe4\x84\xc7\xd5\x09\x8f\xec\x48\x47\x36\xf4\x0c\xec\x94\x19\xf6\x25\x5f\x40\x36\x9d\x43\x6c\xef\xf6\x52\x88\xa7\x21\xef\x10\x21\x1c\xe2\x50\xcf\xb2\x4a\xb4\x6b\xa5\xcc\x15\xed\x6a\xcd\xf5\x28\x93\x33\x9b\x38\xc4\x77\xa1\x8a\x9e\x3f\xbf\x5e\xd6\x55\x11\xe7\x14\x11\xbd\x3d\x93\xeb\x08\x0e\x3b\x46\xf9\x10\xea\xd3\x0a\x89\x0b\xb1\x43\x0d\xc6\x46\x3f\x22\x1b\xe9\x16\x09\xcb\xcc\x0a\xfb\x58\xb1\xa8\xb2\x60\xc7\x5d\x06\x47\x26\x4f\x6c\x72\xa2\xb8\xdb\x78\x0f\x6d\xf3\x47\xb0\x71\xd1\x6f\x4f\xaa\x89\xd4\xb9\x42\xd6\x4f\x7f\x5e\x49\xb2\x57\xfd\x5e\x10\x17\x6a\x5a\xbc\x4c\xdd\xfc\xd7\x9b\xd8\xf7\x84\x3c\x5a\xe2\x44\xd6\x13\xe5\x9f\xac\x9e\xd5\xb5\x37\xbc\xf5\xb2\x60\xc9\xe7\xb5\xfc\x0f\x5a\x80\x48\xf2\xa8\xd3\x13\x75\x2b\xaa\x3c\x7e\x50\xf2\x2c\x55\x73\x7b\x17\x3a\xf3\x93\x83\x9d\xb9\x9c\xbb\xbe\xf4\x1c\x13\x6b\x1a\x72\xce\x63\xf8\x45\xe0\x38\x9c\x9a\x3d\xb5\xa3\x25\x82\x80\xd8\x65\x86\x73\xd7\x8c\x4f\x2c\xac\x31\xbf\x68\x3a\xe2\x11\x35\xaa\x9a\xde\x59\xfd\x72\x0c\x85\x01\x8c\xdd\xe9\xb4\xa2\xa1\xd6\x33\x11\x3e\xd4\x83\xbd\x02\x29\xe8\x16\x82\x5e\xe4\x24\x81\xf9\x00\xb5\x18\xed\x2c\x1c\xb8\x50\x3d\x38\xe3\x71\x51\x25\x96\x05\x5d\x1f\x45\x67\x70\xf7\x8d\x4a\xeb\x3a\x28\x04\x66\x4c\xfb\xb6\x87\x9b\xd7\x75\xb7\xc9\x5f\x2f\x65\xaf\x97\x84\xeb\xa5\xdc\xf5\x52\x3e\xe6\x61\x67\xf9\x28\x99\x9c\xb9\xc2\xd1\x78\x09\xac\x5d\x06\xda\x2c\xbf\x6b\x03\x1f\x72\x40\xfe\x2a\x4c\x0b\x3f\x81\x9d\xbd\xc2\x66\x7f\x02\x5b\xb8\xc4\x3e\x34\xb0\x0f\x61\xe7\xae\xb0\x7f\x86\xf3\xfc\x15\xf6\x4f\x68\xcd\x7d\xc8\x5c\x5c\xac\x8b\x6b\xec\xb8\x94\x58\x2c\x49\x80\x4e\xc8\x24\x58\x77\x69\x5d\x39\x0b\xb1\x4e\x1d\x36\x6d\x21\xfd\x3c\x72\xfc\xa4\xff\x10\x8b\x21\xd7\x9b\x2c\xad\xc3\x3e\x16\xf6\xf7\x35\x7a\x3a\x4a\x1a\x65\xc6\x22\x21\xf2\xd2\x1e\xb1\xa1\x43\xe7\x78\x87\x9c\x40\x92\xf2\x80\x24\x61\x90\x3d\xc6\x2d\x3e\xa0\xe5\x9a\xd4\x12\x58\xd3\x4e\x8d\x74\xd4\x12\xbd\x7c\xb8\x4e\x46\x4f\x83\xba\x44\x1c\x92\x7e\x47\xc6\xd2\x82\xde\xd3\x33\x53\x21\x8e\x4f\x2c\xe8\x3f\x33\x4f\x1d\xac\xa0\x5d\x3a\x60\x22\xa0\xa7\x67\x46\x42\x8e\x45\x22\x98\xa5\x87\x91\x77\x56\x4a\x50\xdc\xc7\x43\x57\x0e\xc4\x9d\x56\x47\x9f\x2e\xb4\x5b\xfd\xaf\x71\x76\x56\x02\x53\x2b\xdd\x1d\x67\x49\xb5\x38\x4b\x73\x4d\xe6\x7f\xb0\xed\x12\x2f\x80\xf4\xba\x36\xc2\xe3\x28\x78\x82\xbb\x7e\x00\x93\xa7\x60\x16\x1f\xc2\xcc\x52\x30\xa3\x00\xf4\x00\xaa\x40\x41\xe5\x1f\xe3\x37\x47\x41\x8d\x02\xf6\x03\xa8\x79\x9a\x92\xd8\x7b\xa8\xff\x56\x2a\xff\x0f\xe6\xf1\xdf\x7f\x2d\xeb\xd8\xf3\x83\xb4\x6a\x62\x4b\xbb\x0e\xa0\xb7\xa5\xf8\xfd\xd7\xb2\x05\x69\xc8\xc7\xd8\x76\x47\x81\xb0\xec\x90\xe0\xdb\x1f\xa6\x87\xf4\x3f\xbf\x9f\xd6\x87\x27\xf5\xe5\xc3\xf5\xe1\x0b\x74\x54\x93\xec\x7b\x62\x8b\xc0\xa0\xcc\x44\xe6\x38\xa9\x94\xd3\xde\x2e\x81\xec\x53\xd2\x59\x14\x4a\x1f\x92\xe4\xf9\x9c\xf6\xc6\x46\x65\x9d\xa8\x4b\xff\xb4\xc1\x4b\xe6\xcd\xbd\x0a\xa3\x9f\xf9\x86\x12\x93\x29\x41\xf7\x6a\x49\xbb\x5e\x8a\x7d\x20\xb1\xc1\x7a\x30\xb0\xdd\xaf\x87\x5f\xf8\xdc\x0e\xfd\xb4\x3c\xfb\x4c\x99\xf1\x99\x63\x13\x46\x1a\x88\x43\x02\x2a\x5e\x14\xc4\x77\xb3\xf1\xe7\x46\x07\xce\x2f\xc7\xe0\xc7\x92\x73\x1f\xde\x62\x1e\xef\x17\xd5\x9f\xfb\xfc\x7e\xfb\x20\xdc\x43\xbe\x73\x0e\x68\xe8\x73\x45\xbb\x3d\xd4\x38\x4a\x9d\x75\xd7\x4c\xee\xa4\x10\xda\xd7\x46\x1c\xcd\x61\xf7\x5e\x7c\xd0\x8f\x20\x08\x5a\x2e\x4f\x1b\x91\xd8\x58\xd3\x76\x11\x83\x36\xc4\xfa\x9c\xb3\xee\xcd\xb4\x53\x65\xd4\xa9\xfc\xa2\xe6\x55\x45\xe3\xae\xfa\x98\xf4\xe5\x64\xe8\xf8\xdc\x83\x1a\x5e\xfa\xb1\x1c\x3f\x0e\x3d\xa9\x09\xb5\xa8\x77\xc2\x8e\x8f\x02\x86\x65\xd2\x11\x65\xf6\xd6\x54\xe9\xbf\x5f\xed\xf3\xe7\x71\xe2\x78\x4a\x12\xeb\xe1\xab\x4b\x1e\x3e\x41\x92\x63\x4a\x8f\xdf\x0c\xa2\xc8\x75\x0f\x20\x7b\x0f\x40\xb8\x07\x90\xbb\x07\x90\x3f\x07\x38\xeb\xde\x39\x85\xd3\xf9\xec\x6d\xaf\x8a\x07\x1d\xd8\xc2\xc1\xe6\xb3\x43\xa6\xe9\x62\xdf\xd5\xef\x83\x2a\x45\x17\xf7\x00\xb2\xf7\x00\x84\x7b\x00\xb9\x7b\x00\xf9\x73\x00\xfa\x08\x22\x29\xf6\x5f\x8a\x78\xd7\xf8\x3f\x07\x9e\xfd\x39\x70\xe1\xe7\xc0\x73\x3f\x07\x7e\xa9\xa8\x6b\x17\x3a\x75\x8b\x1b\xb3\x30\xf3\xbc\xf8\x3c\x06\xd2\x97\xec\x49\xf1\x7d\x36\xc2\x4c\x9e\xa9\x20\x88\x54\x44\xbd\x6c\x31\xf9\x87\x77\xe1\x5e\x72\xff\xc6\x3e\x57\x2d\x39\xb7\x4f\xa2\x0f\x35\xe0\x09\x80\x57\xbd\x36\xfb\x52\x2c\x24\x51\xbd\x6a\xad\xd9\x97\x62\x8e\x9e\x4b\xff\x7a\xc3\xcb\xd3\x37\x8f\xbb\x4f\x7a\xeb\x49\x6f\x2f\x13\xba\xd1\xbf\x3c\x8e\xb1\x76\x37\x5b\x57\xd7\xed\xd0\xb2\xe8\xe0\xbf\x5f\x0f\x13\x76\x61\x9f\x06\x9c\xda\x93\x3f\x63\x2d\xd1\xbe\x9a\xf5\xc0\x68\x4b\xb3\x18\x2d\xb8\x82\xa3\x31\x74\xed\x29\x49\x37\x78\x8f\xcc\x50\xa2\x7d\xb5\xab\x92\x2b\x7f\x23\x0f\xde\x1a\x0d\x1e\x2b\x9b\x63\x49\xb3\xbf\x01\x39\xbc\xe2\xf1\x39\x43\x85\xcb\x80\x24\x13\xdb\x8f\x03\x1f\x91\xe0\x80\x40\x2f\x8e\x99\xe0\x42\xba\xa8\xb5\xe4\x8e\x75\xc9\xae\xd4\x79\x68\x1c\xbb\x27\xe7\x25\x5c\x1b\xfd\xa2\xeb\xfa\x49\xb4\xd8\x19\xec\x93\xee\xa1\x84\x4a\xa6\x5b\x76\x02\x73\x57\x50\x7e\xe3\x9d\xef\x89\xbb\xec\xaf\x0f\x93\x6f\x6b\x6c\xb8\x4e\x5f\x68\xfd\xc6\xdd\x48\x92\x20\x54\xfa\x7f\xc4\x39\xff\xb7\xb8\xa9\xfa\x93\x16\x0d\x12\x46\x6d\x9f\xa8\x11\xdc\x39\xe6\xbe\x43\x4b\x42\x3d\x8e\x85\x3e\xdd\xf3\x85\x47\x36\xc3\xbe\x08\x17\x17\x59\x17\x27\xa3\x98\xfb\x1b\x5d\x3e\xcf\x50\xe0\x37\xbe\xf0\x9c\xe5\x9e\xb3\xb9\x67\xf6\x85\xcd\x7d\x4f\xaa\x5c\x93\x67\x46\x27\x81\xc5\x21\x9e\x0d\xa9\x51\xc5\xf5\xd0\xef\x14\xfe\x1f\xe0\x3a\x26\x1f\x47\xac\x4f\xfa\x0c\x13\x9a\x38\x40\xf1\xdc\x1d\x95\x23\xe2\x8f\x5f\x40\x50\x5a\x5a\x13\x1b\xa6\x15\xe9\x9e\xd6\xbe\x24\x9d\xf4\x13\xac\x8b\x26\xf3\x62\x54\x46\x69\x3d\x69\x32\xdd\xdb\x21\xb1\xb5\x3d\x6a\xf3\x10\xaa\x2e\x03\xcb\x55\x65\x2e\x24\xf6\x27\x74\x2b\x1f\x0d\x77\xd1\xad\xfd\x38\x3f\x62\x87\xdd\x6e\x98\x97\x72\x15\x98\x74\x89\x46\x75\xac\xbb\x97\xd4\xf7\x4c\xad\x2f\x2d\x2b\xad\x12\xdb\xc6\x01\xf3\xa2\x04\x4e\xfa\x30\xbd\x70\x48\xf0\xad\xac\x61\x3f\x8a\x40\xda\xf7\x93\xdb\xb6\xcb\xf7\x11\x4e\xef\x77\xee\xbc\xa8\xf0\x97\x9b\xdc\x5d\x7f\xfb\x1f\x9c\xab\xfe\x3f\x34\xca\x1a\xa7\xe5\x34\x98\xdc\x28\x5f\x35\xd2\x3f\xd1\x28\x27\x67\x90\x72\x7c\x81\x8c\xb4\xd4\x4b\x44\x8e\xa4\x2d\xa8\xa0\x7d\xcd\x41\x7c\xbc\xeb\x55\x3c\x64\xc1\x00\xaf\x62\xa9\xb6\x69\xec\x68\x68\x1d\x0f\xbb\xae\x2d\xbb\xbf\xa7\xa5\xf8\x51\x00\xfd\x79\x3a\xbe\x00\xc0\x01\xb2\xe9\xd7\x01\x89\xb3\xba\x73\xe4\x14\x8d\xd8\x69\xd0\x48\x38\x99\x97\x68\x9f\x77\xd8\xc7\x23\x77\x4c\x10\xf1\x40\x28\xcd\xbd\xe4\x77\xf1\x27\xc9\xca\x37\xde\x28\xa0\x99\x10\x21\x2a\xce\x3f\xbf\x7e\xf9\x92\x74\xf5\xfe\xe5\x8b\x8d\x9d\x43\xc0\xe0\xd9\xd8\x0f\xbf\x7c\x39\x09\x22\xa5\xe2\x71\xed\x20\xc4\x2e\xae\x7c\xf9\x72\x3c\x2d\x42\x7c\x54\xbe\xfc\xeb\xeb\xd7\x7f\xd8\x48\xc3\x90\xf9\x76\x42\xa0\x90\x2f\xb8\xeb\xa8\x66\xf8\x42\x61\xeb\x33\x56\x1e\x48\x44\xec\x07\x44\x4d\x47\xb5\x00\xc4\xce\xfe\xd0\x9f\x27\x93\x5f\xf4\x52\xf4\x1f\x1b\x94\xf1\x90\x8b\x60\xc0\xf8\xaa\x47\x2c\x2b\x52\xf1\x8f\x1d\xfc\x75\xfd\x04\x21\xdc\x3d\xbb\x9e\x53\x9f\x25\xba\xd2\x2e\xd1\x25\xe5\x8f\xab\x1b\x81\xcb\xe5\x7d\xd9\x70\x5c\x3f\x0b\xc7\x07\xf9\x02\x1c\xec\x0b\xd3\x8b\x52\xb2\xc0\xee\x25\x88\x7b\xfa\xbd\x43\xa8\xc8\x09\x0e\x6f\xd1\x9c\x6b\xc7\xc2\xcf\xcc\xc5\x12\xb1\x28\x4b\x8c\x85\x77\x9b\x7d\x1e\x8b\x32\x43\x96\x81\x8f\x35\xb4\x53\x64\xfc\x71\x3e\x22\xff\xbf\x00\x00\x00\xff\xff\x00\xf6\x39\x84\x50\x36\x00\x00"),
		},
		"/_themes/github/templates": &vfsgen۰DirInfo{
			name:    "templates",
			modTime: time.Date(2019, 2, 6, 5, 9, 37, 291710100, time.UTC),
		},
		"/_themes/github/templates/main.html": &vfsgen۰CompressedFileInfo{
			name:             "main.html",
			modTime:          time.Date(2019, 2, 7, 10, 50, 19, 643969700, time.UTC),
			uncompressedSize: 377,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x90\xb1\x4e\xc4\x30\x0c\x86\xf7\x3c\x85\xe9\xc0\x76\xcd\x8a\x84\x2f\x4b\x8f\x13\x03\x12\x88\x96\x81\x31\x34\x3e\x1a\x91\x36\x55\x62\x1d\xaa\xaa\xbe\x3b\x4a\xd3\x72\x53\x62\xff\x9f\xad\x2f\xc1\xbb\xd3\x6b\xd5\x7c\xbe\x3d\x41\xc7\xbd\x53\x02\xf7\x83\xb4\x51\x02\x7b\x62\x0d\x6d\xa7\x43\x24\x3e\x16\x1f\xcd\xf9\xf0\x50\x28\x81\x6c\xd9\x91\x9a\xe7\xb2\x49\x97\x65\x41\x99\x3b\x62\x9e\xa1\xac\x79\x72\x54\x77\x44\x1c\x61\x59\x72\xab\x0d\x76\xfc\x2f\xed\x05\xca\x13\x5d\x53\x05\x90\xe2\x17\x7b\xa5\x77\x72\x5e\x9b\x0c\x6e\x1c\x0d\x66\x5f\xf0\x4c\xda\x50\x48\x15\xca\x4d\xed\xcb\x9b\x49\x09\xd4\x81\x6d\xeb\x08\x5a\xa7\x63\x3c\x16\xbd\x0e\x3f\xc6\xff\x0e\x87\x14\x17\x59\xa8\xf2\x03\xd3\xc0\x79\x7a\xe3\x73\x72\xf6\x9e\x29\xdc\xac\x6a\xcb\xb4\x62\x97\x35\x50\x02\x00\x47\x85\xb1\xd7\xce\xa9\xfb\xd6\x8f\xd3\xe3\x2a\x9c\xb8\xb2\xf2\xe3\x14\xec\x77\x97\x16\x03\xca\x0c\xa1\x1c\x95\x40\xb9\xcf\xdf\x5e\x81\x72\x13\x96\xf9\x87\xff\x02\x00\x00\xff\xff\x6e\xdb\xbe\xc5\x79\x01\x00\x00"),
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