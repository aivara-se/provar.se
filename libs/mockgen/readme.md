# @provar/mockgen

A JavaScript library to create mock feedback data that can be imported to Provar.se.

## Usage

> `Usage: mockgen --type <type> --count <count> --period <period>`

All command line parameters are optional and will fallback to default values.

```bash
# basic usage
npx @provar/mockgen > feedbacks.csv

# with a feedback type and count
npx @provar/mockgen -c 100 -t csat > feedbacks.csv
```
