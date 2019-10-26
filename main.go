package main

import (
	"fmt"
	"github.com/ascode/rset/collection"
	"github.com/ascode/rset/test_data"
)

func main() {
	rset := collection.NewSet(test_data.StudentsGood, test_data.StudentsBad)
	fmt.Println(rset.Set)

	printSet(rset)
	fmt.Println("---------------")
	rset.SortDescDowngradeBy("Id")
	printSet(rset)
	fmt.Println("---------------")
	rset.SortAscDowngradeBy("Id")
	printSet(rset)

}

func printSet(rset *collection.RSet) {
	for _, obj := range rset.Set {
		fmt.Printf("%d,%s,%d,%-v\n", obj["Id"], obj["Name"], obj["Age"], obj["IsNewbie"])
	}
}
