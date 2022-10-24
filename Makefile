BINARY_NAME=build/start-server
COVERAGE_PROFILE=build/coverage.out
COVERAGE_REPORT=build/coverage.html

.PHONY: build test

clean:
	rm -rf build

test:
	go test ./...

coverage:
	mkdir -p build
	go test -coverprofile=${COVERAGE_PROFILE} ./...
	go tool cover -html=${COVERAGE_PROFILE} -o ${COVERAGE_REPORT}

build:
	go build -o ${BINARY_NAME} cmd/start-server.go

run:
	${BINARY_NAME}

all: clean coverage build run
