package collection

// 返回分页数据
func (s RSet)SkipReturn(n int64) RSet {
	s.Lock()
	defer s.Unlock()

	return RSet{
		Set:s.Set[n:],
	}
}

func (s RSet)LimitReturn(n int64) RSet {
	s.Lock()
	s.Unlock()

	return RSet{
		Set:s.Set[:n],
	}
}

// 自己砍自己
func (s RSet)Skip(n int64) {
	s.Lock()
	defer s.Unlock()
	(&s).Set = s.Set[n:]
}

func (s RSet)Limit(n int64) {
	s.Lock()
	defer s.Unlock()

	(&s).Set = s.Set[:n]
}