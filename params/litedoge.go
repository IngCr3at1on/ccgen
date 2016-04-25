package params

import "github.com/btcsuite/btcd/chaincfg"

// LiteDogeParams defines the chopped down LiteDoge parameters.
var LiteDogeParams = chaincfg.Params{
	Name: "litedoge",

	PubKeyHashAddrID: 0x59 + 0x01, // starts with d
	PrivateKeyID:     0x99 + 0x12, // hacky?
}
