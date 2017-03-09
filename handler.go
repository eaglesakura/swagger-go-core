package swagger

import (
	"net/http"
	"github.com/gorilla/mux"
)

/**
 * APIエンドポイント(&METHOD) ごとに用意されるハンドリングデータ
 * swagger-codegenにより自動生成される。
 */
type HandleRequest struct {
	/**
	 * APIへのパスを設定する。
	 * これはswagger.jsonの"{basePath}{apiPath}"で指定される。
	 */
	Path        string

	/**
	 * http methodを指定する。
	 * GET, POST, PUT...
	 */
	Method      string

	/**
	 * ハンドリング用の関数を指定する。
	 */
	HandlerFunc func(context RequestContext, request *http.Request) Responder
}

/**
 * HandleRequestと実際のRouterのマッピングを行なう。
 */
type HandleMapper interface {
	/**
	 * リクエストハンドラを追加する
	 */
	PutHandler(request HandleRequest)

	/**
	 * 最終的なハンドリングを行なうためのRouterを生成する
	 */
	NewRouter(factory ContextFactory) *mux.Router
}
