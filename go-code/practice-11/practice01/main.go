package main

import "fmt"

//创建结构体变量时指定字段值
type Stu struct {
	Name string
	Age int
}

func main() {

	var stu1 Stu= Stu{"小明",19}
	stu2 := Stu{"小李",20}
	var stu3 = Stu{
		Name:"jack",
		Age: 21,
	}
	var stu4 = Stu{
		Age: 22,
		Name:"Mary",
	}
	fmt.Println(stu1,stu2,stu3,stu4)

	//方式2
	var stu5 *Stu =&Stu{"小明",19}
	stu6 := &Stu{"小李",20}
	var stu7 = &Stu{
		Name:"jack",
		Age: 21,
	}
	var stu8 = &Stu{
		Age: 22,
		Name:"Mary",
	}
	fmt.Println(*stu5,*stu6,*stu7,*stu8)
}