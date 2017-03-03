package swagger

import (
	"testing"
	"github.com/go-openapi/swag"
	"github.com/stretchr/testify/assert"
	"github.com/eaglesakura/swagger-go-core/utils"
)

func Test_ValidatorImpl_Valid(t *testing.T) {
	value := swag.String("value")
	var nilValue *string
	var nilInterface interface{} = nilValue

	assert.Nil(t, nilValue)
	assert.Nil(t, nilInterface)

	assert.True(t, swagger.NewValidator(value, value == nil).Required(true).Valid(swagger.NewValidatorFactory()))
	assert.False(t, swagger.NewValidator(nilValue, nilValue == nil).Required(true).Valid(swagger.NewValidatorFactory()))
	assert.False(t, swagger.NewValidator(nil, true).Required(true).Valid(swagger.NewValidatorFactory()))
}