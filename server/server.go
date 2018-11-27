package server

import (
	"encoding/gob"
	"encoding/json"
	"net/http"

	core "github.com/garritfra/blockchain-project/core"
)

var blockchain core.Blockchain

// Start starts the HTTP server on port 8080
func Start() {
	gob.Register(core.Block{})
	gob.Register(core.Transaction{})
	gob.Register(core.Blockchain{})

	blockchain = core.NewBlockchain()
	block := core.Block{}

	transaction := core.Transaction{Sender: "foo", Receiver: "bar", Amount: 100}
	block.AddTransaction(transaction)

	go blockchain.AddBlock(block)

	block = core.Block{}
	transaction = core.Transaction{Sender: "bar", Receiver: "baz", Amount: 500}
	block.AddTransaction(transaction)

	go blockchain.AddBlock(block)

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
	json.NewEncoder(w).Encode(blockchain.Blocks)
}
