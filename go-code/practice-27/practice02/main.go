package main

import "fmt"

//二分查找
func BinaryFind(arr *[6]int, leftIndex int, rightleft int,findVal int){

	if leftIndex > rightleft {
		fmt.Println("找不到")
		return
	}
	middle := (leftIndex + rightleft)/2

	if (*arr)[middle] > findVal {
		BinaryFind(arr ,leftIndex,middle-1,findVal)
	}else if (*arr)[middle] < findVal {
		BinaryFind(arr ,middle+1,rightleft,findVal)
	}else {
		fmt.Println("找到了，下标为",middle)
	}
}

func main() {

	arr := [6]int{1,2,434,555,777,7676}

	BinaryFind(&arr,0,len(arr),434)
}