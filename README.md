# swagger-go-core

[swagger-codegen](https://github.com/eaglesakura/swagger-codegen) で出力されたコードを利用するためのクライアントライブラリ.

# Install

```bash
go get "github.com/eaglesakura/swagger-go-core"
```

# サーバーサイドへの組み込み

```go
// GAE/Goでの例

import (
	"api_package_name"
	"github.com/eaglesakura/swagger-go-core"
	"github.com/eaglesakura/swagger-go-core/utils"
	"net/http"
	"my_server"
)

func init() {

	// swagger.ContextFactory を実装した構造体を生成する
	contextFactory := &my_server.ContextFactoryImpl{}
	mapper := utils.NewHandleMapper()

	// swagger.yamlで [MyTag] タグを付与したAPIのハンドラを登録する
	// タグを指定しない場合、 [Default] タグとして生成される
	{
		api := api_package_name.NewMyTagApiController()
		// api.Handle{Path/To/Api/Name}{MethodName}(HandlerFunction)
		api.HandleHogeFugaGet(HogeFugaHandler)  // GET /hoge/fuga のAPIハンドラを登録する
		api.MapHandlers(mapper)
	}

	http.Handle("/api/", mapper.NewRouter(contextFactory))
}

// GET /hoge/fuga APIのハンドラを定義する
// RequestContextインターフェースはhttpリクエストが発生するごとに生成される
// リクエストパラメータ(引数:param)はすべてswagger-codegenが生成した構造体に収められている
func HogeFugaHandler(_ctx swagger.RequestContext, param *api_package_name.HogeFugaGetParams) swagger.Responder {
	ctx := _ctx.(my_server.RequestContext)

  ctx.HogeFuga()

	// swagger.Responderインターフェースをimplした構造体を返却する
	// swgger.yamlに戻り値が定義されている場合は自動的に生成される.
	// 生成された構造体のほかにも、インターフェースを満たしていれば自由に返却可能.
	// 生成される構造体のメンバはすべてポインタで定義される. これは "omitempty" 動作と空文字を明示的に区別することが必要なため.
	result := &api_package_name.HogeFugaInfo{}
	result.Message = swag.String("Hello World!")

	return result
}

```

# クライアントサイドへの組み込み

```go

import (
	"api_package_name"
	"github.com/eaglesakura/swagger-go-core/swag"
	"github.com/eaglesakura/swagger-go-core/utils"
	"net/http"
	"time"
)

func callHogeFuga() {
	// Fetch Clientを生成する
	// インターフェースを満たしていれば良いため、必要に応じてカスタマイズが可能
	client := utils.NewFetchClient("https://your-api-endpoint.example.com/api", &http.Client{Timeout: (time.Duration(30) * time.Second)})
	// 自動生成されたコードを呼び出す
	api := api_package_name.NewMyTagApi()

	// APIを呼び出す
	// 戻り値構造体も自動生成される
	result := api_package_name.HogeFugaInfo{}
	if err := api.HogeFugaGet(client, request, &result); err != nil {
		fmt.Printf("API error[%v]", err.Error())
		return nil, err
	}

	// Hello World!
	fmt.Printf("%v\n", *result.Message)
}
```
