build:
	docker build -t violas .
run:
	docker run --rm -i -t -v ./tmp_etc:/opt/diem/out_etc violas bash
gen:
	docker run --rm -i -t -v ./tmp_etc:/opt/diem/out_etc violas
rmi:
	docker rmi violas
