package main

import (
	"fmt"
	"net"
	"socket"
)
var clients []net.Conn

func deleteClosedConn(conn net.Conn) {
	for i := 0 ; i <= len(clients) ; i++ {
		if(conn == clients[i]) {
			clients = clients[:i+copy(clients[i:], clients[i+1:])]
			break
		}
	} 
}
func echoToClients(conn net.Conn) {

	for {	
		msg , err:= socket.ReadMsg(conn)
		if(err != nil) {
			fmt.Println("Lost connection to client")
			deleteClosedConn(conn)
			break
		}
		if(msg.Nickname != "" ) {
			fmt.Println("Recived msg : ",  msg.Message , " from ", msg.Nickname)
			for _, conn := range clients {				
				socket.SendMsg(conn,msg)
			}
		}		
	}
}
func initServer() (net.Listener){
	// Listen on TCP port 2000 on all interfaces.
	l, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("[Server] : Error  : ",err)
	}
	fmt.Println("[Server] :Listen done")
	return l
}
func accpetNewClients(l net.Listener) {
	// Wait for a connection.
	fmt.Println("[Server] : Listening to accept")

	for {
		conn, err := l.Accept()
		fmt.Println("[Server] : Accepted")
		if err != nil {
			fmt.Println("[Server] : Error",err)
		}
		clients = append(clients, conn)
		go echoToClients(conn)

	}
}

func main() {
	listener := initServer()
	go accpetNewClients(listener)

	var input string
	fmt.Println("Exit server by pressing enter in console.")
    fmt.Scanln(&input)
    fmt.Println("Exiting server")
}