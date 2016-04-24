Simple(ish) semi-universal address and vanity address generator for cryptocurrency

```
NAME:
   ccgen - Cryptocoin address generator

USAGE:
   ccgen [global options] command [command options] [arguments...]

VERSION:
   0.0.2

COMMANDS:
GLOBAL OPTIONS:
   --type, -t "bitcoin"	Define coin-type to generate an address for
   --compress, -c	Compress the private key and address
   --vanity, -V 	Attempt to generate an address with the provided prefix
   --help, -h		show help
   --version, -v	print the version
```

Currently supports Bitcoin, Paycoin and LiteDoge. Pull requests for coin support are welcome.
