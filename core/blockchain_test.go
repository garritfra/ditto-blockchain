package core

import (
	"testing"
)

func Test_generateGenesisBlock(t *testing.T) {
	genesisBlock := generateGenesisBlock()

	if len(genesisBlock.PreviousHash) > 1 {
		t.Errorf("Previous hash of Genesis Block is not empty!")
	}
}

func TestNewBlockchain(t *testing.T) {
	blockchain := NewBlockchain()

	if len(blockchain.Blocks) != 1 {
		t.Errorf("New Blockchain has too many Blocks (Besides Genesis)")
	}
}

func TestBlockchain_GetLastHash(t *testing.T) {
	blockchain := NewBlockchain()

	blockchain.AddBlock(Block{})
	want := blockchain.Blocks[1].Hash
	got := blockchain.GetLastHash()
	if got != want {
		t.Errorf("Want: %v, but got %v", want, got)
	}
}
