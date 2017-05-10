package swagger

import (
	"io"
)

type DataPayload interface {
	/**
	 * content-length
	 * ex) unknown < 0
	 */
	GetContentLength() int

	/**
	 * content-type
	 */
	GetContentType() string

	/**
	 * data stream
	 */
	OpenStream() io.ReadCloser
}

type FetchClient interface {
	ValidatorFactory

	/**
	 * 文字列をURL Encodeする
	 */
	EscapeString(origin string) string

	/**
	 * エンドポイントのPATHを指定する
	 */
	SetApiPath(path string)

	/**
	 * httpメソッドを指定する
	 */
	SetMethod(method string)

	/**
	 * @param key   query key(not url encoded!)
	 * @param value query value
	 */
	AddQueryParam(key string, value string)

	/**
	 * @param key   query key(not url encoded!)
	 * @param value query value
	 */
	AddHeader(key string, value string)

	SetPayload(payload DataPayload)

	/**
	 * URL Fetch & decode data
	 */
	Fetch(resultPtr interface{}) error
}
