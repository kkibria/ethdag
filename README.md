# ethdag
Saves full dataset on disk, uses go-etherium code.

## Installing
`go get github.com/kkibria/ethdag`

will install the package and dependency, go-etherium.

## Compile
We need to copy the file `sizes.go` from the `ethdag/ethash` directory to the `consensus/ethash` directory of the `go-etherium` source tree. This will expose some unexported functions like cacheSize and datasetSize etc. that we need. Running make will copy the file, compile and install the executable in $GOPATH/bin directory.

`make -f $GOPATH/src/github.com/kkibria/ethdag/Makefile`
