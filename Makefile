GOLINT := golangci-lint

all: dep fmt test

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
	@which $(GOLINT) || curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GO_PATH)/bin v1.46.2
