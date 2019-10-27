package collection

// TODO: - 集合差
func (s RSet) Minus(c []interface{}) {
	s.Lock()
	defer s.Unlock()
}

// TODO: - 和原始集合的集合差
func (s RSet) MinusOrigin(c []interface{}) {
	s.Lock()
	defer s.Unlock()
}