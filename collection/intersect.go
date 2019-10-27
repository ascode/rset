package collection

// TODO: - 集合交
func (s RSet) Intersect(c []interface{}) {
	s.Lock()
	defer s.Unlock()

}

// TODO: - 和原始集合的集合交
func (s RSet) MinusIntersect(c []interface{}) {
	s.Lock()
	defer s.Unlock()
}