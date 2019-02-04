SRCS    := $(shell find . -type f -name '*.go' | grep -v asset)
ASSETS  := $(shell find ./asset/_root)

all: $(SRCS) asset/assets_vfsdata.go
	GO111MODULE=on go build ./cmd/mellowdown

asset/assets_vfsdata.go: asset/assets.go $(ASSETS)
	GO111MODULE=on go generate ./asset
