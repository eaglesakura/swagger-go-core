package utils

import (
	"github.com/eaglesakura/swagger-go-core"
	"net/http"
)

/*
バイナリデータ等、単純なデータを返却するためのレスポンスデフォルト実装を定義する
*/
type RawBufferResponse struct {
	StatusCode  int
	ContentType string
	Payload     []byte
}

/*
別なURL（Google Cloud Storage等）へのリダイレクトを行わせるレスポンスデフォルト実装を定義する
*/
type RedirectResponse struct {
	Location string
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

func (it *RedirectResponse) WriteResponse(w http.ResponseWriter, p swagger.Producer) {
	w.Header().Add("Location", it.Location)
	w.WriteHeader(301)
}
