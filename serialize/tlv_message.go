package serialize

import (
	"bytes"
	"encoding/binary"
	"errors"
	"hash/crc32"
)

const (
	HeanderLen         = 20
	MagicNumber uint16 = 0x3C2B
)

type TlvHeader struct {
	MagicNumber uint16
	Version     uint16
	CRC         uint32
	Len         uint32
	Type        uint32
	Ext         uint32
}

type TlvMessage struct {
	TlvHeader
	Body []byte
}

func NewCustomMessage(tp uint32, body []byte) *TlvMessage {
	return &TlvMessage{
		TlvHeader: TlvHeader{
			MagicNumber: MagicNumber,
			Version:     1,
			Type:        tp,
		},
		Body: body,
	}
}

func (s *TlvMessage) Header() any {
	return s.TlvHeader
}

func (s *TlvMessage) Payload() []byte {
	return s.Body
}

func (s *TlvMessage) UnpackHeader(header []byte) error {
	buf := bytes.NewBuffer(header)
	binary.Read(buf, binary.LittleEndian, &s.MagicNumber)
	binary.Read(buf, binary.LittleEndian, &s.Version)
	binary.Read(buf, binary.LittleEndian, &s.CRC)
	binary.Read(buf, binary.LittleEndian, &s.Len)
	binary.Read(buf, binary.LittleEndian, &s.Type)
	binary.Read(buf, binary.LittleEndian, &s.Ext)
	if s.MagicNumber != MagicNumber {
		return errors.New("magic number error")
	}
	return nil
}
func (s *TlvMessage) UnpackPlayload(payload []byte) error {
	crc := crc32.ChecksumIEEE(payload)
	if s.CRC != crc {
		return errors.New("crc error")
	}
	s.Body = payload
	return nil
}

func (s *TlvMessage) Pack() ([]byte, error) {
	var err error
	buf := bytes.NewBuffer(nil)
	err = binary.Write(buf, binary.LittleEndian, MagicNumber)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, s.Version)
	if err != nil {
		return nil, err

	}
	err = binary.Write(buf, binary.LittleEndian, crc32.ChecksumIEEE(s.Body))
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, uint32(len(s.Body)))
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, s.Type)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, s.Ext)
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.LittleEndian, s.Body)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
