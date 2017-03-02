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

type RequestBinder interface {
	/**
	 * URLからパラメータを取り出す
	 */
	BindPath(key string, resultType string, result interface{}) error

	/**
	 * Query Stringからパラメータを取り出す
	 */
	BindQuery(key string, resultType string, result interface{}) error

	/**
	 * http headerからパラメータを取り出す
	 */
	BindHeader(key string, resultType string, result interface{}) error

	/**
	 * FORM形式でリクエストされたデータからパラメータを取り出す
	 */
	BindForm(key string, resultType string, result interface{}) error

	/**
	 * Bodyからパラメータをパースする
	 */
	BindBody(resultType string, result interface{}) error
}
