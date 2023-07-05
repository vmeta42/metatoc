IMAGE_NAME = meta42/metatoc-violas-config-generator
VERSION = latest

build:
	docker build -t $(IMAGE_NAME) .
run:
	docker run --rm -i -t -v ./etc:/opt/diem/out_etc $(IMAGE_NAME) bash
gen:
	docker run --rm -i -t -v ./etc:/opt/diem/out_etc $(IMAGE_NAME)
rmi:
	docker rmi $(IMAGE_NAME)
tag:
	docker tag violas:latest $(IMAGE_NAME):$(VERSION)
