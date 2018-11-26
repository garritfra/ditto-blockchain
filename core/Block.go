package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
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
func NewBlock(previousHash string) Block {
	block := Block{Timestamp: time.Now(), PreviousHash: previousHash}

	block.calculateHash()

	return block
}

// AddTransaction takes in a transaction and adds it to the block
func (block *Block) AddTransaction(transaction Transaction) error {
	block.Data = append(block.Data, transaction)
	return nil
}

func (block *Block) calculateHash() {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	if err := encoder.Encode(block); err != nil {
		panic(err)
	}
	hasher := sha256.New()
	bytes := buffer.Bytes()
	hasher.Write(bytes)

	sum := hex.EncodeToString(hasher.Sum(nil))

	block.Hash = string(sum)
}
