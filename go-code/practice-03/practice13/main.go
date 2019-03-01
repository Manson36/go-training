package main

import (
	"fmt"
	"github.com/go-training/go-code/practice-03/utils"
)

var age = test()

func test() int {

	fmt.Println("test()...")
	return 90
}
func init(){

	fmt.Println("init()...")
}
func main(){

	fmt.Println("main()...age =",age)
	fmt.Println("Age =",utils.Age,"Name =",utils.Name)

}