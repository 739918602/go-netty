package test

import (
	"go-netty/idl/gen/pb"
	"go-netty/serialize"
	"log"
	"net"
	"testing"
	"time"

	"google.golang.org/protobuf/proto"
)

func BenchmarkClient(b *testing.B) {

	b.RunParallel(func(pbt *testing.PB) {
		conn, err := net.Dial("tcp", "192.168.0.78:9000")
		if err != nil {
			panic(err)
		}
		var count int
		for pbt.Next() {
			count++
			conn.SetDeadline(time.Now().Add(time.Second * 5))
			body := pb.NamedStruct{
				Id:   int64(count),
				Name: "Hello World",
			}
			protoBody, err := proto.Marshal(&body)
			if err != nil {
				panic(err)
			}
			msg := serialize.NewCustomMessage(1, protoBody)
			log.Println("Send Msg:" + body.String())
			if err != nil {
				panic(err)
			}
			bt, err := msg.Pack()
			if err != nil {
				panic(err)
			}
			_, err = conn.Write(bt)
			if err != nil {
				panic(err)
			}
		}
		err = conn.Close()
		if err != nil {
			panic(err)
		}
		b.Log("Close Conn...")
	})
}
