package genp2tr

import (
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
)

func TestIsValidAddressWithBtcutil(t *testing.T) {
	tests := []struct {
		name      string
		address   string
		network   *chaincfg.Params
		wantValid bool
		wantErr   bool
	}{
		{
			name:      "valid mainnet address",
			address:   "bc1p8h9rzsy6puu6key4kykj7n4g2s0nw0ngw00fdmkxtedtk44ql73qdmys8a",
			network:   &chaincfg.MainNetParams,
			wantValid: true,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid, err := IsValidAddressWithBtcutil(tt.address, tt.network)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got nil for %s", tt.name)
				}
			} else {
				if err != nil {
					t.Errorf("did not expect error but got: %v", err)
				}
			}

			if valid != tt.wantValid {
				t.Errorf("expected valid=%v, got %v", tt.wantValid, valid)
			}
		})
	}
}
