all: dep fmt test

dep:
	go mod tidy
	go mod download

fmt:
	go fmt github.com/OrbisSystems/orbis-sdk-go/...

test:
	go test -tags=unit -cover -race -count=1 -timeout=60s ./...
