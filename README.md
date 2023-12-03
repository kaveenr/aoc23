# AoC 2023

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
- scrape year=<year> day=<day_number>: Scrape challange and input from site

Hints: by default <year>,<day> will be set to current EST date
```