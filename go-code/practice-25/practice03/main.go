package main

import (
	"fmt"
	"time"
)

func main() {

	i,p := a()
	fmt.Println("return:",i,p,time.Now())
}
func a() (i int, p *int) {

	defer func(i int) {
		fmt.Println("defer3:",i,&i,time.Now())
	}(i)

	defer fmt.Println("defer2:",i,&i,time.Now())

	defer func() {
		fmt.Println("defer1:",i,&i,time.Now())
	}()

	i++
	func() {
		fmt.Println("print1:",i,&i,time.Now())
	}()

	fmt.Println("print2:",i,&i,time.Now())

	return i,&i
}