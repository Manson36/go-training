package main

import "fmt"

func SelectSort(arr *[6]int) {
	//注意：标准访问（*arr）[1]等价于arr[1]

	//1.先易后难，先死后活：做arr[0]为最大值，与后面比较

	for j := 0; j <len(arr) -1; j++ {

		max := arr[j]
		maxIndex := j
		//2.遍历后面1-[len(arr)-1],比较
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] {
				max = arr[i]
				maxIndex = i
			}
		}
		//交换
		if max != arr[j] {
			arr[maxIndex], arr[j] = arr[j], arr[maxIndex]
		}

		fmt.Printf("第%d次 %v	\n", j +1, *arr)
	}

}

func main() {
	arr := [6]int{12, 22, 3, 55, 23, 88}

	SelectSort(&arr)
}
