# ethdag
Saves dag on disk, uses geth code.

## Installing
`go get github.com/kkibria/ethdag`

will install the package and dependency geth. Copy the file `sizes.go` from the `ethdag/ethash` directory to the `ethash` directory of the geth source tree. This will expose the cacheSize and datasetSize functions we need.

## Compile
Running `go install` will compile and install the executable in $GOPATH/bin directory. 
