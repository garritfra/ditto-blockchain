package main

import (
	"crypto/sha256"
	"fmt"
)

// Blockchain struct
type Blockchain struct {
	blocks []Block
}

func (bc *Blockchain) addBlock(block Block) {
	bc.blocks = append(bc.blocks, block)
}

func newBlockchain() Blockchain {

	blockchain := Blockchain{blocks: make([]Block, 0)}

	genesisBlock := generateGenesisBlock()
	blockchain.addBlock(genesisBlock)

	return blockchain
}

func generateGenesisBlock() Block {
	block := new(Block)
	hasher := sha256.New()
	s := fmt.Sprintf("%v", block)
	sum := hasher.Sum([]byte(s))
	block.hash = sum

	return *block
}
