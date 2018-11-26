package main

import (
	"log"
)

// Blockchain struct
type Blockchain struct {
	blocks []Block
}

// AddBlock adds a block to the chain
func (bc *Blockchain) AddBlock(block Block) {

	block.PreviousHash = bc.GetLastHash()

	block.Hash = calculateHash(block)
	bc.blocks = append(bc.blocks, block)
	log.Print("Block added: %o", block)
}

// NewBlockchain creates a new Blockchain
func NewBlockchain() Blockchain {
	log.Print("Creating Blockchain...")

	blockchain := Blockchain{blocks: make([]Block, 0)}

	genesisBlock := generateGenesisBlock()
	blockchain.AddBlock(genesisBlock)
	return blockchain
}

func generateGenesisBlock() Block {

	block := Block{PreviousHash: "0"}
	transaction := Transaction{Amount: 0, Sender: "0", Receiver: "0", Message: "Genesis"}
	block.AddTransaction(transaction)
	log.Print("Genesis Block created")
	return block
}

// GetLastHash returns the hash of the latest block on the chain
func (bc *Blockchain) GetLastHash() string {

	bcLength := len(bc.blocks)

	if bcLength == 0 {
		return "0"
	}
	return bc.blocks[len(bc.blocks)-1].Hash
}
