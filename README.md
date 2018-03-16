# ethdag
Saves dag on disk, uses geth code.

##Installing
`go get github.com/kkibria/ethdag`

will install the package and dependency geth. Copy the file `size.go.txt` in the ethash directory of the geth source tree and rename it to `sizes.go`. This will expose the cacheSize and datasetSize functions.

##Compile
Running `go install` will compile and install the executable in $GOPATH/bin dierctory. 
