package server

import (
	"errors"
	"go-netty/handler"
	"go-netty/serialize"
	"io"
	"net"
)

type Server struct {
	conn    net.Conn
	handler handler.Handler
	decoder serialize.Decoder
}

func New() *Server {
	return &Server{}
}
func (s *Server) Handler(handler handler.Handler) *Server {
	s.handler = handler
	return s
}

func (s *Server) Decoder(decoder serialize.Decoder) *Server {
	s.decoder = decoder
	return s
}

func (s *Server) ServeTCP(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	for {
		s.accept(l)
	}

	return nil
}

func (s *Server) accept(l net.Listener) {
	defer func() {
		if recErr := recover(); recErr != nil {
			s.handler.OnError(recErr.(error))
			s.conn.Close()
		}
	}()
	var err error
	s.conn, err = l.Accept()
	if err != nil {
		s.handler.OnError(err)
		return
	}
	//Connection Ready
	err = s.handler.OnAccept(s.conn)
	if err != nil {
		s.handler.OnError(err)
		return
	}
	go func(conn net.Conn) {
		for {
			msg, err := s.decoder.Decode(conn)
			if errors.Is(err, io.EOF) {
				s.handler.OnDisConnect(conn)
				break
			}
			if err != nil {
				s.handler.OnError(err)
				break
			}
			s.handler.OnRead(msg, conn)
		}
		if conn != nil {
			err = conn.Close()
		}
		if err != nil {
			s.handler.OnError(err)
		}
		err = s.handler.OnClose()
		if err != nil {
			s.handler.OnError(err)
		}
	}(s.conn)
}
