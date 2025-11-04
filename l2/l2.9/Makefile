.PHONY: check lint vet test run

test:
	go test -v ./...

run:
	go run ./cmd/main.go -file examples.txt

check: vet lint

vet:
	@echo "→ Running go vet..."
	@go vet ./... && echo "✔ vet passed"

lint:
	@echo "→ Running golint..."
	@if ! command -v golint >/dev/null 2>&1; then \
		echo "golint не установлен. Установи с помощью: go install golang.org/x/lint/golint@latest"; \
		exit 1; \
	fi
	@golint ./... && echo "✔ lint passed"