default:
	cp ethhash/sizes.go $(GOPATH)/src/github.com/ethereum/go-ethereum/consensus/ethash
	go build