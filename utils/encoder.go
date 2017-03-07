package swagger

import (
	"io"
	"encoding/json"
	"github.com/eaglesakura/swagger-go-core"
)

func NewJsonProducer() swagger.Producer {
	return swagger.ProducerFunc(func(writer io.Writer, data interface{}) error {
		enc := json.NewEncoder(writer)
		return enc.Encode(data)
	})
}

func NewJsonConsumer() swagger.Consumer {
	return swagger.ConsumerFunc(func(reader io.Reader, data interface{}) error {
		dec := json.NewDecoder(reader)
		dec.UseNumber() // preserve number formats
		return dec.Decode(data)
	})
}
