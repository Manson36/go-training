package main

import "fmt"

//冒泡法
func bubbleSort(arr *[5]int) {
	count := 0
	flag := true

	for i := 0; i <len(*arr)-1; i++ {

		for j := 0; j < len(*arr)-1-i; j++ {
			count++
			if (*arr)[j] > (*arr)[j+1] {
				flag = false
				(*arr)[j],(*arr)[j+1] = (*arr)[j+1],(*arr)[j]
			}
		}
		if flag {
			break
		}
		flag = true
	}
	fmt.Println("count=",count)
}

func main() {

	arr := [5]int{22,3,55,7777,2}
	bubbleSort(&arr)
	fmt.Println("arr=",arr)

}