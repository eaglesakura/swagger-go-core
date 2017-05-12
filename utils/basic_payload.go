package utils

import (
	"io"
	"bytes"
	"github.com/eaglesakura/swagger-go-core"
	"net/url"
	"encoding/json"
)

type BasicPayload struct {
	Buffer      []byte
	ContentType string
}

type readerImpl struct {
	reader io.Reader
}

func (it *BasicPayload)GetContentLength() int {
	return len(it.Buffer)
}

func (it *BasicPayload)GetContentType() string {
	return it.ContentType
}

func (it *BasicPayload)OpenStream() io.ReadCloser {
	reader := bytes.NewReader(it.Buffer)
	return &readerImpl{
		reader:reader,
	}
}

func (it *readerImpl)Read(p []byte) (n int, err error) {
	return it.reader.Read(p)
}

func (it *readerImpl)Close() error {
	return nil
}

func NewFormPayload(values url.Values) swagger.DataPayload {
	result := BasicPayload{
		Buffer:[]byte(values.Encode()),
		ContentType:"application/x-www-form-urlencoded",
	}
	return &result
}

func NewJsonPayload(values interface{}) swagger.DataPayload {
	buf, _ := json.Marshal(values)
	result := BasicPayload{
		Buffer:buf,
		ContentType:"application/json",
	}
	return &result
}