package main

import (
	"fmt"
	"reflect"
)

//反射快速入门
func reflectTest01(b interface{}){

	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=",rTyp)

	rVal := reflect.ValueOf(b)
	n2 := 2 + rVal.Int()
	fmt.Println("n2 =",n2)

	fmt.Printf("rVal=%v,rVal type=%T",rVal,rVal)

	iV := rVal.Interface()
	num2 := iV.(int)
	fmt.Println("num2=",num2)
}

func reflectTest02(b interface{}){

	rTyp := reflect.TypeOf(b)
	fmt.Println("rTyp=",rTyp)

	rVal := reflect.ValueOf(b)

	iV := rVal.Interface()
	fmt.Printf("iV=%v,iV type =%T\n",iV,iV)

	stu,ok := iV.(Student)
	if ok {
		fmt.Printf("stu.name=%v",stu.Name)
	}
}

type Student struct {
	Name string
	Age int
}

type Monster struct {
	Name string
	Age int
}


func main() {

	var num int
	reflectTest01(num)

	stu := Student{
		Name:"Jack",
		Age:23,
	}
	reflectTest02(stu)
}