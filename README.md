# ethdag
Saves full dataset on disk, uses go-etherium code.

## Installing
`go get github.com/kkibria/ethdag`

will install the package and dependency, go-etherium. Copy the file `sizes.go` from the `ethdag/ethash` directory to the `ethash` directory of the `go-etherium` source tree. This will expose some unexported functions like cacheSize and datasetSize etc. that we need.

## Compile
Running `go install` will compile and install the executable in $GOPATH/bin directory. 
