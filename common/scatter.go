package common

import (
	"github.com/astaxie/beego/logs"
	"sync"
)

func NewScatterSlice(data interface{}, do func(todo interface{}) interface{}) []interface{} {
	defer func() {
		if err := recover(); err != nil {
			logs.Info("%s\n", err)
		}
	}()
	//将interface{}转化为[]interface{}
	ret := InterSlice(data)
	//并发安全通道，用于保存处理完毕的数据体
	channel := make(chan interface{}, len(ret))
	//用于返回的数据切片
	resultSlice := make([]interface{}, 0)

	for _, v := range ret {
		go func(todo interface{}) {
			//函数式接口异步处理切片内数据
			resp := do(todo)
			//装入并发安全通道
			channel <- resp
		}(v)
	}

	var wg sync.WaitGroup
	wg.Add(cap(channel))
	go func() {
		for {
			select {
			case v := <-channel:
				resultSlice = append(resultSlice, v)
				wg.Done()
			}
		}
	}()
	wg.Wait()
	return resultSlice
}
