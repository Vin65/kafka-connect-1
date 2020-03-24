build:
	go build -o kafka-connect

test:
	go test ./...

fmt:
	go fmt ./...