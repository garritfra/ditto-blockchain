package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	blockchain := newBlockchain()

	block := newBlock(blockchain.blocks[0].Hash, []byte("I am the first block!!"))
	blockchain.addBlock(block)

	block = newBlock(blockchain.blocks[1].Hash, []byte("This is number two"))
	blockchain.addBlock(block)

	for i := 0; i < len(blockchain.blocks); i++ {

		block := blockchain.blocks[i]
		fmt.Println("Block " + string(i))
		fmt.Println(hex.EncodeToString(block.PreviousHash))
		fmt.Println(hex.EncodeToString(block.Hash))
		fmt.Println()
	}

}
