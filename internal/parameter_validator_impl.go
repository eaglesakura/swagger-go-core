package swagger

import (
	"regexp"
	"github.com/eaglesakura/swagger-go-core"
)


type ParameterValidatorImpl struct {
	Value     interface{}
	IsNil     bool

	pattern   *regexp.Regexp
	required  *bool
	minLength *int
	maxLength *int
}

func (it *ParameterValidatorImpl)Required(set bool) swagger.ParameterValidator {
	it.required = &set
	return it
}

func (it *ParameterValidatorImpl)Pattern(pattern string) swagger.ParameterValidator {
	pattern = (pattern[1:len(pattern) - 2])

	exp, err := regexp.Compile(pattern)
	if err != nil {
		it.pattern = exp
	}
	return it
}

func (it *ParameterValidatorImpl)MinLength(len int) swagger.ParameterValidator {
	it.minLength = &len
	return it
}

func (it *ParameterValidatorImpl)MaxLength(len int) swagger.ParameterValidator {
	it.maxLength = &len
	return it
}

func (it ParameterValidatorImpl)Valid(factory swagger.ValidatorFactory) bool {
	if it.IsNil && it.required != nil && *it.required {
		return false
	}

	// nil && not required
	if it.IsNil && (it.required == nil || !(*it.required)) {
		return true
	}

	if strValue, ok := it.Value.(*string); strValue != nil && ok {
		if it.minLength != nil && len(*strValue) < *it.minLength {
			return false
		}

		if it.maxLength != nil && len(*strValue) > *it.maxLength {
			return false
		}

		if it.pattern != nil && !it.pattern.Match([]byte(*strValue)) {
			return false
		}
	} else {
		if validatable, ok := it.Value.(swagger.Validatable); validatable != nil && ok {
			return validatable.Valid(factory)
		}
	}

	// valid param
	return true
}