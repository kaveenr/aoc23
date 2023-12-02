.PHONY: install run test new

install:
	go install

run:
	go run . $(day)

test:
	go test  -cover ./...

new:
	mkdir day$(day)
	cp dayn/dayn.go day$(day)/day$(day).go
	cp dayn/dayn_test.go day$(day)/day$(day)_test.go
	sed -i '' 's/ayn/ay$(day)/g' day$(day)/day$(day).go
	sed -i '' 's/ayn/ay$(day)/g' day$(day)/day$(day)_test.go

.PHONY: help
help:
	@echo "Available targets:"
	@echo "- install: Install dependencies"
	@echo "- run day=<day_number>: Run the code for a specific day (provide day_number)"
	@echo "- test: Run tests"
	@echo "- newday=<day_number>: Create from template"
