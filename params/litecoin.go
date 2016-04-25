package params

import "github.com/btcsuite/btcd/chaincfg"

// LitecoinParams defines the chopped down Litecoin parameters.
var LitecoinParams = chaincfg.Params{
	Name: "litecoin",

	PubKeyHashAddrID: 0x30, // starts with L
	PrivateKeyID:     0xB0, // starts with 6 (uncompressed) or S (compressed)
}
