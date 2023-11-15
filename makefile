BINARY_NAME=ipinfo

build:
	@go build -o ${BINARY_NAME} .
run:
	@go build -o ${BINARY_NAME} .
	@./${BINARY_NAME}

test:
	go test -v .

clean:
	clear
	@go clean
	@rm ${BINARY_NAME}
