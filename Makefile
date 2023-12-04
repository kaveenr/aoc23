.PHONY: help
help:
	@echo "Available targets:"
	@echo "- install: Install dependencies"
	@echo "- run day=<day_number>: Run the code for a specific day"
	@echo "- test: Run tests"
	@echo "- bench: Run benchmarks"
	@echo "- new day=<day_number>: Create from template"
	@echo "- scrape year=<year> day=<day_number>: Scrape puzzle and input from site"
	@echo "- answer year=<year> day=<day_number> part=<part> answer=<answer>: Answer puzzle"
	@echo ""
	@echo "Hints: by default <year>,<day> will be set to current EST date"

.PHONY: install run test

install:
	go install

run:
	go run . $(day)

test:
	@go test  -cover $(shell go list ./...)

bench:
	@go test -benchmem -run=^$ -cover -bench . $(shell go list ./...)

.PHONY: new scrape
-include .env
year = $(shell env TZ=EST date +"%Y")
day = $(shell env TZ=EST date +"%-d")

new:
	@mkdir day$(day)
	@cp dayn/day.tmpl day$(day)/day$(day).go
	@cp dayn/day_test.tmpl day$(day)/day$(day)_test.go
	@sed -i '' 's/ayn/ay$(day)/g' day$(day)/day$(day).go
	@sed -i '' 's/ayn/ay$(day)/g' day$(day)/day$(day)_test.go
	@echo "Created from template ./day$(day)/"

scrape:
	mkdir -p puzzles
	@curl -s https://adventofcode.com/$(year)/day/$(day)/input \
		-b "session=${AOC_SESSION}" \
		> inputs/day$(day).txt
	@curl -s https://adventofcode.com/$(year)/day/$(day) \
		-b "session=${AOC_SESSION}" \
		| sed -n '/<main>/,/<\/main>/p' \
		| pandoc  --from=html --to=plain \
		> puzzles/day$(day).txt
	@echo "Scraped Day $(day) for $(year)"

answer:
	@curl -s -X 'POST' 'https://adventofcode.com/$(year)/day/$(day)/answer' \
		-b "session=${AOC_SESSION}" \
		-H 'Content-Type: application/x-www-form-urlencoded' \
		--data 'level=$(part)&answer=$(answer)' \
		| sed -n '/<main>/,/<\/main>/p' \
		| pandoc  --from=html --to=plain