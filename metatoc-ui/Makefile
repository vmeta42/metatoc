.PHONY: build
build:
	docker build --tag metatoc-ui:latest .

.PHONY: run
run:
	docker run --name metatoc-ui --rm -d -p 80:80 -e chatgpt_service_proxy_pass=http://metatoc-ui-chatgpt-proxy:3001 -e blockchain_service_proxy_pass=http://172.22.50.202:2929 metatoc-ui:latest

.PHONY: stop
stop:
	docker stop metatoc-ui

.PHONY: pushDev
push:
	docker tag metatoc-ui:latest harbor.dev.21vianet.com/metatoc/metatoc-ui:latest
	docker push harbor.dev.21vianet.com/metatoc/metatoc-ui:latest
