package collection

import (
	"reflect"
	"sort"
)

func (s RSet)SortAscBy(fields ...string) *RSet {

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

func (s RSet)SortDescBy(fields ...string) *RSet {

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
