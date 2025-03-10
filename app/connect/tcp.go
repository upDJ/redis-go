package connect

import (
	"github.com/codecrafters-io/redis-starter-go/app/utils"
	"fmt"
	"net"
	"os"
)

func InitTcp() (net.Listener) {
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

func processRequest(conn net.Conn) {
  for {
    buf := make([]byte, 1024)

    _, err := conn.Read(buf)
    if err!= nil {
      fmt.Println("Error reading:", err)
      return
    }

    res := utils.InputParser(string(buf))
    conn.Write([]byte(res))
  }
}

func HandleConnection(l net.Listener) {
  for {
    conn := getClientConnection(l)
    go processRequest(conn)
  }
}


