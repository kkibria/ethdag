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
	var err error

	pDagCnt := flag.Uint64("r", 1, "Number of dags to generate. Maximum 16 dags.")
	pBlockFlag := flag.Bool("b", false, "Set this flag if the number is a block number.")
	pEpochFlag := flag.Bool("e", false, "Set this flag if the number is an epoch.")
	pStatsFlag := flag.Bool("s", false, "Stats only, do not generate.")
	pOutDir := flag.String("o", "ethdag", "Output directory prefix.")

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("    -r -b|e -s -o outDir number\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	if (flag.NArg() != 1) || (*pBlockFlag == *pEpochFlag) || (*pDagCnt < 1) || (*pDagCnt > 16) {
		flag.Usage()
	}

	var num uint64

	if num, err = strconv.ParseUint(flag.Arg(0), 10, 64); err != nil {
		flag.Usage()
	}

	var block uint64
	var epoch uint64

	if *pEpochFlag {
		epoch = num
		block = epoch * blocksPerEpoch
	}

	if *pBlockFlag {
		block = num
		epoch = block / blocksPerEpoch
	}

	fmt.Fprintf(os.Stdout, "Block = %v\n", block)
	fmt.Fprintf(os.Stdout, "Epoch = %v\n", epoch)

	// sizes
	fmt.Fprintf(os.Stdout, "Cache size = %v\n", ethash.CacheSize(block))
	fmt.Fprintf(os.Stdout, "Dataset size = %v\n", ethash.DatasetSize(block))

	seedhash := ethash.SeedHash(block)
	fmt.Fprintf(os.Stdout, "seedhash = %v\n", seedhash)

	if !*pStatsFlag {
		for i := uint64(0); i < *pDagCnt; i++ {
			e := epoch + i
			ethash.MakeDataset((e)*blocksPerEpoch, fmt.Sprintf("%s-%v", *pOutDir, e))
		}
	}
}
