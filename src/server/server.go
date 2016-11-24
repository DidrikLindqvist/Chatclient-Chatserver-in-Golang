package main

import (
	"fmt"
	"net"
	"socket"
)
var clients []net.Conn

var server_ip string = "127.0.0.1"
var server_port string = "8080"

func deleteClosedConn(conn net.Conn) {
	// Finds the connection in clients and removes it.
	for i := 0 ; i <= len(clients) ; i++ {
		if(conn == clients[i]) {
			clients = clients[:i+copy(clients[i:], clients[i+1:])]
			break
		}
	} 
}

func echoToClients(conn net.Conn) {
	// Loop until client disconnect.
	for {	
		msg , err:= socket.ReadMsg(conn)
		//The connection got closed.
		if(err != nil) {
			fmt.Println("Lost connection to client")
			deleteClosedConn(conn)
			break
		}
		
		fmt.Println("Recived msg : ",  msg.Message , " from ", msg.Nickname)
		// Sends msg to all clients connected to server.
		for _, conn := range clients {				
			socket.SendMsg(conn,msg)
		}
	}
		
}

func initServer() (net.Listener){
	ip := server_ip + ":"+ server_port
	// Binds the ip using TCP protocol.
	l, err := net.Listen("tcp", ip)
	if err != nil {
		fmt.Println("[Server] : Error  : ",err)
	}
	fmt.Println("[Server] :Listen done")
	return l
}

func accpetNewClients(l net.Listener) {
	
	fmt.Println("[Server] : Listening to accept")
	// Waiting for new connection. If new it starts a goroutine for the client.
	for {
		conn, err := l.Accept()
		fmt.Println("[Server] : Accepted")
		if err != nil {
			fmt.Println("[Server] : Error",err)
		}else{
			clients = append(clients, conn)
			go echoToClients(conn)
		}
	}
}

func main() {
	// Inits the server, binding ip/port with tcp.
	listener := initServer()
	// Will take care of accepting new clients.
	go accpetNewClients(listener)

	var input string
	fmt.Println("Exit server by pressing enter in console.")
    fmt.Scanln(&input)
    fmt.Println("Exiting server")
}