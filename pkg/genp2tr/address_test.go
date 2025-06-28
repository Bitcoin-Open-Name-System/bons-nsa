package genp2tr

/**
You need to get the following packages before runing the test.

- go get github.com/btcsuite/btcd/chaincfg
*/

import (
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
)

func TestGetNetworkParams(t *testing.T) {
	tests := []struct {
		name       string
		input      Network
		wantParams *chaincfg.Params
		wantErr    bool
	}{
		{
			name:       "mainnet case",
			input:      Mainnet,
			wantParams: &chaincfg.MainNetParams,
			wantErr:    false,
		},
		{
			name:       "testnet case",
			input:      Testnet,
			wantParams: &chaincfg.TestNet3Params,
			wantErr:    false,
		},
		{
			name:       "regtest case",
			input:      Regtest,
			wantParams: &chaincfg.RegressionNetParams,
			wantErr:    false,
		},
		{
			name:    "invalid network",
			input:   "invalidnet",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getNetworkParams(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error for input %q, got nil", tt.input)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error for input %q: %v", tt.input, err)
				return
			}

			if got != tt.wantParams {
				t.Errorf("got %+v, want %+v", got, tt.wantParams)
			}
		})
	}

}
