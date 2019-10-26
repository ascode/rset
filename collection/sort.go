package collection

import (
	"reflect"
	"sort"
)

// 按照给定的多个字段顺次降低优先级升序排序
func (s RSet)SortAscDowngradeBy(fields ...string) *RSet {

	theSet := &s

	sort.Slice(theSet.Set, func(i, j int) bool {
		switch theSet.Set[i][fields[0]].(type) {
		case int64:
			if theSet.Set[i][fields[0]].(int64) == theSet.Set[j][fields[0]].(int64) {
				return true
			}else {
				return  theSet.Set[i][fields[0]].(int64) < theSet.Set[j][fields[0]].(int64)
			}
		case string:
			return true
		default:
			return false
		}
	})
	return theSet
}

// 按照给定的多个字段顺次降低优先级降序排序
func (s RSet)SortDescDowngradeBy(fields ...string) *RSet {

	theSet := &s

	sort.Slice(theSet.Set, func(i, j int) bool {
		switch reflect.ValueOf(theSet.Set[i][fields[0]]).Interface().(type) {
		case int64:
			if theSet.Set[i][fields[0]].(int64) == theSet.Set[j][fields[0]].(int64) {
				return true
			}else {
				return  theSet.Set[i][fields[0]].(int64) > theSet.Set[j][fields[0]].(int64)
			}
		case string:
			return true
		default:
			return false
		}
	})
	return theSet
}

// 按照给定的多个字段竞争升序排序
func (s RSet)SortAscVieBy(fields ...string) *RSet {
	theSet := &s

	sort.Slice(theSet.Set, func(i, j int) bool {
		switch theSet.Set[i][fields[0]].(type) {
		case int64:
			var a, b int64
			if theSet.Set[i][fields[0]].(int64) > theSet.Set[i][fields[1]].(int64) {
				a = theSet.Set[i][fields[0]].(int64)
			}else {
				a = theSet.Set[i][fields[1]].(int64)
			}
			if theSet.Set[j][fields[0]].(int64) > theSet.Set[j][fields[1]].(int64) {
				b = theSet.Set[j][fields[0]].(int64)
			}else {
				b = theSet.Set[j][fields[1]].(int64)
			}
			return a < b
		case string:
			return true
		default:
			return false
		}
	})
	return theSet
}

// 按照给定的多个字段竞争降序排序
func (s RSet)SortDescVieBy(fields ...string) *RSet {
	theSet := &s

	sort.Slice(theSet.Set, func(i, j int) bool {
		switch reflect.ValueOf(theSet.Set[i][fields[0]]).Interface().(type) {
		case int64:
			var a, b int64
			if theSet.Set[i][fields[0]].(int64) > theSet.Set[i][fields[1]].(int64) {
				a = theSet.Set[i][fields[0]].(int64)
			}else {
				a = theSet.Set[i][fields[1]].(int64)
			}
			if theSet.Set[j][fields[0]].(int64) > theSet.Set[j][fields[1]].(int64) {
				b = theSet.Set[j][fields[0]].(int64)
			}else {
				b = theSet.Set[j][fields[1]].(int64)
			}
			return a > b
		case string:
			return true
		default:
			return false
		}
	})
	return theSet
}