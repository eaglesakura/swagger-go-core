package swagger

import (
	"io"
)

/*
swaggerクライアントのデータ送信用のインターフェースを定義する.
*/
type DataPayload interface {
	/*
	content-length ヘッダに対応する.
	不定の場合、0未満の値を返却する.
	*/
	GetContentLength() int

	/*
	content-typeヘッダに対応する.
	*/
	GetContentType() string

	/*
	送信対象のデータ本体を取得する.
	*/
	OpenStream() io.ReadCloser
}

/*
swaggerクライアントのデータ送受信用インターフェースを定義する.

swaggerから各種パラメータが与えられるため、実際のFetch処理はインターフェース実装に任せられる.
*/
type FetchClient interface {
	/*
	Validatorと同等の機能が必須となる.
	*/
	ValidatorFactory

	/*
	APIのパスを指定する.

	URLに直接パラメータが含まれる場合、swaggerが自動的にパラメータを付与する.
	エンドポイントはdev/prod等の環境を変えることがあるため、swaggerからは渡されずに実装側で文字列連結を行う.

	ex)
		"api/v1/users/123/profile"
	*/
	SetApiPath(path string)

	/*
	httpメソッドを指定する
	ex)
		"GET", "POST", "PUT"...
	*/
	SetMethod(method string)

	/*
	URLに付与するQuery-Stringを定義する.
	値はすべて文字列にエンコードされるため、単純に&key=valueでつなげれば問題ないが、例えば不要なKeyを弾いたり、valueを入れ替えるといった処理を行える.
	*/
	AddQueryParam(key string, value string)

	/*
	URLに付与するQuery-Stringを定義する.
	値はすべて文字列にエンコードされるため、単純に&key=valueでつなげれば問題ないが、例えば不要なKeyを弾いたり、valueを入れ替えるといった処理を行える.
	*/
	AddHeader(key string, value string)

	/*
	送信対象のデータペイロードが指定される.

	例えば、JsonデータをPOSTする場合等に呼び出される.
	*/
	SetPayload(payload DataPayload)

	/*
	パラメータの指定が完了後、実際のhttp接続呼び出しをリクエストする際に呼び出される.

	resultPtrはswaggerで定義された構造体のポインタが渡されるため、json.unmarshal等でデシリアライズする.
	デイシリアライズ処理も実装者に任されているため、任意のライブラリを使用して問題ない.
	*/
	Fetch(resultPtr interface{}) error
}
