package swagger

import (
	"net/http"
)

/**
 * http.Requestが発生するごとに生成されるContext
 *
 * 内部の具体的実装はGAE/Go Appの実装者に任せられる。
 */
type RequestContext interface {
	/**
	 * Validatorを生成させる
	 *
	 * Contextをチェックし、適したValidatorを生成させる。
	 *
	 * ex) swagger.NewValidatorFactory()
	 */
	NewValidatorFactory() ValidatorFactory

	/**
	 * request -> parameterへのバインド制御インターフェースを生成する
	 *
	 * 認証等の事前チェックを行いたい場合はこの呼出前に行い、問題があればerrorを返却する
	 * ex) swagger.NewRequestBinder()
	 */
	NewRequestBinder(request *http.Request) (RequestBinder, error)

	/**
	 * Request -> Parameterのバインド失敗時に呼び出される。
	 */
	NewBindErrorResponse(err error) Responder

	/**
	 * ハンドリングの完了処理を行う。
	 *
	 * このメソッドは制御の最後にかならず呼び出される。
	 * 必要に応じてリソースの開放処理を行う。
	 */
	Done(writer http.ResponseWriter, response Responder)
}

type ContextFactory interface {
	/**
	 * 1ハンドリングごとのコンテキストを生成する
	 */
	NewContext(request *http.Request) RequestContext
}
