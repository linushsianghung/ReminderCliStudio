GO ?= go
.PHONY: client

# Via go help build
# If the arguments to build are a list of .go files from a single directory, build treats them as a list of source files specifying a single package.
client:
	@echo "Build the client binary"
	${GO} build -o bin/client cmd/client/*.go

# Build with Optimisation Decision
clientOpt:
	@echo "Print optimisation decisions"
	${GO} build -gcflags "-m=1" cmd/client/*.go
