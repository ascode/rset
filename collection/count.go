package collection

// TODO: - 返回元素数量
func (s RSet)Count() int64 {
	s.Lock()
	defer s.Unlock()
	return int64(len(s.Set))
}