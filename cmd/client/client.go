package main

import (
	"go-netty/idl/gen/pb"
	"go-netty/serialize"
	"log"
	"net"
	"time"

	"google.golang.org/protobuf/proto"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		panic(err)
	}
	conn.(*net.TCPConn).SetKeepAlive(true)
	conn.SetDeadline(time.Now().Add(time.Second * 5))
	conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	conn.SetReadDeadline(time.Now().Add(time.Second * 5))
	body := pb.NamedStruct{
		Id:   int64(1),
		Name: "hello world",
	}
	protoBody, err := proto.Marshal(&body)
	if err != nil {
		panic(err)
	}
	msg := serialize.NewCustomMessage(1, protoBody)
	bt, err := msg.Pack()
	if err != nil {
		panic(err)
	}
	_, err = conn.Write(bt)
	if err != nil {
		panic(err)
	}
	log.Println("Send Msg:" + body.String())
	err = conn.Close()
	if err != nil {
		panic(err)
	}
	log.Println("Close Conn...")
}
