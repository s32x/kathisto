NOVENDOR_PATH = $$(glide novendor)
.PHONY: test

glide:
	-rm glide.lock
	-rm -r vendor
	glide cache-clear
	glide install

test:
	go clean
	go test ${NOVENDOR_PATH}

build:
	make glide
	make test
	-docker rmi ubuntu
	-docker rmi kathisto
	-rm -r kathisto
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o kathisto
	docker build --no-cache -t entrik/kathisto .
	rm -r kathisto