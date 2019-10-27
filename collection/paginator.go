package collection

// 分页
func (s RSet)Skip(n int64) RSet {
	s.Lock()
	defer s.Unlock()

	return RSet{
		Set:s.Set[n:],
	}
}

func (s RSet)Limit(n int64) RSet {
	s.Lock()
	s.Unlock()

	return RSet{
		Set:s.Set[:n],
	}
}