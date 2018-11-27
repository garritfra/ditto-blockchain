package main

import (
	"github.com/garritfra/blockchain-project/server"
)

func main() {
	server.ServeHTTP(":42000")
}
