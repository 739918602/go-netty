package serialize

import (
	"encoding/binary"
	"io"
	"net"
)

type TlvDecoder struct {
}

func (l TlvDecoder) Decode(conn net.Conn) (Message, error) {
	message := &TlvMessage{}
	var err error
	header := make([]byte, HeanderLen)
	_, err = io.ReadFull(conn, header)
	if err != nil {
		return nil, err
	}
	// 解出头部
	err = message.UnpackHeader(header)
	if err != nil {
		return nil, err
	}
	// 解出包体内容
	payload := make([]byte, message.Len)
	err = binary.Read(conn, binary.LittleEndian, payload)
	if err != nil {
		return nil, err
	}
	err = message.UnpackPlayload(payload)
	if err != nil {
		return nil, err
	}
	return message, nil
}
