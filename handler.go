package swagger

import (
	"github.com/gorilla/mux"
	"net/http"
)

/*
APIエンドポイント(&METHOD) ごとに用意されるハンドリングデータ.
swagger-codegenにより自動生成される.
*/
type HandleRequest struct {
	/*
		internal.

		APIへのパスを設定する。

		swagger.[json|yaml]の"{basePath}{apiPath}"で指定される。
		エンドポイントはdev/stg等の環境で切り替わる可能性があるため、実装側に任されている。

		ex)
			"api/v1/user/"
	*/
	Path string

	/*
		internal.

		http methodを指定する。

		ex)
			"GET", "POST", "PUT"...
	*/
	Method string

	/*
		internal.

		ハンドリング用の関数.
	*/
	HandlerFunc func(context RequestContext, request *http.Request) Responder
}

/*
HandleRequestと実際のRouterのマッピングを行なう。
*/
type HandleMapper interface {

	/*
		リクエストハンドラを追加する
	*/
	PutHandler(request HandleRequest)

	/*
		gorilla/muxで処理するためのRouterを生成する.
	*/
	NewRouter(factory ContextFactory) *mux.Router
}
