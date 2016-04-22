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

func generateAddress(ctype string, compress bool) (string, string, error) {
	priv, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", "", err
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
		return "", "", err
	}

	var spub []byte
	if compress {
		spub = priv.PubKey().SerializeCompressed()
	} else {
		spub = priv.PubKey().SerializeUncompressed()
	}

	addr, err := btcutil.NewAddressPubKey(spub, p)
	if err != nil {
		return "", "", err
	}

	return wif.String(), addr.EncodeAddress(), nil
}

func searchLoop(vanity, ctype string, compress bool) (string, string, error) {
	for {
		wif, addr, err := generateAddress(ctype, compress)
		if err != nil {
			return "", "", err
		}

		if strings.HasPrefix(addr, vanity) {
			return wif, addr, nil
		}
	}
}

func init() {
	app = cli.NewApp()
	app.Name = "ccgen"
	app.Usage = "Cryptocoin address generator"
	app.Version = "0.0.1"

	var ctype string
	var vanity string
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
		cli.StringFlag{
			Name:        "vanity, V",
			Value:       "",
			Usage:       "Attempt to generate an address with the provided prefix",
			Destination: &vanity,
		},
	}

	app.Action = func(c *cli.Context) {
		var wif string
		var addr string
		var err error
		if vanity != "" {
			wif, addr, err = searchLoop(vanity, ctype, compress)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			wif, addr, err = generateAddress(ctype, compress)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		fmt.Printf("%s\n%s\n", wif, addr)
	}
}

func main() {
	app.Run(os.Args)
}
