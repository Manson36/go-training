package main

import "fmt"

func main(){

	//for循环代替While
	var i int =1

	for {
		if i>10{
			break
		}
		fmt.Println("Hello World!")
		i++
	}
	fmt.Println("i =",i)

	//do While 的实现
	var j int = 1

	for{
		fmt.Println("Hello World!")
		j++
		if j>10{
			break
		}
	}

}