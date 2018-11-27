package server

import (
	"encoding/gob"
	"encoding/json"
	"log"
	"net/http"

	"github.com/garritfra/blockchain-project/core"
)

var blockchain core.Blockchain

// ServeHTTP serves a http server on a given port with format ":PORT"
func ServeHTTP(port string) {
	gob.Register(core.Block{})
	gob.Register(core.Transaction{})
	gob.Register(core.Blockchain{})

	blockchain = core.NewBlockchain()

	http.HandleFunc("/", handleMain)
	http.HandleFunc("/mine_block", handleMineBlock)
	http.HandleFunc("/pending_transactions", handleListPendingTransactions)

	log.Print("Listening on port ", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

func handleMain(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleListBlocks(w, r)
	case "POST":
		handleAddTransaction(w, r)
	}
}

func handleError(err error, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Error: " + err.Error()))
	log.Print(err.Error())
}

func handleListBlocks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(blockchain.Blocks)
}

func handleAddTransaction(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var receivedTransaction core.Transaction
	err := decoder.Decode(&receivedTransaction)
	if err == nil {
		log.Println("Transaction from", receivedTransaction.Sender, "to", receivedTransaction.Receiver, "received")

		go blockchain.AddTransaction(receivedTransaction)
	} else {
		handleError(err, w, r)
	}
}

func handleMineBlock(w http.ResponseWriter, r *http.Request) {
	block := blockchain.MineBlock()
	json.NewEncoder(w).Encode(block)
}

func handleListPendingTransactions(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(blockchain.PendingTransactions)
}
