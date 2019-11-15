package swagger

import (
	"github.com/eaglesakura/swagger-go-core"
	"regexp"
)

type ParameterValidatorImpl struct {
	Value interface{}
	IsNil bool

	pattern     *regexp.Regexp
	enumPattern *[]string
	required    *bool
	minLength   *int
	maxLength   *int
}

func NewValidator(value interface{}, isNil bool) swagger.ParameterValidator {
	return &ParameterValidatorImpl{Value: value, IsNil: isNil}
}

func (it *ParameterValidatorImpl) Required(set bool) swagger.ParameterValidator {
	it.required = &set
	return it
}

func (it *ParameterValidatorImpl) Pattern(pattern string) swagger.ParameterValidator {
	pattern = (pattern[1 : len(pattern)-2])

	exp, err := regexp.Compile(pattern)
	if err != nil {
		it.pattern = exp
	}
	return it
}

func (it *ParameterValidatorImpl) MinLength(len int) swagger.ParameterValidator {
	it.minLength = &len
	return it
}

func (it *ParameterValidatorImpl) MaxLength(len int) swagger.ParameterValidator {
	it.maxLength = &len
	return it
}

/*
valuesに指定したいずれかの値に合致する必要がある
*/
func (it *ParameterValidatorImpl) EnumPattern(values []string) swagger.ParameterValidator {
	it.enumPattern = &values
	return it
}

func (it *ParameterValidatorImpl) Valid(factory swagger.ValidatorFactory) bool {

	if it.IsNil {
		// nilチェック
		if it.required != nil && *it.required {
			return false
		}

		// 必須でないならばこれ以上のチェックは必要ない
		if it.required == nil || !(*it.required) {
			return true
		}
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

		// enum一致をチェックする
		if it.enumPattern != nil {
			validEnum := false
			for _, value := range *it.enumPattern {
				//
				if value == (*strValue) {
					validEnum = true
				}
			}

			// enumに一致しないのでNG
			if !validEnum {
				return false
			}
		}
	} else {
		if it.enumPattern != nil {
			if validatable, ok := it.Value.(swagger.EnumValidatable); validatable != nil && ok {
				// 事前指定されたEnumに一致しないのでNG
				if !validatable.Valid(*it.enumPattern) {
					return false
				}
			}
		}

		if validatable, ok := it.Value.(swagger.Validatable); validatable != nil && ok {
			return validatable.Valid(factory)
		}
	}

	// valid param
	return true
}
