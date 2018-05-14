default:
	cp ethhash/sizes.g $(GOPATH)/src/github.com/ethereum/go-ethereum/consensus/ethash
	go build