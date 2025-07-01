package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Bitcoin-Open-Name-System/bons-nsa/pkg/genp2tr"
)

func main() {
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	seed := createCmd.String("s", genp2tr.PrimarySeed, "seed string for address generation")
	net := createCmd.String("n", string(genp2tr.Mainnet), "network to use (mainnet, testnet, regtest)")

	if len(os.Args) < 2 {
		fmt.Printf("asdfasdf\n")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		createCmd.Parse(os.Args[2:])

		netParams, err := genp2tr.GetNetworkParams(genp2tr.Network(*net))
		if err != nil {
			log.Fatalf("Invalid network: %v\n", err)
		}

		address, internalPubKey, outputPubkey, err := genp2tr.CreateTaprootBurnAddress(*seed, netParams)
		if err != nil {
			log.Fatalf("Failed to create burn address: %v\n", err)
		}

		fmt.Printf("================= | Successfully created P2TR burn address | =================\n\n")
		fmt.Printf("Burn Address	: 	%s\n\n", address)
		fmt.Printf("Input Seed	: 	%s\n", *seed)
		fmt.Printf("Network		: 	%s\n", *net)
		fmt.Printf("Internal PubKey	:	%s\n", internalPubKey)
		fmt.Printf("Output PubKey	: 	%s\n", outputPubkey)
	default:
		fmt.Printf("command %s not found\n", os.Args[1])
		os.Exit(1)
	}
}
