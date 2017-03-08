package swagger

const (
	BindExtra_Authentication = "Authentication"
)

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

	/**
	 * その他の特殊データのバインディングを行う
	 * このデータはバリデーションの対象外となるが、errorを返却した場合はその時点でbind失敗とみなす。
	 *
	 * 主に認証情報等で使用される
	 * @param resultType BindExtra_Authentication
	 */
	BindExtra(resultType string, result interface{}) error
}
