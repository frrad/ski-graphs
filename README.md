## Installation

First install the binary with

```
cd /tmp && GO111MODULE=on go get github.com/frrad/ski-graphs
```

Then install InfluxDB 2.0+

## Setup
Setup some crons to fetch from APIs.

To get data from the ikon API do like
```
curl --silent 'https://mtnpowder.com/feed?resortId=61' > "/path/to/data/ikon/resort-61.$(date --rfc-3339=second).json"
curl --silent 'https://mtnpowder.com/feed?resortId=62' > "/path/to/data/ikon/resort-62.$(date --rfc-3339=second).json"
```

Vail resorts are like
```
curl --silent "https://cms.mountain.live/lumiplan/api/moduleplr/2542/detail?lang=en&platform=ANDROID" > "/path/to/data/vail/resort-2542.$(date --rfc-3339=second).json"
```

Also run the binary on a cron to process the results. Make sure to run `ski-graphs -help` to see what all flags must be set.

If you need to re-process things you should be able to clear influx with something like 

```
influx delete --bucket skiing --start 2009-01-02T23:00:00Z --stop 2099-01-02T23:00:00Z --org org
```

then just remove your seen-file and re-run things.

## Development

Test examples are formatted and have keys sorted like this:
```
jq --sort-keys . < resort-61.2021-01-01\ 18:32:01-08:00.json > testinput/test_one.json
```
