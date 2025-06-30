package main

import (
	"fmt"
	"log"

	"github.com/Bitcoin-Open-Name-System/bons-nsa/pkg/genp2tr"
)

func main() {
	seed := genp2tr.PrimarySeed
	net := genp2tr.Regtest

	netParams, err := genp2tr.GetNetworkParams(net)
	if err != nil {
		log.Fatalf("Invalid network: %v", err)
	}

	address, internalPubKey, outputPubkey, err := genp2tr.CreateTaprootBurnAddress(seed, netParams)
	if err != nil {
		log.Fatalf("Failed to create burn address: %v", err)
	}

	fmt.Printf("================= | Successfully created P2TR burn address | =================\n\n")
	fmt.Printf("Burn Address: 	%s\n\n", address)
	fmt.Printf("Input Seed: 	%s\n", seed)
	fmt.Printf("Network: 	%s\n", net)
	fmt.Printf("Internal PubKey:%s\n", internalPubKey)
	fmt.Printf("Output PubKey: 	%s\n", outputPubkey)
}
