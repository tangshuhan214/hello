package common

import (
	"errors"
	"reflect"
	"sync"
)

func NewScatterSlice(data interface{}, do func(todo interface{}) interface{}) (resultSlice []interface{}, e error) {
	defer func() {
		if err := recover(); err != nil {
			e = errors.New(err.(string))
		}
	}()

	v := reflect.ValueOf(data) //使用断言机制判断当前传入类型
	if v.Kind() != reflect.Slice {
		panic("方法体需要接收一个切片类型")
	}
	l := v.Len()
	if l == 0 {
		panic("集合数据为空")
	}

	//并发安全通道，用于保存处理完毕的数据体
	channel := make(chan interface{}, l)

	for i := 0; i < l; i++ {
		go func(todo interface{}) {
			//函数式接口异步处理切片内数据
			resp := do(todo)
			//装入并发安全通道
			channel <- resp
		}(v.Index(i).Interface())
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
	return resultSlice, e
}
