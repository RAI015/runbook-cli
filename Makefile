APP := runbook

.PHONY: test build demo

test:
	go test ./...

build:
	go build -o $(APP) ./cmd/runbook

demo:
	go run ./cmd/runbook generate -i examples/runbook.yaml -o runbook.md
