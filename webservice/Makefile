# Copyright 2022-2023 The VNET Project Authors. All Rights Reserved.

# SPDX-License-Identifier: MIT

build:
	cp -r ../violas_client ./
	docker build -t meta42/metatoc-webservice .
run:
	docker run -d --name=metatoc-webservice -p 2929:5000 meta42/metatoc-webservice
clean:
	docker stop metatoc-webservice
	docker rm metatoc-webservice
	docker rmi meta42/metatoc-webservice
