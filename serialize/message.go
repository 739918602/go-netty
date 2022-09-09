package serialize

type Message interface {
	Header() interface{}
	Payload() interface{}
}
