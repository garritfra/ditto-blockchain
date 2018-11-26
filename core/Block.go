package main

import (
	"time"
)

// Block Struct
type Block struct {
	Timestamp    time.Time
	Hash         string
	PreviousHash string
	Data         []Transaction
}

// NewBlock creates a new Block
func NewBlock() Block {
	block := Block{}
	return block
}

// AddTransaction takes in a transaction and adds it to the block
func (block *Block) AddTransaction(transaction Transaction) error {
	block.Data = append(block.Data, transaction)
	return nil
}
