package collection

import (
	"reflect"
)

func NewSet(elementGroup ...interface{}) *RSet {

	returnSet := []map[string]interface{}{}

	for _, elements := range elementGroup {
		for i := 0; i < reflect.ValueOf(elements).Len(); i ++ {
			getType := reflect.TypeOf(elements).Elem()
			set := map[string]interface{}{}
			for j := 0; j < getType.NumField(); j ++ {
				set[getType.Field(j).Name] = reflect.ValueOf(elements).Index(i).Field(j).Interface()
			}
			returnSet = append(returnSet, set)
		}
	}

	return &RSet{
		Set:returnSet,
	}
}

func Abc(){

}