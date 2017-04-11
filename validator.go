package swagger

/**
 * 入力されたパラメータのバリデーションを行なうインターフェース
 */
type ParameterValidator interface {
	/**
	 * パラメータがNotNullである
	 */
	Required(set bool) ParameterValidator

	/**
	 * 文字列の正規表現パターンを指定する
	 */
	Pattern(pattern string) ParameterValidator

	/**
	 * 文字列の最小長を指定する
	 */
	MinLength(len int) ParameterValidator

	/**
	 * 文字列の最大長を指定する
	 */
	MaxLength(len int) ParameterValidator

	/**
	 * valuesに指定したいずれかの値に合致する必要がある
	 */
	EnumPattern(values []string) ParameterValidator

	/**
	 * 全ての検証に合格していればtrue
	 */
	Valid(factory ValidatorFactory) bool
}

type ValidatorFactory interface {
	NewValidator(value interface{}, isNil bool) ParameterValidator
}

type Validatable interface {
	Valid(factory ValidatorFactory) bool
}

type EnumValidatable interface {
	Valid(pattern []string) bool
}