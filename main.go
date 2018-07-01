package main

import (
	"github.com/ror0/wiki/blockchain"
)

func main() {

	bc := blockchain.NewBlockchain()

	bc.AddBlock([]byte("Test 1"))
	bc.AddBlock([]byte("Test 2"))

	for _, block := range bc.Blocks {
		block.PrintBlockInfo(true)
	}
}

