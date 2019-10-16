package business

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"github.com/astaxie/beego"
	"hello/common"
	"sort"
)

type ZhiFuBaoBiz struct {
}

func (this *ZhiFuBaoBiz) InsertPay(params map[string]interface{}) map[string]interface{} {
	getConfig(&params)

	var sslice []string
	for key, _ := range params {
		sslice = append(sslice, key)
	}
	sort.Strings(sslice)
	//在将key输出
	var check map[string]interface{}
	for _, v := range sslice {
		check[v] = params[v]
	}

	after, _ := json.Marshal(params)
	mString := string(after)
	params["CHECK_CODE"] = md5V(mString)

	url := beego.AppConfig.String("payUrls")

	respData := common.PostJson(url, params)
	return respData
}

func (this *ZhiFuBaoBiz) RefundPay(params map[string]interface{}) map[string]interface{} {
	getConfig(&params)
	url := beego.AppConfig.String("refundUrls")

	respData := common.PostJson(url, params)
	return respData
}

func getConfig(params *map[string]interface{}) {
	par := *params
	par["URL"] = "https://openapi.alipay.com/gateway.do"
	par["APP_ID"] = "2017110609759293"
	par["FORMAT"] = "json"
	par["CHARSET"] = "UTF-8"
	par["ALIPAY_PUBLIC_KEY"] = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAm2pWzdLeOizM3zMf1PkYuctnjn8Vl2s1e9KditLNy8ta/YHgoKjPq+fZIMOsjCJK7zg6RD+IP6X0zbdtMFnhmyUqJwUJzG01Z748RT1SSRnttkodGIt/tgy87Sg535dsMDKHkPWM5IfApkup3ykSq7aEQGhIPf/+v6pMl6pPJ89tc53NAdSkEHPujkReKkUYOIuB2wnUeidThwPCdnmpn+q602+SN793t9AQaLvW9uI3lU8x4xf+KkZD1QnxkCB0xf5DTID5eZgug2LnjyYax0GOLvvR3LGCwgOvLWDVy4pv6GP+kHxq+snpc3h0N8IafOfDE04/U5hbPgJaTGvscwIDAQAB"
	par["APP_PRIVATE_KEY"] = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQC74eJPNFeYZb4H7Tyf88oeRcZ4lUEFJShOVKhnd4rdeRIO6LSCA2lzdVa6/1nHe0oby7sIe6X7SUVtJP9VcWrNfGZcudYNXv8RGy7VSUsVnReD2lZBRmKkqNWXTa2FXrnkGRce+BxSuZh6vTQQwczbxPz3jWx6yqarJ638vu7R1O9DBsOqZoFotMNmKQeuFFR26PhOxzqigX4vQyiOZ9VWAJIjIXSwGlpQcg5c1kBpfoQ8eFHQwSlPAiWXDRzNaIkUoSLWErcW3x0FQVluGnCJT1iJpXREbJVtG07II6kPvXei0cqyS0rSE9S5y5jXxo1fyXgR2LEzcNMShUtFCJNfAgMBAAECggEAIp3k36kKNqcQU0+Pqvz1EYzTm6YMW9FT0tLgUGgDKvlCrYKMt6O/SymkeEHSHeIGboakCFUX6OvAIfL+JJUIE6JHDfCjksUdA97ZSyz865eNHJse51n7lNESwZTrUrZ4U6GX0/ns5gUSJhM0Q54hi6CI52ekRB+Qg9qEwRvAzuI24h5hT3ePlEDyEIEYqLqe2ewLjE/P0btwNtKTAIW1xhppRiGXx6RZHwAu3cn1LNPYzNKHonsAds6DT0fF27TslV3yo0KIImQm0V5RACY2Xxx5+Di6zxi7Dpb/prdGRS9j+wxRm8CAH2mYs0U7lYzkeFAw/+kBv91L8J4R0guaaQKBgQDnDQNjxobHvnzZOkUmAkonBshjBF1UNWmC43UfSe6BccmSVRvkvj+975hWg0k473t9/t2OWh0HUY9RciqL7d9tv71tnWaJnehqX8Sn0Gz5JxZk0JgSPvrhvdc/aB8/UKxot9WxcRR84I1M8My3AmdHRNNMYIJ6buefKcnQze97DQKBgQDQK447HCTGfs8nmhaySnjWvH6vBVrdrgwshNbsThrvciaEqKsvXretSggsB+u76+GI82UH67LY6jLdAqpGV5RLbhtx3KhKHdCzluXn2+XDw252EmuOa2MygE4KYLJ+TSlqHVSTtY/ZwRD9mdiubRM/S6VYsy5dFJFuzoySpya9GwKBgQDTGLOhJFAeDO76dV+aE3t2Xp8UwHcYqdglqvVmSeDsSW72EGZ0vlF0koRnfnmW2E7G1eXM2o4tEppTunAe+o0pM4a5sJZvY2NmVOtSu33kwY6XO3HFMd98Aju4BcSOz4FGB7fo77zdPzg6NMOE7WA44CwBWye6/rsGU2K9MHn9vQKBgGWzgc9bFpRrS6WbYE9GlvbCLFoxkY0QBR6S37WfCwXEjRDunoOMEMx2iLuKOx8aRJt13fwaqMvUz3iuXqXzD54ycvITzZw4KMg0hqnaAsy7Y/IHWcjAqjv39yiWyV1vMTaIkdOANoE6E6TyTqwY2fhoaqWFFLeg3tR10LOtOf3VAoGANNOxSmKExZkNGu++ESJ0c3S2GTlRfum8xSgRt1bMAmZQzFBranBx5C13rm8+vN5exxd3NQL0DBF7uldXnTbHtU3xKfm/8B74r+s6abyfqneWfoRbjePTnewrghN0MGzgpyVQfgbgMsTukTOpT3yGbi/MQi78DYA14GMz0MK0uMw="
	par["SIGN_TYPE"] = "RSA2"
}

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := hex.EncodeToString(h.Sum(nil))
	encodeString := base64.StdEncoding.EncodeToString([]byte(cipherStr))
	return encodeString
}
