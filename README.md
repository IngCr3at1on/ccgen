# ccgen - Cryptocoin Address Generator

![Version](https://img.shields.io/badge/version-0.0.2-blue.svg?style=flat-square)
![MIT License](http://img.shields.io/badge/license-MIT-green.svg?style=flat-square)

ccgen is a simple(ish) semi-universal address and vanity address generator for cryptocurrency.

## Supported Coins

Current List of Available Coins for Address Generation  
-----
|**Coin Short** | **Coin Name** | **Address Prefix**  |
| --------------------------------------- | -------------------------------------------- | ------------ |
|BTC | Bitcoin | 1  |
|DNR | Denarius | D  |
|LDOGE | LiteDoge | d  |
|LTC | Litecoin | L  |
|XPY | Paycoin | P  |

*Pull requests for coin support are welcome.*

## Usage

```
Usage: ccgen [options] <command> [<arguments...>]
```

## Options

```
--type, -t              Define coin-type to generate an address for (eg. "bitcoin")
--compress, -c          Compress the private key and address
--vanity, -V            Attempt to generate an address with the provided prefix
--help, -h              Show help
--version, -v           Print the version
```

## License

The MIT License
