BUILDS_PATH = builds

build:
	go build -o ${BUILDS_PATH}/aos ./cmd/aos.go 

run:
	go run ./cmd/aos.go 

clean:
	rm -fr ${BUILDS_PATH}
