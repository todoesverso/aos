BUILDS_PATH = builds

build:
	go build -o ${BUILDS_PATH}/aos ./cmd/aos.go 

test:
	go test -cover ./...

fmt:
	go fmt ./...

clean:
	rm -fr ${BUILDS_PATH}
