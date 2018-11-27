package core

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
