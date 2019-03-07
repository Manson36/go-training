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

		if !flag { // 冒泡，插入，希尔，快速，桶，堆
			break
		}

		flag = false
	}
	fmt.Println("count", count)
} // (n -1 )(n - 1 - x) n^2 -1n + O(n2)

func main() {
	//arr := [5]int{22, 1, 33, 56, 54454}
	//BubbleSort(&arr)
	//fmt.Println("arr=",arr)

	test()
}

func test() {
	s := make([]int, 5)
	 fmt.Printf("cap=%v \n",cap(s))
	s = append(s, 10)
	s = append(s, 11)
	x := append(s, 12, 13)
	y := s
	y = append(y, 14)

	fmt.Println(s, x, y)
	fmt.Printf("cap=%v",cap(s))
}