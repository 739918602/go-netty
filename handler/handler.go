package handler

import (
	"go-netty/serialize"
	"net"
)

type Handler interface {
	OnAccept(conn net.Conn) error
	OnRead(msg serialize.Message, conn net.Conn)
	OnDisConnect(conn net.Conn)
	OnClose(conn net.Conn) error
	OnError(err error)
}
