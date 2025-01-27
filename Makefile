all: build
.PHONY: all

build:
	go build github.com/openshift/client-go/...
.PHONY: build

build-examples:
	go build -o examples/build/app github.com/openshift/client-go/examples/build/...
.PHONY: build-examples

verify:
	hack/verify-codegen.sh
.PHONY: verify

update: generate
.PHONY: update

generate:
	hack/update-codegen.sh
.PHONY: generate
