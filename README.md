A forex currency CLI tool with caching.

When building, specify cacheExpiryDurationInHours.

```bash

go build -ldflags "-X 'main.cacheExpiryDurationInHours=12' -X 'main.version=1.0.0'" -o forex

```

After building move binary to `/usr/local/bin/`, this will allow you to call it from any directory

Todo

- Write unit tests
- Assign a single user specific cache dir (eg. ~/.cache/forex/), current behaviour: create a cache file in the current dir
