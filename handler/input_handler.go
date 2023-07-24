package handler

import (
	"go-netty/idl/gen/pb"
	"go-netty/serialize"
	"log"
	"net"

	"google.golang.org/protobuf/proto"
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
	msg.Payload()
	protoBody := pb.NamedStruct{}
	proto.Unmarshal(msg.Payload(), &protoBody)
	log.Println(conn.RemoteAddr().String(), protoBody.String())
	//conn.Write([]byte("echo:" + string(msg.Payload()) + "\n"))
}

func (InputHandler) OnClose(conn net.Conn) error {
	log.Println("Server Close")
	return conn.Close()
}

func (InputHandler) OnError(err error) {
	log.Println("Error:", err.Error())
}
