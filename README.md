A currency CLI tool with caching which allows to type in a source currency and
multiple target currencies. USD to currency conversion rates are fetched from `currencyapi` API and are cached.
Source to target currency is then calculated based on the USD to each currency rate.

Caching in an app that converts currencies while their values keep changing every second is not ideal. The reason for caching functionality is to satisfy the API rate quota in the free tier. The app was intended to be general CLI tool for currency conversion. The cache duration can be set to zero for real time use when building the binary and a paid API tier can be used.

$$
converted amount := amount \times \frac{target currency rate w.r.t USD}{source currency rate w.r.t USD}
$$

Example usage:

```bash

forex 15 GBP INR SAR

```

Example output:

```bash

Version: 1.0.0
GBP 15.00 equals:

INR 1677.79
SAR 75.15

```

When building binary, specify cacheExpiryDurationInHours.

```bash

go build -ldflags "-X 'main.cacheExpiryDurationInHours=12' -X 'main.version=1.0.0'" -o forex

```

After building move binary to `/usr/local/bin/`, this will allow you to call it from any directory

Todo

- Write unit tests
- Assign a single user specific cache dir (eg. ~/.cache/forex/), current behaviour: create a cache file in the current dir
- Remove api key and move to .env
