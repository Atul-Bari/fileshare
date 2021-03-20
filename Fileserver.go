// Very basic socket server
// https://golangr.com/

package main

import ( 	
	"net"
	"fmt"
	"bufio"
)

var userData = make(map[string]*user)

func main() {
  fmt.Println("Start server...")

  // listen on port 8000
  ln, _ := net.Listen("tcp", ":8000")

  // accept connection
  conn, _ := ln.Accept()
  fmt.Printf("Data %V", conn)

  // run loop forever (or until ctrl-c)
  for {
    // get message, output
    message, _ := bufio.NewReader(conn).ReadString('\n')
    fmt.Print("Message Received:", string(message))

  }
}