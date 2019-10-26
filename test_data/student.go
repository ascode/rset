package test_data

type Student struct {
	Id int64
	Name string
	Age int64
	IsNewbie bool
}

var StudentsGood []Student
var StudentsBad []Student

func init()  {
	std := Student{
		Id:1,
		Name:"金飞",
		Age:18,
		IsNewbie:false,
	}
	StudentsGood = append(StudentsGood, std)

	std = Student{
		Id:2,
		Name:"李平",
		Age:12,
		IsNewbie:true,
	}

	StudentsGood = append(StudentsGood, std)


	std = Student{
		Id:3,
		Name:"金飞1",
		Age:18,
		IsNewbie:false,
	}
	StudentsBad = append(StudentsBad, std)

	std = Student{
		Id:4,
		Name:"李平1",
		Age:12,
		IsNewbie:true,
	}

	StudentsBad = append(StudentsBad, std)
}