package test_data

type Student struct {
	Name string
	Age int64
	IsNewbie bool
}

var Students []*Student

func init()  {
	std := Student{
		Name:"金飞",
		Age:18,
		IsNewbie:false,
	}
	Students = append(Students, &std)

	std = Student{
		Name:"李平",
		Age:12,
		IsNewbie:true,
	}

	Students = append(Students, &std)
}