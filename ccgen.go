package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/IngCr3at1on/ccgen/gen"

	"github.com/codegangsta/cli"
)

var app *cli.App

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
		if vanity != "" {
			vane := gen.NewVanityGen(ctype, vanity, compress)
			vane.Start()
		out:
			for !vane.Finished {
				select {
				case <-vane.Quit:
					vane.Wg.Wait()
					break out
				default:
				}
			}

			wif = vane.Wif
			addr = vane.Addr
		} else {
			var err error
			wif, addr, err = gen.GenerateAddress(ctype, compress)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		fmt.Printf("%s\n%s\n", wif, addr)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	app.Run(os.Args)
}
