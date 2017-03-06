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

func Test_ValidatorImpl_Enum(t *testing.T) {
	assert.True(t, swagger.NewValidator(swag.String("a"), false).EnumPattern([]string{"this", "is", "a", "pen"}).Valid(swagger.NewValidatorFactory()))
	assert.False(t, swagger.NewValidator(swag.String("A"), false).EnumPattern([]string{"this", "is", "a", "pen"}).Valid(swagger.NewValidatorFactory()))
	assert.False(t, swagger.NewValidator(swag.String(""), false).EnumPattern([]string{"this", "is", "a", "pen"}).Valid(swagger.NewValidatorFactory()))
	assert.False(t, swagger.NewValidator(swag.String("swagger"), false).EnumPattern([]string{"this", "is", "a", "pen"}).Valid(swagger.NewValidatorFactory()))
	assert.True(t, swagger.NewValidator(nil, true).EnumPattern([]string{"this", "is", "a", "pen"}).Valid(swagger.NewValidatorFactory()))
}