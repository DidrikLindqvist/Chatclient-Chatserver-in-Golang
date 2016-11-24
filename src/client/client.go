package main


import (
	"fmt"
	"net"
	"socket"
	"sync"
	"bufio"
	"os"
)

var server_ip string = "127.0.0.1"
var server_port string = "8080"

func initConnection(host string) (net.Conn,error){
	c , err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println("[Client] : Error : ",err)
	}else{
		fmt.Println("[Client] : Connected")
	}
	return c, err
}
func clientRead(conn net.Conn , nickname string) {
	for {
		msg , err:= socket.ReadMsg(conn)
		if(err != nil) {
			fmt.Println("Lost connection with chatserver")
			break;
		}
		fmt.Print("[" + msg.Nickname + "] : " + msg.Message)
		//fmt.Print("[You] : ")
	}
}

func clientSend(conn net.Conn, wg *sync.WaitGroup, nickname string) {
	socket.SendMsg(conn, socket.TCP_Message{nickname,"Joined the chat.\n"})
	reader := bufio.NewReader(os.Stdin)
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
		socket.SendMsg(conn, msg)
		//socket.WriteToSocket(conn,nickname+input)
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
	ip := server_ip + ":"+ server_port
	conn, err := initConnection(ip)
	fmt.Println("Welcome to the chatclient, exit chat with -1")
	if(err != nil) {
		fmt.Println("Error whith connection to host msg : ,",err)
	}else {
		wg.Add(1)
		//temporary
		nickname := setNickName()
		go clientRead(conn,nickname)
		go clientSend(conn,&wg,nickname)
	}
	
	wg.Wait()
}