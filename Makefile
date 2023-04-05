BINARY_NAME=build/start-server
COVERAGE_PROFILE=build/coverage.out
COVERAGE_REPORT=build/coverage.html
CONTAINER_IMAGE=rbender-hbo/hello-go-rest:latest

.PHONY: build test

clean:
	rm -rf build

test:
	go test ./...

integration-test-local:
	go test -v -tags=integration ./itest/...
.PHONY: integration-test-local

start-wiremock:
	docker run -it --rm \
	-p 8081:8080 \
	--name wiremock \
	wiremock/wiremock:2.35.0
.PHONY: start-wiremock

start-integration:
	PRODUCT_BASE_URL="http://localhost:8081" ${BINARY_NAME}
.PHONY: start-integration

integration-test: docker-build
	docker-compose up -d
	sleep 5
	# Store the exit code so we can cleanup first
	go test -tags=integration ./itest/... ; echo "$$?" > test.result
	docker-compose down
	exit $$(cat test.result)
.PHONY: integration-test

coverage:
	mkdir -p build
	go test -coverprofile=${COVERAGE_PROFILE} ./...
	go tool cover -html=${COVERAGE_PROFILE} -o ${COVERAGE_REPORT}

build:
	go build -o ${BINARY_NAME} cmd/start-server.go

run:
	${BINARY_NAME}

docker-build:
	DOCKER_BUILDKIT=1 docker build --tag $(CONTAINER_IMAGE) .
.PHONY: docker-build

docker-run: docker-build
	CONTAINER_IMAGE=$(CONTAINER_IMAGE) docker-compose up
.PHONY: docker-run

all: clean coverage build run
