package collection

// TODO: - 移除元素
func (s RSet) Remove(e interface{}) {
	s.Lock()
	defer s.Unlock()


}

// TODO: - 根据条件移除元素
func (s RSet) RemoveByCondition(condi interface{}) {

}

