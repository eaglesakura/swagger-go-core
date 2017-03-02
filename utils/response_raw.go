package swagger

import (
	"net/http"
	"github.com/go-openapi/runtime"
)

type RawBufferResponse struct {
	StatusCode  int
	ContentType string
	Payload     []byte
}

func (it *RawBufferResponse) WriteResponse(w http.ResponseWriter, p runtime.Producer) {
	w.WriteHeader(it.StatusCode)
	if it.ContentType != "" {
		w.Header().Add("Content-Type", it.ContentType)
	}

	if it.Payload != nil {
		w.Write(it.Payload)
	}
}
