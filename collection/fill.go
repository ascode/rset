package collection

import (
	"reflect"
)

// 填充指定的容器（对象）
func (s RSet) Fill(can *interface{}) {

	canType := reflect.ValueOf(*can).Type().Elem() // reflect.TypeOf(can)
	canValue := reflect.ValueOf(*can)

	//fmt.Println("1", canType)
	for i := 0; i < len(s.Set); i ++ {
		canItemTempReflectValue := reflect.New(canType)
		//fmt.Println("4", )
		//canItemTemp :=
		canItemTempType := reflect.TypeOf(canItemTempReflectValue.Interface()).Elem()
		//fmt.Println("5", canItemTempType)
		//canItemTempValuePrt := reflect.ValueOf(canItemTempReflectValue.Interface())
		//fmt.Println("6", canItemTempValuePrt)
		canItemTempValue := reflect.ValueOf(canItemTempReflectValue.Interface()).Elem()
		//fmt.Println("7", canItemTempValue)
		for j := 0; j < canItemTempType.NumField(); j ++ {
			fieldName := canItemTempType.Field(j).Name

			//fmt.Println("2", canItemTempValue.Field(j))
			//fmt.Println("21", canItemTempType.Field(j).Name)
			//fmt.Println("22", s.Set[i][fieldName])

			switch s.Set[i][fieldName].(type) {
			case int64:
				canItemTempValue.Field(j).SetInt(s.Set[i][fieldName].(int64))
			case string:
				canItemTempValue.Field(j).SetString(s.Set[i][fieldName].(string))
			case float64:
				canItemTempValue.Field(j).SetFloat(s.Set[i][fieldName].(float64))
			case bool:
				canItemTempValue.Field(j).SetBool(s.Set[i][fieldName].(bool))
			default:

			}

			//fmt.Println("8", canItemTempValue)
		}
		canValue = reflect.Append(canValue, canItemTempValue)

		//fmt.Println(reflect.ValueOf(can).Elem())
		//fmt.Println("result 1 ", canValue)
	}
	*can = canValue.Interface()
}