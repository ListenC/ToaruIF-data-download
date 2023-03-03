package HTTP

import (
	"io"
	"net/http"
)

/**
 * @description: 发起GET请求
 * @param {string} URL 请求地址
 * @return {[]byte} 返回体
 * @return {*http.Response} 返回内容
 * @return {error} 错误信息
 */
func Get(URL string) ([]byte, *http.Response, error) {
	Request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, nil, err
	}

	Client := &http.Client{}

	Response, err := Client.Do(Request)
	if err != nil {
		return nil, Response, err
	}
	defer Response.Body.Close()

	Body, err := io.ReadAll(Response.Body)
	if err != nil {
		return nil, Response, err
	}

	return Body, Response, nil
}
