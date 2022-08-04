build:
	cp -r ../violas_client ./
	docker build -t metatoc-webservice .
build-test:
	cp -r ../violas_client ./
	docker build -t metatoc-webservice-test .
run:
	docker run -d --name=metatoc-webservice -p 2929:5000 metatoc-webservice
run-test:
	docker run -d --name=metatoc-webservice-test -p 2929:5000 metatoc-webservice-test
clean:
	docker stop metatoc-webservice
	docker rm metatoc-webservice
	docker rmi metatoc-webservice
clean-test:
	docker stop metatoc-webservice-test
	docker rm metatoc-webservice-test
	docker rmi metatoc-webservice-test
