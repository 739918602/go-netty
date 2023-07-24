package server

import (
	"go-netty/handler"
	"go-netty/serialize"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	boss      BossEventLoopGroup
	worker    WorkerEventLoopGroup
	chanClose chan bool
}

func New() *Server {
	return &Server{
		chanClose: make(chan bool),
	}
}

func (s *Server) BootStrap(bossSize, workerSize int) *Server {
	s.worker = WorkerEventLoopGroup{
		chanConn:   make(chan net.Conn),
		chanClose:  make(chan bool, 1),
		workerSize: workerSize,
		decoder:    serialize.TlvDecoder{},
		handler:    handler.InputHandler{},
	}
	s.boss = BossEventLoopGroup{
		size:     bossSize,
		chanExit: make(chan bool, 1),
		worker:   s.worker,
	}
	return s
}
func (s *Server) RegisterClose() error {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	<-c
	return s.Close()
}
func (s *Server) Close() error {
	log.Println("Close Begin")
	err := s.boss.Close()
	if err != nil {
		return err
	}
	s.worker.Close()
	s.chanClose <- true
	return nil
}
func (s *Server) ServeTCP(addr string) {
	go s.RegisterClose()
	s.boss.ServeTCP(addr)
	s.worker.Start()
	<-s.chanClose
}
