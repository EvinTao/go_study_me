package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func ApiGet(url string) {
	/* 示例1：GET */
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++")
	/* 发请求收应答 */
	ack, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	/* 读取应答正文 */
	ackBody, err := io.ReadAll(ack.Body)
	/* 关闭应答正文，释放资源，无论是否异常 */
	ack.Body.Close()
	if err != nil {
		panic(err)
	}

	/* 输出应答状态 */
	fmt.Printf("HTTP Response StatusCode: %d\n", ack.StatusCode)
	fmt.Printf("HTTP Response Status: %s\n", ack.Status)

	/* 输出应答头域 */
	fmt.Printf("HTTP Response HEADER: %s\n", ack.Header.Get("my-http-head"))

	/* 输出应答正文 */
	fmt.Printf("HTTP Response BODY: %s\n", ackBody)
}

func ApiPost(url string, query map[string]string, body []byte) ([]byte, error) {
	fmt.Println("---------------------------------------------")
	// 内网 使用http短连接
	tr := &http.Transport{DisableCompression: false}
	var ack *http.Response
	defer func() {
		if ack != nil && ack.Body != nil {
			_ = ack.Body.Close()
		}
		tr.CloseIdleConnections()
	}()

	/* 创建请求对象 */
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		//panic(err)
		return nil, err
	}

	if query != nil && len(query) > 0 {
		q := req.URL.Query()
		for k, v := range query {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	/* 设置请求头域 */
	req.Header.Set("X-Token", "token")

	/* 发请求收应答 */
	ack, err = http.DefaultClient.Do(req)
	if err != nil {
		//panic(err)
		return nil, err
	}

	/* 读取应答正文 */
	content, err := ioutil.ReadAll(ack.Body)
	/* 关闭应答正文，释放资源，无论是否异常 */
	ack.Body.Close()
	if err != nil {
		//panic(err)
		return nil, err
	}

	/* 输出应答状态 */
	if len(content) == 0 {
		return nil, errors.New("empty body")
	}
	return content, err

}

func main() {
	//var query = make(map[string]string)
	//var body = "{\"geo\": \"go\",\"offset\": 1,\"num\": 10}"
	//content, err := ApiPost("api.example.com/list", query, []byte(body))
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf(string(content))
}
