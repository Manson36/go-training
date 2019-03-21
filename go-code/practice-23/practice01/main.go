package main

import "fmt"

func main() {

	a := 13
	if a > 20 {
		fmt.Println("a大于20")
	}else if a < 10 {
		fmt.Println("a小于10")
	}else if a == 11 {
		fmt.Println("a等于11")
	}else {
		fmt.Println("a大于10")
		fmt.Println("a小于20")
		fmt.Println("a不等于11")
	}
	fmt.Println("a的值是",a)
}