/*
port github.com/go-openapi/runtime/

LICENSE https://github.com/go-openapi/runtime/blob/master/LICENSE

go-openapiは内部でunsafeを使用しているため、unsafeを取り除いたコンパクト実装を定義する
*/
package swagger

import (
	"net/http"
	"io"
)

// ProducerFunc represents a function that can be used as a producer
type ProducerFunc func(io.Writer, interface{}) error

func (it ProducerFunc) Produce(w io.Writer, v interface{}) error {
	return it(w, v)
}

// Producer implementations know how to turn the provided interface into a valid
// HTTP response
type Producer interface {
	// Produce writes to the http response
	Produce(io.Writer, interface{}) error
}

// ConsumerFunc represents a function that can be used as a consumer
type ConsumerFunc func(io.Reader, interface{}) error

// Consumer implementations know how to bind the values on the provided interface to
// data provided by the request body
type Consumer interface {
	// Consume performs the binding of request values
	Consume(io.Reader, interface{}) error
}

func (it ConsumerFunc) Consume(r io.Reader, v interface{}) error {
	return it(r, v)
}


// Responder is an interface for types to implement
// when they want to be considered for writing HTTP responses
type Responder interface {
	WriteResponse(http.ResponseWriter, Producer)
}
