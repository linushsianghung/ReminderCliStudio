.PHONY: client

client:
	@echo "Build the client binary"
	go build -o bin/client cmd/client/main.go
got
