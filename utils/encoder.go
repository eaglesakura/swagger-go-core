package utils

import (
	"io"
	"encoding/json"
	"github.com/eaglesakura/swagger-go-core"
)

/*
swaggerで生成された構造体からJson形式へ変換するためのデフォルト実装を提供する
*/
func NewJsonProducer() swagger.Producer {
	return swagger.ProducerFunc(func(writer io.Writer, data interface{}) error {
		enc := json.NewEncoder(writer)
		return enc.Encode(data)
	})
}

/*
Json形式からswaggerで生成された構造体へ変換するためのデフォルト実装を提供する
*/
func NewJsonConsumer() swagger.Consumer {
	return swagger.ConsumerFunc(func(reader io.Reader, data interface{}) error {
		dec := json.NewDecoder(reader)
		dec.UseNumber() // preserve number formats
		return dec.Decode(data)
	})
}
