package main

import "fmt"

func main() {

	ageArr := []int{7,9,2,4,1}
	f1(ageArr...)
}

func f1(arr ...int) {
	f2(arr...)
	fmt.Println("")
	f3(arr)
}

func f2(arr ...int) {
	for _,char := range arr {
		fmt.Printf("%d",char)
	}
}

func f3(arr []int) {
	for _,char := range arr {
		fmt.Printf("%d",char)
	}
}