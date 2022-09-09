package serialize

import (
	"bytes"
	"errors"
	"net"
)

type Decoder interface {
	Decode(conn net.Conn) (Message, error)
}
type LineDecoder struct {
	Limit int
}

func (l LineDecoder) Decode(conn net.Conn) (Message, error) {
	temp := make([]byte, 1)
	buf := bytes.NewBuffer(nil)
	var count int
	for {
		count++
		if count > l.Limit {
			return nil, errors.New("size exceeds the limit")
		}
		n, err := conn.Read(temp)
		if err != nil {
			return nil, err
		}
		r := rune(temp[0])
		if r == '\r' || r == '\n' {
			break
		}
		_, err = buf.Write(temp)
		if err != nil {
			buf.Reset()
			return nil, err
		}
		if n == 0 {
			break
		}
	}
	return StringMessage{
		header:  "",
		payload: buf.String(),
	}, nil
}
