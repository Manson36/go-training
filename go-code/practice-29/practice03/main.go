package main

import (
	"fmt"
	"github.com/go-training/go-code/practice-29/practice03/practice-"
)

func main() {

	var stu = utils.NewStudent("tom",99.8)

	fmt.Println(*stu)
	fmt.Println("name=",stu.Name,"score=",stu.GetScore())

}