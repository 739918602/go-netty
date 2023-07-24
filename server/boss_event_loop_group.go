package server

import (
	"log"
	"net"
)

type BossEventLoopGroup struct {
	l        net.Listener
	size     int
	chanExit chan bool
	worker   WorkerEventLoopGroup
}

func (b *BossEventLoopGroup) ServeTCP(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	b.l = l
	b.Start()
	return nil
}

func (b *BossEventLoopGroup) Close() error {
	b.chanExit <- true
	return b.l.Close()
}

func (b *BossEventLoopGroup) Start() {
	for i := 0; i < b.size; i++ {
		go func(l net.Listener) {
			for {
				select {
				case <-b.chanExit:
					log.Println("BossEventLoopGroup exit")
					b.worker.Close()
					return
				default:
					conn, err := l.Accept()
					if err != nil {
						log.Println(err)
						return
					}
					b.worker.AcceptConnection(conn)
				}
			}
		}(b.l)
	}
}
