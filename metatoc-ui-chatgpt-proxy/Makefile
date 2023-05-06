.PHONY: build
build:
	docker build --tag metatoc-ui-chatgpt-proxy:latest .

.PHONY: run
run:
	docker run --name metatoc-ui-chatgpt-proxy --rm -d -e API_KEY=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 -e PROXY_URL=http://172.16.31.118:15732  metatoc-ui-chatgpt-proxy:latest

.PHONY: stop
stop:
	docker stop metatoc-ui-chatgpt-proxy

.PHONY: push
push:
	docker tag metatoc-ui-chatgpt-proxy:latest harbor.dev.21vianet.com/metatoc/metatoc-ui-chatgpt-proxy:latest
	docker push harbor.dev.21vianet.com/metatoc/metatoc-ui-chatgpt-proxy:latest
