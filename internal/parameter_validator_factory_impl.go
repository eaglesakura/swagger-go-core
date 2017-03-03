package swagger

import "github.com/eaglesakura/swagger-go-core"

type ParameterValidatorFactoryImpl struct {

}

func (it *ParameterValidatorFactoryImpl)NewValidator(value interface{}, isNil bool) swagger.ParameterValidator {
	return &ParameterValidatorImpl{
		Value:value,
		IsNil:isNil,
	}
}
