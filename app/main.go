package main 
import ( "fmt"
	"net"
	"os"
  "strings"
)


var inputMap map[string][]string

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

//func response()
func inputParser(data string) (string) {
  var ping = string("ping")
  var echo = string("echo")
  var set = string("set")
  var get = string("get")
  var delim = string("\r\n")
  var dataArr = strings.Split(data, delim)
  
  if strings.Contains(strings.ToLower(data), ping) {
      return string("+PONG\r\n")

  } else if strings.Contains(strings.ToLower(data), echo) {
      return string("+" + dataArr[4] + delim)
  
  } else if strings.Contains(strings.ToLower(data), set) {
      key := dataArr[4] 
      val := dataArr[5:7]
      inputMap[key] = val
      return string("+OK" + delim)

  } else if strings.Contains(strings.ToLower(data), get) {
      val := inputMap[dataArr[4]] 
      return strings.Join(val, delim) + delim

  }

  return string("")
}

func connectionEcho(conn net.Conn) {
  for {
    buf := make([]byte, 1024)

    _, err := conn.Read(buf)
    if err!= nil {
      fmt.Println("Error reading:", err)
      return
    }

    res := inputParser(string(buf))
    conn.Write([]byte(res))
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
  inputMap = make(map[string][]string)

  l := initTcp()
  handleConnection(l)
}
