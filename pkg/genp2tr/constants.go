package genp2tr

type Network string

const (
	Mainnet Network = "mainnet"
	Testnet Network = "testnet"
	Regtest Network = "regtest"
)

// Seed used to generate the Name Server Address.
// Genesis block miner address(Nakamoto Satoshi)
const PrimarySeed string = "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"
