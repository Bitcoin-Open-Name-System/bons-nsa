package genp2tr

import (
	"fmt"

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
