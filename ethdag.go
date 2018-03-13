package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/consensus/ethash"
)

const (
	blocksPerEpoch = 30000
)

func main() {
	repeatPtr := flag.Uint64("r", 1, "Number of dags to generate. Maximum 16 dags.")
	blockFlagPtr := flag.Bool("b", false, "Set this flag if the number is a block number.")
	epochFlagPtr := flag.Bool("e", true, "Set this flag if the number is an epoch.")
	outDirPtr := flag.String("o", "outDir", "Output directory.")

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("    -r -b|e -o outDir number\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	if (flag.NArg() != 1) || (*blockFlagPtr == *epochFlagPtr) || (*repeatPtr < 1) || (*repeatPtr > 16) {
		flag.Usage()
	}

	var err error
	var num uint64

	if num, err = strconv.ParseUint(flag.Arg(0), 10, 64); err != nil {
		flag.Usage()
	}

	var block uint64
	var epoch uint64

	if *epochFlagPtr {
		epoch = num
		block = epoch * blocksPerEpoch
	}

	if *blockFlagPtr {
		block = num
		epoch = block / blocksPerEpoch
	}

	fmt.Fprintf(os.Stdout, "Block = %v\n", block)
	fmt.Fprintf(os.Stdout, "Epoch = %v\n", epoch)
	seedhash := ethash.SeedHash(block)
	fmt.Fprintf(os.Stdout, "seedhash = %v\n", seedhash)

	for i := uint64(0); i < *repeatPtr; i++ {
		ethash.MakeDataset(epoch*blocksPerEpoch+i, *outDirPtr)
	}

}
