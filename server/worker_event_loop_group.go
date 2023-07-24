package server

import (
	"errors"
	"go-netty/handler"
	"go-netty/serialize"
	"io"
	"log"
	"net"
)

type WorkerEventLoopGroup struct {
	chanConn   chan net.Conn
	chanClose  chan bool
	workerSize int
	decoder    serialize.Decoder
	handler    handler.Handler
}

func NewWorkerEventLoopGroup(handler handler.Handler, decoder serialize.Decoder, workerSize int) *WorkerEventLoopGroup {
	return &WorkerEventLoopGroup{
		chanConn:   make(chan net.Conn, 1024),
		chanClose:  make(chan bool),
		workerSize: workerSize,
		decoder:    decoder,
		handler:    handler,
	}
}
func (w WorkerEventLoopGroup) AcceptConnection(conn net.Conn) {
	w.chanConn <- conn
}

func (w WorkerEventLoopGroup) Start() {
	//启动worker
	for i := 0; i < w.workerSize; i++ {
		go func() {
			for {
				conn := <-w.chanConn //阻塞等待新连接
				//处理新来的连接
				w.handle(conn)
			}
		}()
	}
}

// 处理连接生命周期
func (w WorkerEventLoopGroup) handle(conn net.Conn) {
	w.handler.OnAccept(conn)
	for {
		select {
		case <-w.chanClose:
			w.handler.OnClose(conn)
			return
		default:
			msg, err := w.decoder.Decode(conn)
			if err != nil {
				if errors.Is(err, io.EOF) {
					w.handler.OnDisConnect(conn)
					return
				}
				w.handler.OnError(err)
				return
			}
			w.handler.OnRead(msg, conn)
		}
	}
}

func (w WorkerEventLoopGroup) Close() {
	w.chanClose <- true
	log.Println("WorkerEventLoopGroup exit")
}
