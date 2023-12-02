.PHONY: install run test new

install:
	go install

run:
	go run . $(day)

test:
	go test  -cover ./...

.PHONY: help
help:
	@echo "Available targets:"
	@echo "- install: Install dependencies"
	@echo "- run day=<day_number>: Run the code for a specific day (provide day_number)"
	@echo "- test: Run tests"
