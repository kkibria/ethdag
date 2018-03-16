package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"unsafe"

	eth "github.com/ethereum/go-ethereum/consensus/ethash"
)

const (
	blocksPerEpoch = 30000
)

var (
	endian     string
	pStatsFlag *bool
	pOutDir    *string
	seedhash   []byte
)

func main() {
	var err error

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
		*pBlock = *pEpoch * blocksPerEpoch
	} else {
		*pEpoch = *pBlock / blocksPerEpoch
	}

	fmt.Fprintf(os.Stdout, "Block = %v\n", *pBlock)
	fmt.Fprintf(os.Stdout, "Epoch = %v\n", *pEpoch)

	if !isLittleEndian() {
		endian = ".be"
	}

	for i := uint64(0); i < *pDatasetCnt; i++ {
		e := *pEpoch + i
		b := e * blocksPerEpoch
		// sizes
		seedhash = eth.SeedHash(b)
		fmt.Fprintf(os.Stdout, "\nseedhash(%v) = %x\n", e, seedhash)
		fmt.Fprintf(os.Stdout, "Cache size = %v\n", eth.CacheSize(b))
		fmt.Fprintf(os.Stdout, "Dataset size = %v\n", eth.DatasetSize(b))
		if !*pStatsFlag {
			eth.MakeDatasetFinalize(b, *pOutDir)
			if err = fixName(e); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func isLittleEndian() bool {
	n := uint32(0x01020304)
	return *(*byte)(unsafe.Pointer(&n)) == 0x04
}

func fixName(e uint64) error {
	filename := fmt.Sprintf("full-R%d-%x%s", eth.GetRev(), seedhash[:8], endian)
	oldPath := filepath.Join(*pOutDir, filename)
	newPath := filepath.Join(*pOutDir, fmt.Sprintf("epoch-%v-full", e))
	return os.Rename(oldPath, newPath)
}
