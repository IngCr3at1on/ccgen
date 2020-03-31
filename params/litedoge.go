package params

import "github.com/btcsuite/btcd/chaincfg"

// LiteDogeParams defines the chopped down LiteDoge parameters.
var LiteDogeParams = chaincfg.Params{
	Name: "litedoge",

	PubKeyHashAddrID: 0x5A, // starts with d
	PrivateKeyID:     0xAB, // starts with 6 (uncompressed) or S (compressed)
}
