package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

var blockchain Blockchain

// StartServer starts the HTTP server on port 8080
func StartServer() {

	blockchain = newBlockchain()

	data, _ := json.Marshal(Transaction{Sender: "foo", Receiver: "bar", Amount: 100})
	block := newBlock(blockchain.blocks[0].Hash, data)
	blockchain.addBlock(block)

	data2, _ := json.Marshal(Transaction{Sender: "bar", Receiver: "baz", Amount: 5000})
	block = newBlock(blockchain.blocks[1].Hash, data2)
	blockchain.addBlock(block)

	http.HandleFunc("/", sayHello)
	http.HandleFunc("/blockchain", listBlocks)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}

func listBlocks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(blockchain.blocks)
}
