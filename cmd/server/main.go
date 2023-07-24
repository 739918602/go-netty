package main

import (
	"go-netty/server"
)

func main() {
	server.New().BootStrap(2, 8).ServeTCP(":9000")
}
