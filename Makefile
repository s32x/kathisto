init:
	-rm -rf ./vendor go.mod go.sum
	GO111MODULE=on go mod init

deps:
	-rm -rf ./vendor go.sum
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor
	
test:
	go test ./...
	
run:
	go run main.go