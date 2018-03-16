// place this file in the ethash directory of the go-ethereum source tree
package ethash

import "math"

// CacheSize is exported version of cacheSize
func CacheSize(block uint64) uint64 {
	return cacheSize(block)
}

// DatasetSize is exported version of datasetSize
func DatasetSize(block uint64) uint64 {
	return datasetSize(block)
}

// GetRev returns algorith revision
func GetRev() int {
	return algorithmRevision
}

// MakeDatasetFinalize is same as MakeDataset(). It generates a new ethash dataset and
// optionally stores it to disk. It additionally finalizes to close the file.
func MakeDatasetFinalize(block uint64, dir string) {
	d := dataset{epoch: block / epochLength}
	d.generate(dir, math.MaxInt32, false)
	d.finalizer()
}
