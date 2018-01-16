package params

import "github.com/btcsuite/btcd/chaincfg"

// DenariusParams defines the chopped down Denarius parameters.
var DenariusParams = chaincfg.Params{
	Name: "denarius",

	PubKeyHashAddrID: 0x1E, // starts with D
	PrivateKeyID:     0x9E, // starts with 6 (uncompressed) or Q (compressed)
}
