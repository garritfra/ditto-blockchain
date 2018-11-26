package main

// Blockchain struct
type Blockchain struct {
	blocks []Block
}

// AddBlock adds a block to the chain
func (bc *Blockchain) AddBlock(block Block) {
	bc.blocks = append(bc.blocks, block)
}

// NewBlockchain creates a new Blockchain
func NewBlockchain() Blockchain {

	blockchain := Blockchain{blocks: make([]Block, 0)}

	genesisBlock := generateGenesisBlock()
	blockchain.AddBlock(genesisBlock)

	return blockchain
}

func generateGenesisBlock() Block {

	block := NewBlock("0")
	transaction := Transaction{Amount: 0, Sender: "0", Receiver: "0", Message: "Genesis"}
	block.AddTransaction(transaction)
	return block
}

// GetLastBlock returns the latest block on the chain
func (bc *Blockchain) GetLastBlock() Block {
	return bc.blocks[len(bc.blocks)-1]
}
