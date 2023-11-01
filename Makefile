GOLINT := golangci-lint

PACKAGES_FOR_TEST := $(shell go list ./... | grep -v config | grep -v bin | grep -v interfaces | grep -v model | grep -v utils)

all: gen-mock dep dep-update fmt lint test

dep:
	go mod tidy
	go mod download

dep-update:
	go get -t -u ./...

fmt:
	go fmt github.com/OrbisSystems/orbis-sdk-go/...

lint: dep check-lint
	$(GOLINT) run --timeout=5m -c .golangci.yml

test: gen-mock
	@go test -tags=unit -cover -race -count=1 -timeout=60s $(PACKAGES_FOR_TEST)

check-lint:
	@which $(GOLINT) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.55.1

dc-up:
	docker-compose up -d

dc-down:
	docker-compose down

gen-mock: check-mockgen
	mockgen -package mock -source interfaces/interface.go -destination interfaces/mock/interface.go

check-mockgen:
	@which mockgen || go install github.com/golang/mock/mockgen@v1.6.0
