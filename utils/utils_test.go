package utils

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddPath(t *testing.T) {
	assert.Equal(t, AddPath("test/", "path"), "test/path")
	assert.Equal(t, AddPath("test", "path"), "test/path")
	assert.Equal(t, AddPath("test", "/path"), "test/path")
	assert.Equal(t, AddPath("test/", "/path"), "test/path")
}

type testStruct struct {
	IntValue    int
	StringValue string
}

func (it testStruct) String() string {
	buf, _ := json.Marshal(it)
	return string(buf)
}

func TestParameterToString_nil(t *testing.T) {
	var temp *int

	assert.Equal(t, ParameterToString(temp), "")
}

func TestParameterToString_nil_struct(t *testing.T) {
	var temp *testStruct

	assert.Equal(t, ParameterToString(temp), "")
}

func TestParameterToString_primitive(t *testing.T) {
	var temp int = 100

	assert.Equal(t, ParameterToString(&temp), "100")
}

func TestParameterToString_struct(t *testing.T) {
	temp := testStruct{
		StringValue: "test",
		IntValue:    100,
	}

	assert.Equal(t, ParameterToString(&temp), `{"IntValue":100,"StringValue":"test"}`)
}

func TestParameterToString_struct_array(t *testing.T) {
	temp := testStruct{
		StringValue: "test",
		IntValue:    100,
	}

	tempArray := []testStruct{temp}

	assert.Equal(t, ParameterToString(&tempArray), `[{"IntValue":100,"StringValue":"test"}]`)
}
