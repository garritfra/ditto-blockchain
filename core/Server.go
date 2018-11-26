package main

import (
	"encoding/gob"
	"encoding/json"
	"net/http"
)

var blockchain Blockchain

// StartServer starts the HTTP server on port 8080
func StartServer() {

	gob.Register(Block{})
	gob.Register(Transaction{})
	gob.Register(Blockchain{})

	blockchain = NewBlockchain()
	block := Block{}

	transaction := Transaction{Sender: "foo", Receiver: "bar", Amount: 100}
	block.AddTransaction(transaction)

	blockchain.AddBlock(block)

	http.HandleFunc("/", redirect)
	http.HandleFunc("/blockchain", listBlocks)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/blockchain", 301)
}

func listBlocks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(blockchain.blocks)
}
