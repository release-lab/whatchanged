test:
	go test -v --cover -covermode=count -coverprofile=coverage.out ./...

lint:
	golangci-lint run ./... -v

format:
	go fmt ./...

format-check:
	gofmt -l -d ./..
