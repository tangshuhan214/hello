package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)


// GO的PostJson方法，发送POST请求，参数要是一个JSON字符串，返回一个MAP(必须要用go来使其异步才能通，因为用了通道)
func PostJson(url string, data []byte) map[string]interface{} {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	httpResp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(httpResp.Body)
	defer httpResp.Body.Close()

	respData := map[string]interface{}{}
	_ = json.Unmarshal(body, &respData)
	return respData
}
