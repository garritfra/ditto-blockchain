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

	http.HandleFunc("/", handleListBlocks)
	log.Print("Listening on port ", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

func handleListBlocks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(blockchain.Blocks)
}
