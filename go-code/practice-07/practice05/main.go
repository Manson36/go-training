package main

import "fmt"

func M(arr *[5]int){

		fmt.Println("arr排序前",*arr)
		temp :=0

		for i:=0; i< len(*arr)-1; i++{

			for j:=0; j < len(*arr)-1-i; j++{
				if (*arr)[j] > (*arr)[j+1]{
					temp = (*arr)[j]
					(*arr)[j] = (*arr)[j+1]
					(*arr)[j+1] = temp
				}
			}
		}
	fmt.Println("arr排序后",*arr)
	}

func main(){

	arr:= [5]int{22,33,11,55,6}

	M(&arr)

	fmt.Println("main arr",arr)
}