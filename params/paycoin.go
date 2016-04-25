package params

import "github.com/btcsuite/btcd/chaincfg"

// PaycoinParams defines the chopped down Paycoin parameters.
var PaycoinParams = chaincfg.Params{
	Name: "paycoin",

	PubKeyHashAddrID: 0x37,               // starts with P
	PrivateKeyID:     0x99 + 0x19 + 0x05, // hacky?
}
