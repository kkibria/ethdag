mkfile_dir := $(dir $(realpath $(firstword $(MAKEFILE_LIST))))
ethash_dir := $(GOPATH)/src/github.com/ethereum/go-ethereum/consensus/ethash

default:
	echo $(mkfile_dir)
	cd $(mkfile_dir); cp ethhash/sizes.go $(ethash_dir); go build; go install