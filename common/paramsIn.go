package common

import (
	"bytes"
	"encoding/json"
)

func ParamsInCtrl(data []byte) map[string]interface{} {
	resp := map[string]interface{}{}
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	_ = json.Unmarshal(data, &resp)
	_ = d.Decode(&resp)
	return resp
}