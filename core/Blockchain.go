package core

import (
	"log"
	"strings"
	"time"

	"github.com/garritfra/blockchain-project/crypto"
)

// Blockchain struct
type Blockchain struct {
	Blocks []Block
}

// AddBlock adds a block to the chain
func (bc *Blockchain) AddBlock(block Block) {

	log.Print("Mining Block...")

	// Mine Block
	for {
		hash := crypto.CalculateHash(block)
		if strings.HasPrefix(hash, "0000") {
			block.Hash = hash
			block.Timestamp = time.Now()
			block.PreviousHash = bc.GetLastHash()
			bc.Blocks = append(bc.Blocks, block)
			log.Print("Block Added: ", block.Hash)
			break
		}
		block.Nonce++
	}

}

// NewBlockchain creates a new Blockchain
func NewBlockchain() Blockchain {
	log.Print("Creating Blockchain...")

	blockchain := Blockchain{Blocks: make([]Block, 0)}

	genesisBlock := generateGenesisBlock()
	blockchain.AddBlock(genesisBlock)
	return blockchain
}

func generateGenesisBlock() Block {

	block := Block{PreviousHash: "0", Timestamp: time.Now()}
	transaction := Transaction{Amount: 0, Sender: "0", Receiver: "0", Message: "Genesis"}
	block.AddTransaction(transaction)
	log.Print("Genesis Block created")
	return block
}

// GetLastHash returns the hash of the latest block on the chain
func (bc *Blockchain) GetLastHash() string {

	bcLength := len(bc.Blocks)

	if bcLength == 0 {
		return "0"
	}
	return bc.Blocks[len(bc.Blocks)-1].Hash
}
