package core

import (
	"time"

	"github.com/garritfra/blockchain-project/crypto"
)

// Block Struct
type Block struct {
	Timestamp    time.Time
	PreviousHash string
	Data         []Transaction
	Proof        int
}

// JSONBlock representation
type JSONBlock struct {
	Timestamp    time.Time     `json:"timestamp"`
	Hash         string        `json:"hash"`
	PreviousHash string        `json:"previous_hash"`
	Data         []Transaction `json:"data"`
	Proof        int           `json:"nonce"`
}

// Hash returns the hash of the block
func (block *Block) Hash() string {
	return crypto.CalculateHash(block)
}

// AsJSON is needed, because the field `hash` is calculated dynamically, and cannot be presented in the standard `Block` struct
func (block *Block) AsJSON() JSONBlock {
	return JSONBlock{Timestamp: block.Timestamp, Hash: block.Hash(), PreviousHash: block.PreviousHash, Data: block.Data, Proof: block.Proof}
}
