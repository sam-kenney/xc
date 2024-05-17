# XC

Convert between currencies.

## Usage
```sh
xc 150 AUD --to USD
```

### Parameters
- quantity `float`: The amount of currency to exchange
- from `string`: Which currency the quantity is in
- to `--to` `string`: The currency to convert to

## Installation
Get your free API token from [https://app.exchangerate-api.com/sign-up](https://app.exchangerate-api.com/sign-up).

```sh
git clone git@github.com:sam-kenney/xc.git && cd xc
export EXCHANGE_API_TOKEN="<your-token-here>"
make install
```

## Build
Requires API token from [Installation](#installation)

```sh
git clone git@github.com:sam-kenney/xc.git && cd xc
export EXCHANGE_API_TOKEN="<your-token-here>"
make build
bin/xc 150 AUD --to USD
```
