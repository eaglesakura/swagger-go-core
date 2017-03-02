package swagger

type ParameterValidator interface {
	Required(set bool) ParameterValidator

	Pattern(pattern string) ParameterValidator

	MinLength(len int) ParameterValidator

	MaxLength(len int) ParameterValidator

	Valid(factory ValidatorFactory) bool
}

type ValidatorFactory interface {
	NewValidator(value interface{}, isNil bool) ParameterValidator
}

type Validatable interface {
	Valid(factory ValidatorFactory) bool
}