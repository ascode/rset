package collection

// TODO: - 相对于 full 的补集，结果存入 s
func (s RSet) Complement(fullSet []interface{}) {
	s.Lock()
	defer s.Unlock()

}

// TODO: - 和原始集合的集合补
func (s RSet) MinusComplement(c []interface{}) {
	s.Lock()
	defer s.Unlock()
}