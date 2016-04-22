package params

import "github.com/btcsuite/btcd/chaincfg"

// Paycoin chopped down Paycoin parameters
var Paycoin = chaincfg.Params{
	Name: "paycoin",

	PubKeyHashAddrID: 0x37,               // starts with p
	PrivateKeyID:     0x99 + 0x19 + 0x05, // hacky?
}
