.PHONY: build
build: create-bin-dir
	$(MAKE) build-protoc-gen-go-extractors

build-protoc-gen-go-extractors:
	go build ./cmd/protoc-go-extractors
	mv protoc-go-extractors bin/

.PHONY: build-protoc-gen-go-extractors-debug
build-protoc-gen-go-extractors-debug:
	go build -gcflags="all=-N -l" ./cmd/protoc-go-extractors
	mv protoc-go-extractors bin/

create-bin-dir:
	mkdir -p bin