# AoC 2023

## Prerequsites

1. GNU + Darwin/Linux enviroment
2. Go installed
3. Curl / Pandoc installed for scraping (optional)
4. Make ENV file from `.env.sample`

## Instuctions

```sh
$ make help
Available targets:
- install: Install dependencies
- run day=<day_number>: Run the code for a specific day
- test: Run tests
- new day=<day_number>: Create from template
- scrape year=<year> day=<day_number>: Scrape challange and input from site
```