package main

import (
	"flag"
	"fmt"
	"kkutils"
	"math"
	"os"
	"time"

	eth "github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/kkibria/ethdag/dataset"
)

var (
	endian     string
	pStatsFlag *bool
	pOutDir    *string
	seedhash   []byte
)

func main() {
	var err error

	defer kkutils.TimeTrack(time.Now(), "ethdag")

	pDatasetCnt := flag.Uint64("r", 1, "Number of datasets to generate. Maximum 16 datasets.")
	pBlock := flag.Uint64("b", math.MaxUint64, "Block number. Epoch number must not be specified.")
	pEpoch := flag.Uint64("e", math.MaxUint64, "Epoch number. Block number must not be specified.")
	pStatsFlag = flag.Bool("s", false, "Stats only, does not generate dataset.")
	pOutDir = flag.String("o", "eth-dataset", "Output directory.")

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("    -r [-b block|-e epoch] -s -o outDir\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	if (flag.NArg() != 0) || (*pBlock == *pEpoch) || (*pDatasetCnt < 1) || (*pDatasetCnt > 16) {
		flag.Usage()
	}

	if (*pBlock == math.MaxUint64) && (*pEpoch == math.MaxUint64) {
		flag.Usage()
	}

	if (*pBlock != math.MaxUint64) && (*pEpoch != math.MaxUint64) {
		flag.Usage()
	}

	if *pEpoch != math.MaxUint64 {
		*pBlock = *pEpoch * dataset.BlocksPerEpoch
	} else {
		*pEpoch = *pBlock / dataset.BlocksPerEpoch
	}

	fmt.Fprintf(os.Stdout, "Block = %v\n", *pBlock)
	fmt.Fprintf(os.Stdout, "Epoch = %v\n", *pEpoch)

	if !kkutils.IsLittleEndian() {
		endian = ".be"
	}

	for i := uint64(0); i < *pDatasetCnt; i++ {
		e := *pEpoch + i
		b := e * dataset.BlocksPerEpoch
		// sizes
		seedhash = eth.SeedHash(b)
		fmt.Fprintf(os.Stdout, "\nseedhash(%v) = %x\n", e, seedhash)
		fmt.Fprintf(os.Stdout, "Cache size = %v\n", eth.CacheSize(b))
		fmt.Fprintf(os.Stdout, "Dataset size = %v\n", eth.DatasetSize(b))
		if !*pStatsFlag {
			if err = dataset.Dataset(e, *pOutDir); err != nil {
				fmt.Println(err)
			}
		}
	}
}
