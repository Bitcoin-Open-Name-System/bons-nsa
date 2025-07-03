package genp2tr

import (
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

// VerifyAddressWithBtcutil returns true if the address is valid, false is not.
func VerifyAddressWithBtcutil(address string, netParams *chaincfg.Params) (bool, error) {
	_, err := btcutil.DecodeAddress(address, netParams)
	if err != nil {
		return false, fmt.Errorf("failed to decode address %s: %v", address, err)
	}
	return true, nil
}
