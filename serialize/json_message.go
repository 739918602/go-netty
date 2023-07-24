package serialize

import "encoding/json"

type JsonMessage struct {
	header  []byte
	payload []byte
}

func (s JsonMessage) Header() interface{} {
	return string(s.header)
}

func (s JsonMessage) Payload() []byte {
	return s.payload
}
func (s JsonMessage) UnmarshalPayload(target interface{}) error {
	err := json.Unmarshal(s.payload, target)
	if err != nil {
		return err
	}
	return nil
}

func (s JsonMessage) UnmarshalHeader(target interface{}) error {
	err := json.Unmarshal(s.header, target)
	if err != nil {
		return err
	}
	return nil
}
