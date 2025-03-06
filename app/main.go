package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this block to pass the first stage
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
	  fmt.Println("Failed to bind to port 6379")
    os.Exit(1)
	}

  var conn net.Conn
	conn, err = l.Accept()
	if err != nil {
	  fmt.Println("Error accepting connection: ", err.Error())
    os.Exit(1)
	}
 
  for {
    buf := make([]byte, 1024)

    _, err := conn.Read(buf)
    if err!= nil {
      fmt.Println("Error reading:", err)
      return
    }

    conn.Write([]byte("+PONG\r\n"))
  }
}
