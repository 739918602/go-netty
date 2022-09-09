package handler

import (
	"go-netty/serialize"
	"log"
	"net"
)

type InputHandler struct{}

func (InputHandler) OnAccept(conn net.Conn) error {
	log.Println(conn.RemoteAddr().String() + " OnAccept")
	return nil
}
func (InputHandler) OnDisConnect(conn net.Conn) {
	log.Println(conn.RemoteAddr().String() + " OnDisConnect")
}
func (InputHandler) OnRead(msg serialize.Message, conn net.Conn) {
	log.Println(conn.RemoteAddr().String(), msg.Payload())
	conn.Write([]byte("echo:" + msg.Payload().(string) + "\n"))
}

func (InputHandler) OnClose() error {
	log.Println("Server Close")
	return nil
}

func (InputHandler) OnError(err error) {
	log.Println("Error", err.Error())
}
