package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/IngCr3at1on/ccgen/params"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/codegangsta/cli"
)

var app *cli.App

func init() {
	app = cli.NewApp()
	app.Name = "ccgen"
	app.Usage = "Cryptocoin address generator"
	app.Version = "0.0.1"

	var ctype string
	var compress bool

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "type, t",
			Value:       "bitcoin",
			Usage:       "Define coin-type to generate an address for",
			Destination: &ctype,
		},
		cli.BoolFlag{
			Name:        "compress, c",
			Usage:       "Compress the private key and address",
			Destination: &compress,
		},
	}

	app.Action = func(c *cli.Context) {
		priv, err := btcec.NewPrivateKey(btcec.S256())
		if err != nil {
			fmt.Println(err)
			return
		}

		var p *chaincfg.Params
		switch {
		case strings.ToLower(ctype) == "bitcoin" || strings.ToLower(ctype) == "btc":
			p = &chaincfg.MainNetParams
			break
		case strings.ToLower(ctype) == "litedoge" || strings.ToLower(ctype) == "ldoge":
			p = &params.Litedoge
			break
		case strings.ToLower(ctype) == "paycoin" || strings.ToLower(ctype) == "xpy":
			p = &params.Paycoin
			break
		default:
			p = &chaincfg.MainNetParams
			break
		}

		wif, err := btcutil.NewWIF(priv, p, compress)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s\n", wif.String())

		var spub []byte
		if compress {
			spub = priv.PubKey().SerializeCompressed()
		} else {
			spub = priv.PubKey().SerializeUncompressed()
		}

		addr, err := btcutil.NewAddressPubKey(spub, p)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s\n", addr.EncodeAddress())
	}
}

func main() {
	app.Run(os.Args)
}
