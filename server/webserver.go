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
		handleAddBlock(w, r)
	}
}

func handleError(err error, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Error: " + err.Error()))
	log.Print(err.Error())
}

func handleListBlocks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(blockchain.Blocks)
}

func handleAddBlock(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var receivedBlock core.Block
	err := decoder.Decode(&receivedBlock)
	if err == nil {
		log.Println(receivedBlock.Data, " received")
		go blockchain.AddBlock(receivedBlock)
	} else {
		handleError(err, w, r)
	}
}
