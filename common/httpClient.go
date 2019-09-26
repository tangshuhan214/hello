package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// GO的PostJson方法，发送POST请求，参数要是一个JSON的BYTE字符集，返回一个MAP
func PostJson(url string, data []byte) map[string]interface{} {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	httpResp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(httpResp.Body)
	defer httpResp.Body.Close()

	respData := map[string]interface{}{}
	_ = json.Unmarshal(body, &respData)
	return respData
}
