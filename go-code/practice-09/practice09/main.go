package main

import "fmt"

func BubbleSort(arr *[5]int) {
	count := 0
	flag := false

	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-1-i; j++ {
			count++
			if (*arr)[j] > (*arr)[j+1] {
				flag = true
				(*arr)[j] = (*arr)[j] ^ (*arr)[j + 1]
				(*arr)[j + 1] = (*arr)[j] ^ (*arr)[j + 1]
				(*arr)[j] = (*arr)[j] ^ (*arr)[j + 1]
			}
		}

		if !flag {
			break
		}

		flag = false
	}
	fmt.Println("count", count)
}

func main() {
	arr := [5]int{22, 1, 33, 56, 54454}
	BubbleSort(&arr)
	fmt.Println("arr=",arr)

}
