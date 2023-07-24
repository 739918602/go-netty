package serialize

const (
	ERROR_DATA = iota
	DATA_READY
	DATA_NOT_READY
)

type StringMessage struct {
	header  string
	payload string
}

func (s StringMessage) Header() interface{} {
	return s.header
}

func (s StringMessage) Payload() []byte {
	return []byte(s.payload)
}
