default: go-dep go-lint test

go-dep:
	go get -v ./...

go-fmt:
	@diff -u <(echo -n) <(gofmt -d -s .)

go-lint: go-fmt
	@go get github.com/golang/lint/golint; \
	golint $(GO_EXTRAFLAGS) -set_exit_status ./...; \
	go vet -v ./...

test:
	go test $(GO_EXTRAFLAGS) -v -race ./...

go-build: go-dep go-lint test
	go clean -v
	go build -v $(GOFLAGS)
