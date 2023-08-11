CUR_DIR=`pwd`

shell:
	docker run -it --mount type=bind,source=${CUR_DIR},destination=/usr/src/app --rm animals:latest /bin/bash

all: build run

build:
	docker build -t animals:latest .

dev:
	docker run -it --mount type=bind,source=${CUR_DIR},destination=/usr/src/app --rm animals:latest go run main.go

run:
	docker run -it --rm --name my-running-app animals
