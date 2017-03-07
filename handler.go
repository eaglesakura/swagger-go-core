package swagger

import (
	"net/http"
	"github.com/gorilla/mux"
)

/**
 * APIエンドポイント(&METHOD) ごとに用意されるハンドリングデータ
 */
type HandleRequest struct {
	// /path/to/api
	Path        string

	// GET, POST, PUT...
	Method      string

	// Function
	// DefaultApiController.
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
