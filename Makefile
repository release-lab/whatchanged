test:
	go test -v --cover -covermode=count -coverprofile=coverage.out ./...

build:
	goreleaser release --snapshot --rm-dist --skip-publish

lint:
	golangci-lint run ./... -v

format:
	go fmt ./...

format-check:
	gofmt -l -d ./..

generate-test:
	go run cmd/whatchanged/main.go --project=./__test__/v HEAD~ > ./__test__/v.CHANGELOG.md
	go run cmd/whatchanged/main.go --project=./__test__/vscode-deno HEAD~ > ./__test__/vscode-deno.CHANGELOG.md
	go run cmd/whatchanged/main.go 23448a5482359f28a0089b17280dd2a0a0eaef26~9dff4fc6a9d746ffd9dd10215cf04d2fec2edd2a > ./__test__/whatchanged-[23448a5482359f28a0089b17280dd2a0a0eaef26~9dff4fc6a9d746ffd9dd10215cf04d2fec2edd2a].CHANGELOG.md
	go run cmd/whatchanged/main.go v0.2.0 > ./__test__/whatchanged-[v0.2.0].CHANGELOG.md
	go run cmd/whatchanged/main.go v0.2.0~v0.1.0 > ./__test__/whatchanged-[v0.2.0~v0.1.0].CHANGELOG.md

update-go-deps:
	@echo ">> updating Go dependencies"
	@for m in $$(go list -mod=readonly -m -f '{{ if and (not .Indirect) (not .Main)}}{{.Path}}{{end}}' all); do \
		go get $$m; \
	done
	go mod tidy -go=1.18
ifneq (,$(wildcard vendor))
	go mod vendor
endif