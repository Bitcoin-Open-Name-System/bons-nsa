package genp2tr

import (
	"strings"
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
)

func TestGetNetworkParams(t *testing.T) {
	tests := []struct {
		name       string
		network    Network
		wantParams *chaincfg.Params
		wantErr    bool
	}{
		{
			name:       "mainnet case",
			network:    Mainnet,
			wantParams: &chaincfg.MainNetParams,
			wantErr:    false,
		},
		{
			name:       "testnet case",
			network:    Testnet,
			wantParams: &chaincfg.TestNet3Params,
			wantErr:    false,
		},
		{
			name:       "regtest case",
			network:    Regtest,
			wantParams: &chaincfg.RegressionNetParams,
			wantErr:    false,
		},
		{
			name:    "invalid network",
			network: "invalidnet",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getNetworkParams(tt.network)
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error for network %q, got nil", tt.network)
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error for network %q: %v", tt.network, err)
				return
			}

			if got != tt.wantParams {
				t.Errorf("got %+v, want %+v", got, tt.wantParams)
			}
		})
	}

}

func TestCreateTaprootBurnAddress(t *testing.T) {
	tests := []struct {
		name            string
		arbitraryString string
		network         *chaincfg.Params
		wantAddrPrefix  string
		wantErr         bool
	}{
		{
			name:            "valid naminet input with primary seed",
			arbitraryString: PrimarySeed,
			network:         &chaincfg.MainNetParams,
			wantAddrPrefix:  "bc1p",
			wantErr:         false,
		},
		{
			name:            "valid naminet input with primary seed",
			arbitraryString: PrimarySeed,
			network:         &chaincfg.TestNet3Params,
			wantAddrPrefix:  "tb1p",
			wantErr:         false,
		}, {
			name:            "valid test input with primary seed",
			arbitraryString: PrimarySeed,
			network:         &chaincfg.RegressionNetParams,
			wantAddrPrefix:  "bcrt1p",
			wantErr:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addr, internal, output, err := createTaprootBurnAddress(tt.arbitraryString, tt.network)

			if tt.wantErr && err == nil {
				t.Errorf("expected error but got nil for %s", tt.name)
			}
			if !tt.wantErr && err != nil {
				t.Errorf("unexpected error for %s: %v", tt.name, err)
			}
			if !tt.wantErr {
				if !strings.HasPrefix(addr, tt.wantAddrPrefix) {
					t.Errorf("expected address prefix %s, got %s", tt.wantAddrPrefix, addr)
				}
				if internal == "" || output == "" {
					t.Errorf("expected non-empty internal/output pubkeys, got internal: %s, output: %s", internal, output)
				}
			}
		})
	}
}
