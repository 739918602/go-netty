package serialize

import (
	"net"
)

type Decoder interface {
	Decode(conn net.Conn) (Message, error)
}
