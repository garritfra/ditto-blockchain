package main

// Blockchain struct
type Blockchain struct {
	blocks []Block
}

func (bc *Blockchain) addBlock(block Block) {
	bc.blocks = append(bc.blocks, block)
}

// NewBlockchain creates a new Blockchain
func NewBlockchain() Blockchain {

	blockchain := Blockchain{blocks: make([]Block, 0)}

	genesisBlock := generateGenesisBlock()
	blockchain.addBlock(genesisBlock)

	return blockchain
}

func generateGenesisBlock() Block {

	transaction := Transaction{Amount: 0, Sender: "0", Receiver: "0", Message: "Genesis"}
	block := NewBlock("0", []Transaction{transaction})
	return block
}
