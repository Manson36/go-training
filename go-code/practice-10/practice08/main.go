package main

import "fmt"

type Student struct {
	name string
	gender string
	age int
	id int
	score float64
}

func (Student Student)say() string{
	infostr := fmt.Sprintf("student的信息 name=[%v] gender=[%v]age=[%v]id=[%v]score=[%v]",
		Student.name,Student.gender,Student.age,Student.id,Student.score)
	return infostr
}

func main() {
	var stu = Student{
		name:"Tom",
		gender:"male",
		age:18,
		id:1000,
		score:99.4,
	}
	fmt.Println(stu.say())
}