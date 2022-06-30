build:
	docker build -t metatoc-webservice .
run:
	docker run -d --name=metatoc-webservice -p 2929:5000 metatoc-webservice
clean:
	docker stop metatoc-webservice
	docker rm metatoc-webservice
	docker rmi metatoc-webservice
