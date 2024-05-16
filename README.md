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

```sh
git clone git@github.com:sam-kenney/xc.git
make install
```

## Build

```sh
git clone git@github.com:sam-kenney/xc.git
make build
bin/xc 150 AUD --to USD
```
