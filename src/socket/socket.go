package socket

import (
	"net"
	"encoding/gob"
)

type TCP_Message struct {
	Nickname string
	Message string
}

func SendMsg(conn net.Conn, msg_s TCP_Message ) (error){
	enc := gob.NewEncoder(conn)
	err := enc.Encode(msg_s)
	return err
}
func ReadMsg(conn net.Conn) (TCP_Message , error) {
	dec := gob.NewDecoder(conn)
	
	var msg TCP_Message
	err := dec.Decode(&msg)

	return msg, err
}