package common

import (
	"reflect"
	"sync"
)

type QueueThreadUtils struct {
	channel     chan interface{}
	resultSlice []interface{}
}

func NewQueue(size int) *QueueThreadUtils {
	s := &QueueThreadUtils{
		channel:     make(chan interface{}, size),
		resultSlice: make([]interface{}, 0),
	}
	return s
}

func (in *QueueThreadUtils) ScatterSlice(data interface{}, do func(todo interface{}) interface{}) []interface{} {
	v := reflect.ValueOf(data) //使用断言机制判断当前传入类型
	if v.Kind() != reflect.Slice {
		panic("方法体需要接收一个切片类型") //不是切片立即抛错
	}
	if data == nil {
		panic("集合数据为空") //
	}
	l := v.Len()
	ret := make([]interface{}, l) //开始将传入切片转换为[]interface{}类型
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}

	for _, v := range ret {
		go func(todo interface{}) {
			resp := do(todo)
			in.channel <- resp
		}(v)
	}

	var wg sync.WaitGroup
	wg.Add(cap(in.channel))
	go func() {
		for {
			select {
			case v := <-in.channel:
				in.resultSlice = append(in.resultSlice, v)
				wg.Done()
			}
		}
	}()
	wg.Wait()
	return in.resultSlice
}
