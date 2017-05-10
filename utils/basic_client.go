package utils

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/eaglesakura/swagger-go-core"
	"net/url"
	"fmt"
	"github.com/eaglesakura/swagger-go-core/errors"
)

type BasicFetchClient struct {
	/**
	 * ex) "http://example.com"
	 */
	Endpoint string
	Client   *http.Client
	Request  *http.Request

	apiPath  string
	queries  url.Values
	payload  swagger.DataPayload
}

/**
 * gen / default client
 */
func NewFetchClient(endpoint string, client *http.Client) *BasicFetchClient {
	req, _ := http.NewRequest("GET", "https://example.com", nil)
	return &BasicFetchClient{
		Client:client,
		Request:req,
		Endpoint:endpoint,
		queries:url.Values{},
	}
}

func (it*BasicFetchClient) NewValidator(value interface{}, isNil bool) swagger.ParameterValidator {
	return NewValidator(value, isNil)
}

/**
 * エンドポイントのPATHを指定する
 */
func (it*BasicFetchClient)SetApiPath(path string) {
	it.apiPath = path
}

/**
 * httpメソッドを指定する
 */
func (it*BasicFetchClient)SetMethod(method string) {
	it.Request.Method = method
}

/**
 * @param key   query key(not url encoded!)
 * @param value query value
 */
func (it*BasicFetchClient)AddQueryParam(key string, value string) {
	it.queries.Add(key, url.QueryEscape(value))
}

/**
 * @param key   query key(not url encoded!)
 * @param value query value
 */
func (it*BasicFetchClient)AddHeader(key string, value string) {
	it.Request.Header.Add(key, url.QueryEscape(value))
}

func (it*BasicFetchClient)SetPayload(payload swagger.DataPayload) {
	it.payload = payload
}

func (it*BasicFetchClient)Fetch(resultPtr interface{}) error {

	// build url
	{
		reqUrl, err := url.Parse(AddPath(it.Endpoint, it.apiPath))
		if err != nil {
			return err
		}
		if len(it.queries) > 0 {
			reqUrl.RawQuery = it.queries.Encode()
		}

		it.Request.URL = reqUrl
		it.Request.Host = reqUrl.Host
	}

	// request payload
	if it.payload != nil {
		it.Request.Header.Set("Content-Type", it.payload.GetContentType())
		it.Request.Header.Set("Content-Length", fmt.Sprintf("%v", it.payload.GetContentLength()))
		it.Request.Body = it.payload.OpenStream()
		defer it.Request.Body.Close()
	}

	resp, err := it.Client.Do(it.Request)
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return err
	} else if (resp.StatusCode / 100) != 2 {
		// statuscode error
		return errors.New(int32(resp.StatusCode), "FetchError[%v]%v]", it.Request.Method, it.Request.URL.Path)
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// decode json
	return json.Unmarshal(buf, resultPtr)
}