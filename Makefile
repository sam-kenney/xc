NAME := xc

build:
	go build -ldflags="-X 'main.apiKey=$(EXCHANGE_API_KEY)'" -o bin/$(NAME) cmd/main.go

install:
	go install -ldflags="-X 'main.apiKey=$(EXCHANGE_API_KEY)'"
