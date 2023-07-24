package serialize

type Message interface {
	Header() interface{}
	Payload() []byte
}

type TLVMessage interface {
	GetMagicNumber() uint16
	GetVersion() uint16
	GetType() uint32
	GetExt() uint32
	Len() uint32
	GetPayload() []byte
}
