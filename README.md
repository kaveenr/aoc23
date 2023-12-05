# AoC 2023

[![Go Tests](https://github.com/kaveenr/aoc23/actions/workflows/tests.yml/badge.svg)](https://github.com/kaveenr/aoc23/actions/workflows/tests.yml)

## Prerequsites

1. GNU + Darwin/Linux enviroment
2. Go installed
3. Curl / Pandoc installed for scraping (optional)
4. Make `.env` file from `.env.sample` (optional)

## Instuctions

```sh
$ make
Available targets:
- install: Install dependencies
- run day=<day_number>: Run the code for a specific day
- test: Run tests
- bench: Run benchmarks
- new day=<day_number>: Create from template
- scrape year=<year> day=<day_number>: Scrape puzzle and input from site
- scrape-all year=<year> day=<day_number>: Scrape from day 1 to specified
- answer year=<year> day=<day_number> part=<part> answer=<answer>: Answer puzzle

Hints: by default <year>,<day> will be set to current EST date
```
