export GO111MODULE=on
BINARY_NAME=tech-challenge-time

all: deps build
install:
	go install
build:
	go build
test:
	go vet ./...
	go test ./...
clean:
	go clean
	rm -f $(BINARY_NAME)
deps:
	go get -v ./...
upgrade:
	go get -u
