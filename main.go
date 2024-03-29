package main

import (
	"fmt"
	"github.com/ascode/rset/collection"
	"github.com/ascode/rset/test_data"
)

func main() {

	// 新建一个RSet
	rset := collection.NewSet(test_data.StudentsGood, test_data.StudentsBad)

	printSet(rset)
	fmt.Println("---------------")
	rset.SortDescDowngradeBy("Id") 	// 集合按照排序字段顺次降级排序
	printSet(rset)
	fmt.Println("---------------")
	rset.SortAscDowngradeBy("Id") 	// 集合按照排序字段顺次升级排序
	printSet(rset)

	fmt.Println("")
	fmt.Println("----------------------------------------------")


	s := rset.SkipReturn(1).LimitReturn(2) // 留下第1个开始连续2个元素
	printSet(&s)

	fmt.Println("----------------------------------------------")

	std1 := test_data.Student{
		Id:5,
		Name:"小金李",
		Age:0,
		IsNewbie:true,
	}
	rset.Add(std1)	 // 添加一个元素
	printSet(rset)

	fmt.Println("----------------------------------------------")

	var newStudentCanInterface interface{} = []test_data.Student{}	// 新建一个装载结果的interface{}容器
	fmt.Println("fill前的newStudentCanInterface", newStudentCanInterface)

	rset.Fill(&newStudentCanInterface)								// 将rset填充到容器
	fmt.Println("fill之后的newStudentCanInterface", newStudentCanInterface)

	newStudentCan := newStudentCanInterface.([]test_data.Student)	// 将interface{}转换成其原始类型
	fmt.Println("强转对象之后的newStudentCan：", newStudentCan)

}

func printSet(rset *collection.RSet) {
	for _, obj := range rset.Set {
		fmt.Printf("%d,%s,%d,%-v\n", obj["Id"], obj["Name"], obj["Age"], obj["IsNewbie"])
	}
}
