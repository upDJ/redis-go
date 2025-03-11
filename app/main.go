package main

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/connect"
	"github.com/codecrafters-io/redis-starter-go/app/config"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")
	config.GetConfig()
  	l := connect.InitTcp()
	connect.HandleConnection(l)
}
