package collection

import (
	"fmt"
	"reflect"
)

func NewSet(elements ...[]*interface{}) *RSet {
	//returnSet := []interface{}{}
	for _, element := range elements  {
		//returnSet = append(returnSet, element)
		for _, item := range element {
			getType := reflect.TypeOf(item)
			for i := 0; i < getType.NumField(); i ++ {
				fmt.Println(getType.Field(i).Name)
				fmt.Println(getType.Field(i).Type)
			}
		}
	}
	return nil
}

func Abc(){

}