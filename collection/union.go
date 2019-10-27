package collection

// TODO: - 集合并
func (s RSet) Union(c []interface{}) {
	// 留着锁试试感觉
	s.Lock()
	defer s.Unlock()


}


// TODO: - 和原始集合的集合并
func (s RSet) MinusUnion(c []interface{}) {

}