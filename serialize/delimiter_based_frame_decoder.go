package serialize

import (
	"bytes"
	"errors"
	"io"
	"net"
	"strings"
)

type DelimiterBasedFrameDecoder struct {
	Limit     int
	Delimiter string
}

func (l DelimiterBasedFrameDecoder) Decode(conn net.Conn) (Message, error) {
	temp := make([]byte, 1)
	buf := bytes.NewBufferString("")
	var size int
	for {
		size++
		if size > l.Limit {
			return nil, errors.New("size exceeds the limit")
		}

		n, err := conn.Read(temp)
		if errors.Is(err, io.EOF) {
			return StringMessage{
				header:  "",
				payload: strings.TrimSuffix(buf.String(), l.Delimiter),
			}, err
		}
		if err != nil {
			return nil, err
		}
		size += n
		_, err = buf.Write(temp)
		if err != nil {
			buf.Reset()
			return nil, err
		}
		if n == 0 {
			break
		}
		if strings.HasSuffix(buf.String(), l.Delimiter) {
			break
		}
	}
	return StringMessage{
		header:  "",
		payload: strings.TrimSuffix(buf.String(), l.Delimiter),
	}, nil
}
