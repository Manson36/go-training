package main

import "fmt"

func BubbleSort(arr *[5]int){

	temp:= 0
	for i:=0; i < len(*arr)-1; i++ {

		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
}

func main(){
	arr := [5]int{11,21,434,55,2}
	BubbleSort(&arr)
	fmt.Println("arr=",arr)
}