package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
)

// Transaction struct
type Transaction struct {
	Sender   string
	Receiver string
	Amount   int
}

func main() {
	blockchain := newBlockchain()

	StartServer()

	data, _ := json.Marshal(Transaction{Sender: "foo", Receiver: "bar", Amount: 100})
	block := newBlock(blockchain.blocks[0].Hash, data)
	blockchain.addBlock(block)

	data2, _ := json.Marshal(Transaction{Sender: "bar", Receiver: "baz", Amount: 5000})
	block = newBlock(blockchain.blocks[1].Hash, data2)
	blockchain.addBlock(block)

	for i := 0; i < len(blockchain.blocks); i++ {

		block := blockchain.blocks[i]
		fmt.Println("Block " + string(i))
		fmt.Println(hex.EncodeToString(block.PreviousHash))
		fmt.Println(hex.EncodeToString(block.Hash))

		data := string(block.Data)
		fmt.Println("Data: " + data)
	}

}
