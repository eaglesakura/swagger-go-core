/*
port github.com/go-openapi/runtime/

LICENSE https://github.com/go-openapi/runtime/blob/master/LICENSE

go-openapiは内部でunsafeを使用しているため、unsafeを取り除いたコンパクト実装を定義する
*/
package errors

import "fmt"

/*
httpステータスコードと表示メッセージを定義したエラー
*/
type apiError struct {
	code    int32
	message string
}

func (a *apiError) Error() string {
	return a.message
}

func (a *apiError) Code() int32 {
	return a.code
}

// New creates a new API error with a code and a message
func New(code int32, message string, args ...interface{}) error {
	if len(args) > 0 {
		return &apiError{code, fmt.Sprintf(message, args...)}
	}
	return &apiError{code, message}
}

/*
panic()からrecoverされた際に目印としてラップされる
*/
type PanicError struct {
	Origin error
}

func (a *PanicError) Error() string {
	return a.Origin.Error()
}
