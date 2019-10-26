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
	fmt.Println("------------------------------")


	s := rset.Skip(1).Limit(2) // 留下第1个开始连续2个元素
	printSet(&s)
}

func printSet(rset *collection.RSet) {
	for _, obj := range rset.Set {
		fmt.Printf("%d,%s,%d,%-v\n", obj["Id"], obj["Name"], obj["Age"], obj["IsNewbie"])
	}
}
