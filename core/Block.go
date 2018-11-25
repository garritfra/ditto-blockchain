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
	Hash         []byte
	PreviousHash []byte
	Data         interface{}
}

// BlockJSON is a modified Block struct for json representation
type BlockJSON struct {
	Timestamp     time.Time
	Hash          string
	PreviouseHash string
	Data          interface{}
}

// NewBlock creates a new Block
func NewBlock(previousHash []byte, data interface{}) Block {
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

// EncodeJSON encodes a Block to a JSON representable struct
func (block *Block) EncodeJSON() BlockJSON {
	return BlockJSON{Timestamp: block.Timestamp, PreviouseHash: hex.EncodeToString(block.PreviousHash), Hash: hex.EncodeToString(block.Hash), Data: block.Data}
}
