package utils

import (
	"encoding/json"
	"fmt"
	"github.com/eaglesakura/swagger-go-core"
	"github.com/eaglesakura/swagger-go-core/errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

type FetchFunction func(it *BasicFetchClient, resultPtr interface{}) error

/*
saggerクライアントとして使用可能なデフォルトFetchClient実装を提供する
*/
type BasicFetchClient struct {
	/*
		ex)
			"http://example.com/api"
	*/
	Endpoint string

	/*
		デフォルトで使用するHttpクライアント
	*/
	Client *http.Client

	/*
		セットアップ対象のHttp Request
	*/
	Request *http.Request

	/*
		Fetch時の処理を移譲するDelegate
	*/
	CustomFetch FetchFunction

	ApiPath string

	Queries url.Values

	Payload swagger.DataPayload

	/*
		Raw Response.
	*/
	Response *http.Response

	/*
		Raw Response Body.
	*/
	ResponseBody []byte
}

/*
saggerクライアントとして使用可能なデフォルトFetchClient実装を生成する

ex)
	NewFetchClient("https://your-gcp-project.appspot.com", httpClient)
*/
func NewFetchClient(endpoint string, client *http.Client) *BasicFetchClient {
	req, _ := http.NewRequest("GET", "https://example.com", nil)
	return &BasicFetchClient{
		Client:      client,
		Request:     req,
		Endpoint:    endpoint,
		CustomFetch: basicFetchFunction,
		Queries:     url.Values{},
	}
}

func (it *BasicFetchClient) NewValidator(value interface{}, isNil bool) swagger.ParameterValidator {
	return NewValidator(value, isNil)
}

func (it *BasicFetchClient) SetApiPath(path string) {
	it.ApiPath = path
}

func (it *BasicFetchClient) SetMethod(method string) {
	it.Request.Method = method
}

func (it *BasicFetchClient) AddQueryParam(key string, value string) {
	it.Queries.Add(key, url.QueryEscape(value))
}

func (it *BasicFetchClient) AddHeader(key string, value string) {
	it.Request.Header.Add(key, value)
}

func (it *BasicFetchClient) SetPayload(payload swagger.DataPayload) {
	it.Payload = payload
}

func basicFetchFunction(it *BasicFetchClient, resultPtr interface{}) error {
	// request Payload
	if it.Payload != nil {
		it.Request.Header.Set("Content-Type", it.Payload.GetContentType())
		it.Request.Header.Set("Content-Length", fmt.Sprintf("%v", it.Payload.GetContentLength()))
		it.Request.Body = it.Payload.OpenStream()
		defer func() { _ = it.Request.Body.Close() }()
	}

	resp, err := it.Client.Do(it.Request)
	it.Response = resp
	if resp != nil && resp.Body != nil {
		defer func() { _ = resp.Body.Close() }()
	}

	if err != nil {
		return err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	it.ResponseBody = buf

	if (resp.StatusCode / 100) != 2 {
		// status code error
		return errors.New(int32(resp.StatusCode), "FetchError[%v]%v]: code='%v' / body=%v", it.Request.Method, it.Request.URL.Path, resp.StatusCode, string(buf))
	}

	// decode json
	if resultPtr != nil {
		return json.Unmarshal(buf, resultPtr)
	} else {
		return nil
	}
}

func (it *BasicFetchClient) Fetch(resultPtr interface{}) error {
	// build url
	{
		reqUrl, err := url.Parse(AddPath(it.Endpoint, it.ApiPath))
		if err != nil {
			return err
		}
		if len(it.Queries) > 0 {
			reqUrl.RawQuery = it.Queries.Encode()
		}

		it.Request.URL = reqUrl
		it.Request.Host = reqUrl.Host
	}

	return it.CustomFetch(it, resultPtr)
}
