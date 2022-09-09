package main

import (
	"go-netty/handler"
	"go-netty/serialize"
	"go-netty/server"
)

func main() {
	server.New().
		Handler(handler.InputHandler{}).
		Decoder(serialize.LineDecoder{Limit: 10}).
		ServeTCP(":9000")
}
