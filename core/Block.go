package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
)

// Block Struct
type Block struct {
	Timestamp    time.Time
	Hash         []byte
	PreviousHash []byte
	Data         []byte
}

func newBlock(previousHash []byte, data []byte) Block {
	block := Block{Timestamp: time.Now(), PreviousHash: previousHash, Data: data}

	block.calculateHash()

	return block
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

	sum := hasher.Sum(nil)

	block.Hash = sum
}
