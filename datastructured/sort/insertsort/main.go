package main

import "fmt"

func InsertSort(arr *[5]int) {

	for i:= 1; i < len(arr); i++ {

	//完成第一次，给第二个元素找到位置并插入
	insertVal := arr[i]
	insertIndex := i - 1

	//从大到小排列
	for insertIndex >=0 && insertVal > arr[insertIndex] {
		arr[insertIndex + 1] = arr[insertIndex]
		insertIndex--
	}

	//插入
	if insertIndex + 1 != i {
		arr[insertIndex + 1] = insertVal
	}

	fmt.Printf("第%d次插入后%v\n", i, *arr)
	}
}

func main() {
	arr := [5]int{2, 0, 13, 56, 43}
	fmt.Println("排序前", arr)

	InsertSort(&arr)

	fmt.Println("排序后", arr)
}
