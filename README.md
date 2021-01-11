to get data from the ikon API do like
```
curl --silent 'https://mtnpowder.com/feed?resortId=61' > "/path/to/data/ikon/resort-61.$(date --rfc-3339=second).json"
curl --silent 'https://mtnpowder.com/feed?resortId=62' > "/path/to/data/ikon/resort-62.$(date --rfc-3339=second).json"
```

vail resorts like
```
curl --silent "https://cms.mountain.live/lumiplan/api/moduleplr/2542/detail?lang=en&platform=ANDROID" > "/path/to/data/vail/resort-2542.$(date --rfc-3339=second).json"
```


tests are craeted like this:
```
jq --sort-keys . < resort-61.2021-01-01\ 18:32:01-08:00.json > testinput/test_one.json
```

