/*
httpリクエストと最終的なハンドラの関連付けを行う
*/
package swagger

import (
	"net/http"
	"github.com/gorilla/mux"
	"sort"
	"strings"
	"github.com/eaglesakura/swagger-go-core"
	"github.com/eaglesakura/swagger-go-core/errors"
)

/*
１つのPathに対してメソッドのマッピングを行う
*/
type MethodMapper struct {
	Path     string
	handlers map[string]swagger.HandleRequest
}

/*
ハンドルマッピングを行う.

最終的には gorilla/mux でバインディングされる.
*/
type HandleMapperImpl struct {
	Mappers map[string]*MethodMapper
}

/*
デフォルトのMapperを生成する
*/
func NewHandleMapper() swagger.HandleMapper {
	return &HandleMapperImpl{
		Mappers: map[string]*MethodMapper{},
	}
}

func initMapper(factory swagger.ContextFactory, r *mux.Route, handler swagger.HandleRequest) {
	r.Methods(handler.Method).Path(handler.Path).HandlerFunc(func(write http.ResponseWriter, request *http.Request) {
		//write.WriteHeader(http.StatusInternalServerError)
		//write.Write([]byte(fmt.Sprintf("Request[%v:%v] Mapping[%v:%v]", request.Method, request.URL.Path, handler.Method, handler.Path)))

		context := factory.NewContext(request)
		var responder swagger.Responder
		responderRef := &responder
		defer func() {
			// エラーからの復旧を行う
			if err := recover(); err != nil && (*responderRef) == nil {
				errorRef, ok := err.(error)
				if !ok {
					errorRef = errors.New(http.StatusInternalServerError, "Unknown Error")
				} else {
					errorRef = &errors.PanicError{
						Origin: errorRef,
					}
				}
				*responderRef = context.NewBindErrorResponse(errorRef)
			}
			context.Done(write, *responderRef)
		}()
		responder = handler.HandlerFunc(context, request)
	})
}

/*
Routerに登録する

GAE/Goの場合、 http.Handle("/path/to/api", router) でハンドリングが提供できる
*/
func (it *HandleMapperImpl) NewRouter(factory swagger.ContextFactory) *mux.Router {
	router := mux.NewRouter()

	for _, mapping := range it.ListMethodMappers() {
		for _, handler := range mapping.handlers {
			initMapper(factory, router.NewRoute(), handler)
		}
	}

	return router
}

/**
 * 同一URL内部でのメソッドバインディングを行う
 */
func (it *MethodMapper) PutHandler(request swagger.HandleRequest) {
	it.handlers[request.Method] = request
}

/**
 * ハンドラのバインディングを行う
 */
func (it *HandleMapperImpl) PutHandler(request swagger.HandleRequest) {
	list, ok := it.Mappers[request.Path]
	if !ok {
		list = &MethodMapper{
			Path:     request.Path,
			handlers: map[string]swagger.HandleRequest{},
		}
		it.Mappers[request.Path] = list
	}

	list.PutHandler(request)
}

type MethodMapperList []*MethodMapper

func (it *HandleMapperImpl) ListMethodMappers() MethodMapperList {
	result := MethodMapperList{}

	// map to list
	for _, value := range it.Mappers {
		result = append(result, value)
	}

	sort.Sort(result)

	return result
}

// Len is the number of elements in the collection.
func (it MethodMapperList) Len() int {
	return len(it)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (it MethodMapperList) Less(i, j int) bool {
	// 長さが違うならば、パスが長い方を優先してハンドリング
	if len(it[i].Path) != len(it[j].Path) {
		return len(it[i].Path) > len(it[j].Path)
	}

	// 長さが同じならば、compareする
	return strings.Compare(it[i].Path, it[j].Path) > 0
}

// Swap swaps the elements with indexes i and j.
func (it MethodMapperList) Swap(i, j int) {
	tmp := it[i]
	it[i] = it[j]
	it[j] = tmp
}
