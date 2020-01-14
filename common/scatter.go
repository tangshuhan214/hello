package common

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"reflect"
	"sync"
)

func NewScatterSlice(data interface{}, do func(todo interface{}) interface{}) []interface{} {
	defer func() {
		if err := recover(); err != nil {
			logs.Info("%s\n", err)
		}
	}()

	v := reflect.ValueOf(data) //使用断言机制判断当前传入类型
	if v.Kind() != reflect.Slice {
		panic("方法体需要接收一个切片类型")
	}
	if data == nil {
		panic("集合数据为空")
	}
	l := v.Len()

	//并发安全通道，用于保存处理完毕的数据体
	channel := make(chan interface{}, l)
	//用于返回的数据切片
	resultSlice := make([]interface{}, 0)

	for i := 0; i < l; i++ {
		go func(todo interface{}) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("%s\n", err)
				}
			}()
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
