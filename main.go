package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/garritfra/blockchain-project/server"
)

func main() {
	portFlag := flag.Int("port", 42000, "The port this service is running on")
	flag.Parse()

	var port string
	port += ":"
	port += strconv.Itoa(*portFlag)
	fmt.Println(*portFlag)
	server.ServeHTTP(port)
}
