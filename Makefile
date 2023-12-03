.PHONY: install run test new scrape
include .env

install:
	go install

run:
	go run . $(day)

test:
	go test  -cover ./...

new:
	@cp dayn/dayn.go day$(day)/day$(day).go
	@cp dayn/dayn_test.go day$(day)/day$(day)_test.go
	@sed -i '' 's/ayn/ay$(day)/g' day$(day)/day$(day).go
	@sed -i '' 's/ayn/ay$(day)/g' day$(day)/day$(day)_test.go
	@echo "Created from template ./day$(day)/"

scrape:
	mkdir -p challanges
	@curl https://adventofcode.com/$(year)/day/$(day)/input \
		-b "session=${AOC_SESSION}" \
		> inputs/day$(day).txt
	@curl https://adventofcode.com/$(year)/day/$(day) \
		-b "session=${AOC_SESSION}" \
		| sed -n '/<main>/,/<\/main>/p' \
		| pandoc  --from=html --to=plain \
		> challanges/day$(day).txt
	@echo "Scraped Day $(day) for $(year)"

.PHONY: help
help:
	@echo "Available targets:"
	@echo "- install: Install dependencies"
	@echo "- run day=<day_number>: Run the code for a specific day"
	@echo "- test: Run tests"
	@echo "- new day=<day_number>: Create from template"
	@echo "- scrape year=<year> day=<day_number>: Scrape challange and input from site"
