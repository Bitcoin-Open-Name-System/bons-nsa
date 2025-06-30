package genp2tr

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

// getNetworkParams returns the chain configuration params for the given Bitcoin network.
func getNetworkParams(network Network) (*chaincfg.Params, error) {
	switch network {
	case Mainnet:
		return &chaincfg.MainNetParams, nil
	case Testnet:
		return &chaincfg.TestNet3Params, nil
	case Regtest:
		return &chaincfg.RegressionNetParams, nil
	default:
		return nil, fmt.Errorf("invalid network: %s", network)
	}
}

func taggedHash(tag string, data []byte) []byte {
	tagHash := sha256.Sum256([]byte(tag))
	h := sha256.New()
	h.Write(tagHash[:])
	h.Write(tagHash[:])
	h.Write(data)

	return h.Sum(nil)
}

// createTaprootBurnAddress returns taproot burn address, x only internalkey(encoded to string), output publickey, error
func createTaprootBurnAddress(arbitraryString string, network *chaincfg.Params) (string, string, string, error) {

	hash := sha256.Sum256([]byte(arbitraryString))
	internalKeyBytes := hash[:]

	var internalPubKey *btcec.PublicKey
	var err error

	for i := 0; i < 256; i++ {
		internalPubKey, err = btcec.ParsePubKey(append([]byte{0x02}, internalKeyBytes...))
		if err == nil {
			break
		}
		hashInput := append([]byte(arbitraryString), byte(i))
		hash = sha256.Sum256(hashInput)
		internalKeyBytes = hash[:]
	}

	if err != nil {
		return "", "", "", fmt.Errorf("failed to create internal pub key: %v", err)
	}

	// tweaking
	internalKeyXOnly := internalPubKey.SerializeCompressed()[1:] // x-only (remove the first byte)
	tweakData := taggedHash("TapTweak", internalKeyXOnly)

	// tweak to scalar
	var tweakScalar btcec.ModNScalar
	tweakScalar.SetBytes((*[32]byte)(tweakData))

	// tweaked pub key : output public key = internal_pubkey + tweak * G
	var outputKey btcec.JacobianPoint
	btcec.ScalarBaseMultNonConst(&tweakScalar, &outputKey)

	// internal key to Jacobian
	var internalKeyJacobian btcec.JacobianPoint
	internalPubKey.AsJacobian(&internalKeyJacobian)

	btcec.AddNonConst(&internalKeyJacobian, &outputKey, &outputKey)

	outputKey.ToAffine()

	outputPubKeyBytes := outputKey.X.Bytes()
	var outputPubKey32 [32]byte
	copy(outputPubKey32[32-len(outputPubKeyBytes):], outputPubKeyBytes[:])

	// Create the taproot address
	taprootAddr, err := btcutil.NewAddressTaproot(outputPubKey32[:], network)
	if err != nil {
		return "", "", "", fmt.Errorf("failed to create taproot addr: %v", err)
	}

	return taprootAddr.String(), hex.EncodeToString(internalKeyXOnly), hex.EncodeToString(outputPubKey32[:]), nil
}
