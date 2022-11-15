build:
	go build -o bin/crypto-fetcher

run: build
	./bin/crypto-fetcher