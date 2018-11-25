package main

import (
	"fmt"
)

func main() {
	blockchain := newBlockchain()

	for _, v := range blockchain.blocks {
		fmt.Println(v.hash)
	}

}
