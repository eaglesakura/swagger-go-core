package utils

import (
	"net/http"
	"github.com/eaglesakura/swagger-go-core"
)

type RawBufferResponse struct {
	StatusCode  int
	ContentType string
	Payload     []byte
}

func (it *RawBufferResponse) WriteResponse(w http.ResponseWriter, p swagger.Producer) {
	w.WriteHeader(it.StatusCode)
	if it.ContentType != "" {
		w.Header().Add("Content-Type", it.ContentType)
	}

	if it.Payload != nil {
		w.Write(it.Payload)
	}
}
