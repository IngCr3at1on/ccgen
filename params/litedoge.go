package params

import "github.com/btcsuite/btcd/chaincfg"

// Litedoge chopped down parameters
var Litedoge = chaincfg.Params{
	Name: "litedoge",

	PubKeyHashAddrID: 0x59 + 0x01, // starts with d
	PrivateKeyID:     0x99 + 0x12, // hacky?
}
