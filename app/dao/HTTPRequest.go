package dao

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 定义一个通用的HTTP响应结构体
type HTTPResponse struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

// 定义一个通用的HTTP请求方法
func HTTPRequest(method, url string, body []byte, headers map[string]string) (*HTTPResponse, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// 设置请求头
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 封装响应到HTTPResponse结构体
	httpResponse := &HTTPResponse{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       bodyBytes,
	}

	return httpResponse, nil
}

// 示例函数，展示如何使用HTTPRequest
func callAPI() error {
	url := "http://example.com/api/some-endpoint"
	body := []byte(`{"key": "value"}`)
	headers := map[string]string{"Content-Type": "application/json"}

	// 发起请求
	httpResponse, err := HTTPRequest("POST", url, body, headers)
	if err != nil {
		return err
	}

	// 处理响应
	fmt.Printf("Status Code: %d\n", httpResponse.StatusCode)
	fmt.Println("Headers:")
	for k, values := range httpResponse.Headers {
		for _, v := range values {
			fmt.Printf("%s: %s\n", k, v)
		}
	}
	fmt.Println("Body:")
	fmt.Println(string(httpResponse.Body))

	// 如果需要，可以将响应体解析为特定类型
	var result map[string]interface{}
	if err := json.Unmarshal(httpResponse.Body, &result); err != nil {
		return err
	}

	return nil
}

func main() {
	err := callAPI()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
