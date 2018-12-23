package core

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

// Blockchain struct
type Blockchain struct {
	Blocks              []Block
	PendingTransactions []Transaction
	Peers               []string
}

// MineBlock adds a block to the chain
func (bc *Blockchain) MineBlock() Block {
	block := Block{}
	block.Data = bc.PendingTransactions
	block.Timestamp = time.Now()
	block.PreviousHash = bc.GetLastHash()
	// Mine Block
	log.Print("Mining Block...")
	for {
		if strings.HasPrefix(block.Hash(), "0000") {

			bc.Blocks = append(bc.Blocks, block)
			bc.PendingTransactions = []Transaction{}

			bc.NotifyPeers()
			log.Print("Block Added: ", block.Hash())
			return block
		}
		block.Proof++
	}
}

// NewBlockchain creates a new Blockchain
func NewBlockchain() Blockchain {
	log.Print("Creating Blockchain...")

	blockchain := Blockchain{Blocks: make([]Block, 0), PendingTransactions: make([]Transaction, 0), Peers: make([]string, 0)}

	// Mine Genesis Block
	blockchain.MineBlock()
	return blockchain
}

// AddTransaction takes in a transaction and adds it to the block
func (bc *Blockchain) AddTransaction(transaction Transaction) error {
	bc.PendingTransactions = append(bc.PendingTransactions, transaction)
	return nil
}

// GetLastHash returns the hash of the latest block on the chain
func (bc *Blockchain) GetLastHash() string {

	bcLength := len(bc.Blocks)

	if bcLength == 0 {
		return "0"
	}
	return bc.Blocks[len(bc.Blocks)-1].Hash()
}

// AddPeer appends the IP address to the list of peers known to the chain
func (bc *Blockchain) AddPeer(peer string) (resp *http.Response, err error) {
	resp, err = http.Get("http://" + peer)

	// Return Error, if an error occured
	if err != nil {
		return nil, err
	}

	// Add Peer, if it is not already in the list
	if !contains(bc.Peers, peer) {
		bc.Peers = append(bc.Peers, peer)
	}
	return resp, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// IsValid checks, if the chain has any faulty blocks
func (bc Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		if bc.Blocks[i-1].Hash() != bc.Blocks[i].PreviousHash {
			return false
		}
	}

	return true
}

// NotifyPeers requests an update to all peers
func (bc *Blockchain) NotifyPeers() {
	for _, peer := range bc.Peers {
		http.Get("http://" + peer + "/update")
	}
}

// Update compares the length of the own chain to the blocklength of every known peer, and updates the chain accordingly
func (bc *Blockchain) Update() bool {

	var hasBeenReplaced = false

	for _, peer := range bc.Peers {
		resp, err := http.Get("http://" + peer)

		decoder := json.NewDecoder(resp.Body)
		var receivedBlockchain JSONBlockchain
		err = decoder.Decode(&receivedBlockchain)
		if err == nil {
			if receivedBlockchain.Blockcount > bc.AsJSON().Blockcount && receivedBlockchain.FromJSON().IsValid() {
				bc.replaceChain(receivedBlockchain.FromJSON())
				log.Println("Blockchain has been updated by " + peer)
				hasBeenReplaced = true
			}
		}
	}
	return hasBeenReplaced
}

func (bc *Blockchain) replaceChain(newChain Blockchain) {
	newBlocks := make([]Block, 0)
	for _, block := range newChain.Blocks {
		newBlocks = append(newBlocks, block)
	}

	bc.Blocks = newBlocks
	bc.NotifyPeers()
}

// JSONBlockchain is needed, because the hash of each block is calculated dynamically, and therefore is not stored in the `Block` struct
type JSONBlockchain struct {
	Blocks              []JSONBlock
	Blockcount          int
	PendingTransactions []Transaction
	Peers               []string
}

// AsJSON returns the Blockchain as a JSON Blockchain
func (bc *Blockchain) AsJSON() JSONBlockchain {
	jsonChain := JSONBlockchain{PendingTransactions: bc.PendingTransactions, Peers: bc.Peers}

	for _, block := range bc.Blocks {
		jsonChain.Blocks = append(jsonChain.Blocks, block.AsJSON())
	}
	jsonChain.Blockcount = len(jsonChain.Blocks)

	return jsonChain
}

// FromJSON casts a `JSONBlockchain` to a regular `Blockchain` struct
func (bc *JSONBlockchain) FromJSON() Blockchain {
	chain := Blockchain{PendingTransactions: bc.PendingTransactions, Peers: bc.Peers}
	for _, block := range bc.Blocks {
		chain.Blocks = append(chain.Blocks, block.FromJSON())
	}

	return chain
}
