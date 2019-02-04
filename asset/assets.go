// +build dev

package asset

import "net/http"

var Assets http.FileSystem = http.Dir("./_root")
