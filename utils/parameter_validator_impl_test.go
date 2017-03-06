package swagger

import (
	"testing"
	"github.com/go-openapi/swag"
	"github.com/stretchr/testify/assert"
)

func Test_ValidatorImpl_Valid(t *testing.T) {
	value := swag.String("value")
	var nilValue *string
	var nilInterface interface{} = nilValue

	assert.Nil(t, nilValue)
	assert.Nil(t, nilInterface)

	assert.True(t, NewValidator(value, value == nil).Required(true).Valid(NewValidatorFactory()))
	assert.False(t, NewValidator(nilValue, nilValue == nil).Required(true).Valid(NewValidatorFactory()))
	assert.False(t, NewValidator(nil, true).Required(true).Valid(NewValidatorFactory()))
}

func Test_ValidatorImpl_Enum(t *testing.T) {
	assert.True(t, NewValidator(swag.String("a"), false).EnumPattern([]string{"this", "is", "a", "pen"}).Valid(NewValidatorFactory()))
	assert.False(t, NewValidator(swag.String("A"), false).EnumPattern([]string{"this", "is", "a", "pen"}).Valid(NewValidatorFactory()))
	assert.False(t, NewValidator(swag.String(""), false).EnumPattern([]string{"this", "is", "a", "pen"}).Valid(NewValidatorFactory()))
	assert.False(t, NewValidator(swag.String("swagger"), false).EnumPattern([]string{"this", "is", "a", "pen"}).Valid(NewValidatorFactory()))
	assert.True(t, NewValidator(nil, true).EnumPattern([]string{"this", "is", "a", "pen"}).Valid(NewValidatorFactory()))
}