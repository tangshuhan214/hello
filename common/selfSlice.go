package common

// active object对象
type Service struct {
	channel chan interface{} `desc:"即将加入到数据slice的数据"`
	data    []interface{}    `desc:"数据slice"`
}

// 新建一个size大小缓存的active object对象
func NewService(size int, done func()) *Service {
	s := &Service{
		channel: make(chan interface{}, size),
		data:    make([]interface{}, 0),
	}

	go func() {
		s.schedule()
		done()
	}()
	return s
}

// 把管道中的数据append到slice中
func (s *Service) schedule() {
	for v := range s.channel {
		s.data = append(s.data, v)
	}
}

// 增加一个值
func (s *Service) Add(v interface{}) {
	s.channel <- v
}

// 管道使用完关闭
func (s *Service) Close() {
	close(s.channel)
}

// 返回slice
func (s *Service) Slice() []interface{} {
	return s.data
}
