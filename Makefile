GOLINT := golangci-lint

all: dep dep-update fmt lint test

dep:
	go mod tidy
	go mod download

dep-update:
	go get -t -u ./...

fmt:
	go fmt github.com/OrbisSystems/orbis-sdk-go/...

lint: dep check-lint
	$(GOLINT) run --timeout=5m -c .golangci.yml

test:
	go test -tags=unit -cover -race -count=1 -timeout=60s ./...

check-lint:
	@which $(GOLINT) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.51.1

dc-up:
	docker-compose up -d

dc-down:
	docker-compose down