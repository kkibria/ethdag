package dataset

import (
	"fmt"
	"os"
	"path/filepath"

	eth "github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/kkibria/kkutils"
)

const (
	// BlocksPerEpoch is the number of blocks per epoch
	BlocksPerEpoch = 30000
)

// Dataset generates a raw dataset file in a directory
func Dataset(epoch uint64, dir string) error {
	b := epoch * BlocksPerEpoch
	eth.MakeDatasetFinalize(b, dir)

	// fix the name
	var endian string
	if !kkutils.IsLittleEndian() {
		endian = ".be"
	}
	seedhash := eth.SeedHash(b)
	filename := fmt.Sprintf("full-R%d-%x%s", eth.GetRev(), seedhash[:8], endian)
	oldPath := filepath.Join(dir, filename)
	newPath := filepath.Join(dir, fmt.Sprintf("epoch-%v-full", epoch))
	return os.Rename(oldPath, newPath)
}
