package main

import "fmt"

func isOdd(v int)bool {
	if v%2 == 0 {
		return false
	}
	return true
}
func isEven(v int)bool {
	if v%2 == 0 {
		return true
	}
	return false
}

type boolFunc func(int) bool

func filter(slice []int,f boolFunc)[]int {
	var result []int
	for _,value := range slice{
		if f(value) {
			result = append(result,value)
		}
	}
	return result
}

func main() {

	slice := []int{2,4,5,6,8,1}
	fmt.Println("slice=",slice)
	odd := filter(slice,isOdd)
	fmt.Println("odd:",odd)
	even := filter(slice,isEven)
	fmt.Println("even:",even)
}