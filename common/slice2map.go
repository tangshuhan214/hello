package common

import (
	"encoding/json"
	"reflect"
	"strings"
	"unsafe"
)

//这里接收一个interface{}空接口将切片结构体集合转为一个切片map集合，用于处理在对象返回时需要装载特殊更多的返回值
func StructSlice2Map(list interface{}) []map[string]interface{} {
	var final []map[string]interface{}
	v := reflect.ValueOf(list) //使用断言机制判断当前传入类型
	if v.Kind() != reflect.Slice {
		panic("方法体需要接收一个切片类型") //不是切片立即抛错
	}
	l := v.Len()
	ret := make([]interface{}, l) //开始将传入切片转换为[]interface{}类型
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}

	for _, obj := range ret {
		jsonBytes, _ := json.Marshal(obj)          //将结构体转为JSON字符串
		var only map[string]interface{}            //这里做两次转换的原因是去除首字母大写
		only = createJsonUseNum(string(jsonBytes)) //将JSON字符串转为Map[string]interface{}
		for key, value := range only { //判断该属性是否为json.Number,如果是调用转换方法
			v := reflect.ValueOf(value)
			of := reflect.TypeOf(value) //使用断言机制判断当前传入类型
			if of.Name() == "Number" {
				number := v.Interface().(json.Number)
				only[key] = accuracy(number)
			}
		}
		final = append(final, only)
	}
	return final
}

//将超大整形属性转换为json.Number以免丢失精度
func createJsonUseNum(source string) map[string]interface{} {
	data := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(source))
	dec.UseNumber()
	dec.Decode(&data)
	return data
}

//将json.Number转化为int
func accuracy(number json.Number) int {
	i, _ := number.Int64()
	idPointer := (*int)(unsafe.Pointer(&i))
	idd16 := *idPointer
	return idd16
}
