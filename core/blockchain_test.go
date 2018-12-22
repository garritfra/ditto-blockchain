package core

import (
	"testing"
)

func Test_hasGenesisBlock(t *testing.T) {
	genesisBlock := NewBlockchain().Blocks[0]

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

	blockchain.MineBlock()
	want := blockchain.Blocks[1].Hash()
	got := blockchain.GetLastHash()
	if got != want {
		t.Errorf("Want: %v, but got %v", want, got)
	}
}

func TestBlockchain_IsValid(t *testing.T) {
	blockchain := NewBlockchain()
	blockchain.MineBlock()

	want := true
	got := blockchain.IsValid()
	if got != want {
		t.Errorf("Want: %v, but got %v", want, got)
	}
}

func TestBlockchain_IsNotValid(t *testing.T) {
	blockchain := NewBlockchain()
	blockchain.MineBlock()
	blockchain.MineBlock()
	blockchain.Blocks[1].Proof++

	want := false
	got := blockchain.IsValid()
	if got != want {
		t.Errorf("Want: %v, but got %v", want, got)
	}
}
