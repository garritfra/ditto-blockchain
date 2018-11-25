package main

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
	block := newBlock([]byte{}, []byte("Genesis"))
	return block
}
