package gen

import (
	"strings"

	"github.com/IngCr3at1on/ccgen/params"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

func GenerateAddress(ctype string, compress bool) (string, string, error) {
	priv, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", "", err
	}

	var p *chaincfg.Params
	switch {
	case strings.ToLower(ctype) == "bitcoin" || strings.ToLower(ctype) == "btc":
		p = &chaincfg.MainNetParams
		break
	case strings.ToLower(ctype) == "litecoin" || strings.ToLower(ctype) == "ltc":
		p = &params.LitecoinParams
		break
	case strings.ToLower(ctype) == "litedoge" || strings.ToLower(ctype) == "ldoge":
		p = &params.LiteDogeParams
		break
	case strings.ToLower(ctype) == "paycoin" || strings.ToLower(ctype) == "xpy":
		p = &params.PaycoinParams
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
