.PHONY: help
help:
	@echo "Available targets:"
	@echo "- install: Install dependencies"
	@echo "- run day=<day_number>: Run the code for a specific day"
	@echo "- test: Run tests"
	@echo "- bench: Run benchmarks"
	@echo "- new day=<day_number>: Create from template"
	@echo "- scrape year=<year> day=<day_number>: Scrape challange and input from site"
	@echo ""
	@echo "Hints: by default <year>,<day> will be set to current EST date"

.PHONY: install run test

install:
	go install

run:
	go run . $(day)

test:
	@go test  -cover $(shell go list ./... | grep -e "/day\\d")

bench:
	@go test -benchmem -run=^$ -cover -bench . $(shell go list ./... | grep -e "/day\\d")

.PHONY: new scrape
-include .env
year = $(shell env TZ=EST date +"%Y")
day = $(shell env TZ=EST date +"%-d")

new:
	@mkdir day$(day)
	@cp dayn/dayn.go day$(day)/day$(day).go
	@cp dayn/dayn_test.go day$(day)/day$(day)_test.go
	@sed -i '' 's/ayn/ay$(day)/g' day$(day)/day$(day).go
	@sed -i '' 's/ayn/ay$(day)/g' day$(day)/day$(day)_test.go
	@echo "Created from template ./day$(day)/"

scrape:
	mkdir -p puzzles
	@curl https://adventofcode.com/$(year)/day/$(day)/input \
		-b "session=${AOC_SESSION}" \
		> inputs/day$(day).txt
	@curl https://adventofcode.com/$(year)/day/$(day) \
		-b "session=${AOC_SESSION}" \
		| sed -n '/<main>/,/<\/main>/p' \
		| pandoc  --from=html --to=plain \
		> puzzles/day$(day).txt
	@echo "Scraped Day $(day) for $(year)"