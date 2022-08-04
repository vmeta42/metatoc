goSrc := $(shell find . -name "*.go")

nats-everywhere.docker: $(goSrc)
	CGO_ENABLED=0 go build -o $@ -tags timetzdata

.PHONY: docker
docker: Dockerfile
ifneq ($(dtag),)
	docker build --tag nats-everywhere:$(dtag) --build-arg NATS_URL=$(nurl) --build-arg NATS_SUBJECT=$(nsubject) --build-arg NATS_DURABLE=$(ndurable) .
else
	# Missing dtag, try again. Example: make docker dtag=1.2.3
	exit 1
endif

.PHONY: run
run:
	docker run -it -d -v /root/nats-everywhere/logs:/opt/logs nats-everywhere:$(dtag) /bin/bash
