package rset

import (
	"github.com/ascode/rset/collection"
	"github.com/ascode/rset/test_data"
	_ "github.com/ascode/rset/test_data"
)

func main() {
	collection.NewSet(test_data.Students)
}
