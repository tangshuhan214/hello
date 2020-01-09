package common

import (
	"fmt"
	"reflect"
	"sync"
)

type FiberSliceUtils struct {
	channel     chan interface{}
	resultSlice []interface{}
}

func NewScatterSlice(data interface{}, do func(todo interface{}) interface{}) []interface{} {
	v := reflect.ValueOf(data) //使用断言机制判断当前传入类型
	if v.Kind() != reflect.Slice {
		panic("方法体需要接收一个切片类型")
	}
	if data == nil {
		panic("集合数据为空")
	}
	l := v.Len()
	ret := make([]interface{}, l) //开始将传入切片转换为[]interface{}类型
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}

	s := &FiberSliceUtils{
		channel:     make(chan interface{}, len(ret)),
		resultSlice: make([]interface{}, 0),
	}

	for _, v := range ret {
		go func(todo interface{}) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("%s\n", err)
				}
			}()
			//函数式接口异步处理切片内数据
			resp := do(todo)
			//装入并发安全通道
			s.channel <- resp
		}(v)
	}

	var wg sync.WaitGroup
	wg.Add(cap(s.channel))
	go func() {
		for {
			select {
			case v := <-s.channel:
				s.resultSlice = append(s.resultSlice, v)
				wg.Done()
			}
		}
	}()
	wg.Wait()
	return s.resultSlice
}
