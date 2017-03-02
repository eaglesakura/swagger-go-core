package swagger

import (
	"github.com/eaglesakura/swagger-go-core"
	swagger_internal "github.com/eaglesakura/swagger-go-core/internal"
	"net/http"
	"github.com/go-openapi/runtime"
)

/**
 * デフォルトのValidatorを生成する
 *
 * nil値を直接interface{}に変換した場合、==nilチェックが行えない仕様上の問題がある。
 * 現時点ではworkaroundとして、isNilをコード生成で賄う。
 */
func NewValidator(value interface{}, isNil bool) swagger.ParameterValidator {
	return &swagger_internal.ValidatorImpl{Value:value, IsNil:isNil}
}

/**
 * デフォルトのMapperを生成する
 */
func NewHandlerMapper() swagger.HandleMapper {
	return &swagger_internal.HandleMapperImpl{
		Mappers:map[string]*swagger_internal.MethodMapper{},
	}
}

/**
 * デフォルトのRequestBinderを生成する
 *
 * Consumerの取得はFunctionに任せられる。
 */
func NewRequestBinder(req *http.Request, consumerFactory func(contentType string) runtime.Consumer) swagger.RequestBinder {
	return &swagger_internal.BasicRequestBinder{
		Request:req,
		ConsumerFactory:consumerFactory,
	}
}