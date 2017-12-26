package utils

import (
	"github.com/eaglesakura/swagger-go-core"
	swagger_internal "github.com/eaglesakura/swagger-go-core/internal"
	"net/http"
)

/*
デフォルトのValidatorを生成する.
*/
func NewValidator(value interface{}, isNil bool) swagger.ParameterValidator {
	return swagger_internal.NewValidator(value, isNil)
}

/*
デフォルトのValidatorFactoryを生成する
*/
func NewValidatorFactory() swagger.ValidatorFactory {
	return swagger_internal.NewValidatorFactory()
}

/*
デフォルトのMapperを生成する
*/
func NewHandleMapper() swagger.HandleMapper {
	return swagger_internal.NewHandleMapper()
}

/*
デフォルトのRequestBinderを生成する.

Consumerの取得はFunctionに任せられる.
*/
func NewRequestBinder(req *http.Request, consumerFactory func(contentType string) swagger.Consumer) swagger.RequestBinder {
	return swagger_internal.NewRequestBinder(req, consumerFactory)
}
