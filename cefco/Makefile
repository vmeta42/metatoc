idxGenOut := $(shell grep -l -R "DO NOT EDIT" pkg/apis/)
idxGenIn:= $(shell grep -l -R -F "// +k8s:" pkg/apis)
idxSrc := $(shell find cmd pkg controllers -name "*.go")

now := $(shell date +%Y-%m-%dT%H:%M:%S%z)
VERSION ?= version-not-set

ifneq ($(gitInfo),)
	linkerVars := -X main.BuildTime=$(now) -X main.GitInfo=$(gitInfo) -X main.Version=$(VERSION)
else
	gitBranch := $(shell git rev-parse --abbrev-ref HEAD)
	gitCommit := $(shell git rev-parse --short HEAD)
	repoDirty := $(shell git diff --quiet || echo "-dirty")
	linkerVars := -X main.BuildTime=$(now) -X main.GitInfo=$(gitBranch)-$(gitCommit)$(repoDirty) -X main.Version=$(VERSION)
endif


vendor: go.mod go.sum
	go mod vendor
	touch $@

$(idxGenOut): vendor $(idxGenIn) hack/custom-boilerplate.go.txt
	bash vendor/k8s.io/code-generator/generate-groups.sh all \
		../pkg/generated \
		../pkg/apis \
		"filesync:v1alpha1" \
		--go-header-file hack/custom-boilerplate.go.txt


.PHONY: run filesync-controller filesync-controller-docker clean-vendor
run: $(idxSrc) vendor
	go run cmd/filesync/main.go --kubeconfig /root/.kube/config

filesync-controller: $(idxSrc) vendor
ifneq ($(noRace),)
	go build -o $@ \
		-ldflags "$(linkerVars)" \
		./cmd/filesync
else
	go build -race -o $@ \
		-ldflags "$(linkerVars)" \
		./cmd/filesync
endif

filesync-controller-docker:
ifneq ($(filesyncVersion),)
	docker build --tag idx/filesync-controller:$(filesyncVersion) \
		--build-arg VERSION=$(filesyncVersion) \
		--build-arg GIT_INFO=$(gitBranch)-$(gitCommit)$(repoDirty) \
		--file docker/filesync/Dockerfile .
else
	# Missing filesyncVersion, try again.
	# make filesync-controller-docker filesyncVersion=0.2.0
	exit 1
endif

clean-vendor:
	rm -rf ./vendor
