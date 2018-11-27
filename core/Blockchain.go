package core

import (
	"log"
	"strings"
	"time"

	"github.com/garritfra/blockchain-project/crypto"
)

// Blockchain struct
type Blockchain struct {
	Blocks              []Block
	PendingTransactions []Transaction
}

// MineBlock adds a block to the chain
func (bc *Blockchain) MineBlock() Block {
	block := Block{}
	block.Data = bc.PendingTransactions
	block.Timestamp = time.Now()
	block.PreviousHash = bc.GetLastHash()
	// Mine Block
	log.Print("Mining Block...")
	for {
		hash := crypto.CalculateHash(block)
		if strings.HasPrefix(hash, "00000") {
			block.Hash = hash

			bc.Blocks = append(bc.Blocks, block)
			bc.PendingTransactions = []Transaction{}

			log.Print("Block Added: ", block.Hash)
			return block
		}
		block.Nonce++
	}

}

// NewBlockchain creates a new Blockchain
func NewBlockchain() Blockchain {
	log.Print("Creating Blockchain...")

	blockchain := Blockchain{Blocks: make([]Block, 0), PendingTransactions: make([]Transaction, 0)}

	// Mine Genesis Block
	blockchain.MineBlock()
	return blockchain
}

// AddTransaction takes in a transaction and adds it to the block
func (bc *Blockchain) AddTransaction(transaction Transaction) error {
	bc.PendingTransactions = append(bc.PendingTransactions, transaction)
	return nil
}

// GetLastHash returns the hash of the latest block on the chain
func (bc *Blockchain) GetLastHash() string {

	bcLength := len(bc.Blocks)

	if bcLength == 0 {
		return "0"
	}
	return bc.Blocks[len(bc.Blocks)-1].Hash
}
