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
	Nonce        int
}

// AddTransaction takes in a transaction and adds it to the block
func (block *Block) AddTransaction(transaction Transaction) error {
	block.Data = append(block.Data, transaction)
	return nil
}
