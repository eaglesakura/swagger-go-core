package swagger

/*
httpリクエストから構造体へ値をバインドする機能を提供する.
*/
type RequestBinder interface {
	ValidatorFactory

	/*
	URLからパラメータを取り出す
	*/
	BindPath(key string, resultType string, result interface{}) error

	/*
	Query Stringからパラメータを取り出す
	*/
	BindQuery(key string, resultType string, result interface{}) error

	/*
	http headerからパラメータを取り出す
	*/
	BindHeader(key string, resultType string, result interface{}) error

	/*
	FORM形式でリクエストされたデータからパラメータを取り出す
	*/
	BindForm(key string, resultType string, result interface{}) error

	/*
	Bodyからパラメータをパースする
	*/
	BindBody(resultType string, result interface{}) error
}
