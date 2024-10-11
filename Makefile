build:
	@go build -o bin/mtgCollectionWebAPI

run: build
	@./bin/mtgCollectionWebAPI

test:
	@go test -v ./...