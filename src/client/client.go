package main


import (
	"fmt"
	"net"
	"socket"
	"sync"
	"bufio"
	"os"
)
// Changes these to match the server ip and port.
var server_ip string = "127.0.0.1"
var server_port string = "8080"

func initConnection() (net.Conn,error){
	
	ip := server_ip + ":"+ server_port
	
	// Connects to the chatserver using TCP.
	c , err := net.Dial("tcp", ip)
	if err != nil {
		fmt.Println("[Client] : Error : ",err)
	}else{
		fmt.Println("[Client] : Connected")
	}
	return c, err
}
func clientRead(conn net.Conn , nickname string) {
	for {
		// Reads any msg from chatserver.
		msg , err:= socket.ReadMsg(conn)
		if(err != nil) {
			fmt.Println("Lost connection with chatserver")
			break;
		}
		// Displays the message from chatserver to client.
		fmt.Print("[" + msg.Nickname + "] : " + msg.Message)
	}
}

func clientSend(conn net.Conn, wg *sync.WaitGroup, nickname string) {
	socket.SendMsg(conn, socket.TCP_Message{nickname,"Joined the chat.\n"})
	reader := bufio.NewReader(os.Stdin)
	// Loop until user writes "-1".
	for {

	  	var msg socket.TCP_Message
	  	var err error

	  	msg.Nickname = nickname
		msg.Message , err = reader.ReadString('\n') 
		if(err != nil) {
			fmt.Println("Stdin error ", err)
		}

		if(msg.Message[0] == '-' && msg.Message[1] == '1') {
			fmt.Println("Exiting chatsession.")
			socket.SendMsg(conn, socket.TCP_Message{nickname,"Leaved the chat."})
			break
		}
		// Sends msg to chatserver.
		socket.SendMsg(conn, msg)
	}
	wg.Done()
}
func setNickName() (string) {
	var nickname string
	fmt.Print("Enter your nickname : ")
	fmt.Scanln(&nickname)
	return nickname
}
func main() {

	var wg sync.WaitGroup
	// Inits the connection to the chatserver.
	conn, err := initConnection()
	fmt.Println("Welcome to the chatclient, exit chat with -1")
	if(err != nil) {
		fmt.Println("Error whith connection to host msg : ,",err)
	}else {
		wg.Add(1)
		nickname := setNickName()
		// Starts the goroutine for reading/sending from/to chatserver.
		go clientRead(conn,nickname)
		go clientSend(conn,&wg,nickname)
	}
	// Waiting for goroutine to be done.
	wg.Wait()
}