package swagger

import (
	"net/http"
)

/*
swaggerサーバーの1リクエスト(http.Request)が発生するごとに生成・管理されるContext.

内部の具体的実装はサーバーアプリの実装者に任せられる.
*/
type RequestContext interface {
	/*
	request -> parameterへのバインド制御インターフェースを生成する.

	http/httpsの1リクエストごとに呼び出される.
	認証等の事前チェックを行いたい場合はこの呼出前に行い、問題があればerrorを返却する.
	errorが返却された場合、swaggerは RequestContext.NewBindErrorResponse に戻り値生成を移譲する。

	デフォルト実装が用意 / Utilされている.
		import(
			swagger_utils "github.com/eaglesakura/swagger-go-core/utils"
		)
		swagger_utils.NewRequestBinder()
	*/
	NewRequestBinder(request *http.Request) (RequestBinder, error)

	/*
	Request -> Parameterのバインド失敗時に呼び出され、エラーレスポンスを生成する。
	*/
	NewBindErrorResponse(err error) Responder

	/*
	ハンドリングの完了処理を行う。
	このメソッドは制御の最後にかならず呼び出される。
	このメソッドでは基本的に response.WriteResponse(writer, {Producer})を呼び出す。
	どのProducerが利用されるかは、Contextの実装者に任される。
	必要に応じてリソースの開放処理を行う。aetestで生成されたContext等。
	*/
	Done(writer http.ResponseWriter, response Responder)
}

type ContextFactory interface {
	/*
	1ハンドリングごとのコンテキストを生成する
	*/
	NewContext(request *http.Request) RequestContext
}
