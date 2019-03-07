package main

import "fmt"

func BubbleSort(arr *[5]int){
	fmt.Println("排序前arr=",*arr)

	temp := 0
	for i := 0; i< len(*arr)-1; i++{
		for j:=0; j < len(*arr)-1-i; j++{
			if (*arr)[j] > (*arr)[j+1]{
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
	fmt.Println("排序后arr=",*arr)
}

func main(){

	arr := [5]int {25,68,80,57,13}

	BubbleSort(&arr)

	fmt.Println("main arr=",arr)
}