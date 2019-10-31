package common

import (
	"bytes"
	"encoding/json"
)

//入参转为MAP集合
func ParamsInCtrl(data []byte, v interface{}) interface{} {
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	_ = json.Unmarshal(data, &v)
	_ = d.Decode(&v)
	return v
}
