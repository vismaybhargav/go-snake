BINARY_NAME=go-snake

build:
	go build -o ${BINARY_NAME}.exe src/*.go

run:
	go build -o ${BINARY_NAME}.exe src/*.go
clean:
	go clean
	rm ${BINARY_NAME}
dep:
	go mod download
