package collection

import "reflect"

// 添加一个map[string]interface{}
func (s *RSet) Add(e interface{}) {
	s.Lock()
	defer s.Unlock()

	newElement := map[string]interface{}{}
	for i := 0; i < reflect.TypeOf(e).NumField(); i ++ {
		newElement[reflect.TypeOf(e).Field(i).Name] = reflect.ValueOf(e).Field(i).Interface()
	}
	s.Set = append(s.Set, newElement)
}
