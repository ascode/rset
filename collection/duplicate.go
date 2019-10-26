package collection


// 创建副本
func (s RSet) Duplicate() *RSet {
	s.Lock()
	defer s.Unlock()
	cp := &RSet{
		Set: []map[string]interface{}{},
	}
	for _, e := range s.Set {
		cp.Set = append(cp.Set, e)
	}
	return cp
}