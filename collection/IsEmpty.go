package collection

// TODO: - 判断集合是不是空的
func (s RSet) IsEmpty() bool {
	s.Lock()
	defer s.Unlock()
	return len(s.Set) == 0
}
