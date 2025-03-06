package main

import (
	"fmt"
	"net"
	"os"
)

func initTcp() (net.Listener) {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
	  fmt.Println("Failed to bind to port 6379")
    os.Exit(1)
	}
  
  return l
}

func getClientConnection(l net.Listener) (net.Conn) {
  conn, err := l.Accept()
	if err != nil {
	  fmt.Println("Error accepting connection: ", err.Error())
    os.Exit(1)
	}

  return conn
}

func connectionEcho(conn net.Conn) {
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

func handleConnection(l net.Listener) {
  for {
    conn := getClientConnection(l)
    go connectionEcho(conn)
  }
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")
  
  l := initTcp()
  handleConnection(l)

 }
