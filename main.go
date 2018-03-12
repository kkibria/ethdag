package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/consensus/ethash"
)

func main() {
	var block uint64 = 5140155
	epoch := int(block / 30000)

	fmt.Fprintln(os.Stdout, "Epoch = ", epoch)
	seedhash := ethash.SeedHash(block)
	fmt.Fprintln(os.Stdout, seedhash)
	ethash.MakeDataset(block, "testcache")
}
