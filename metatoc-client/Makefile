goSrc := $(shell find . -name "*.go")

metatoc-client.docker: $(goSrc)
	CGO_ENABLED=0 go build -o $@ -tags timetzdata

.PHONY: docker
docker: Dockerfile
ifneq ($(dtag),)
	docker build --tag metatoc-client:$(dtag) .
else
	# Missing dtag, try again. Example: make docker dtag=1.2.3
	exit 1
endif
