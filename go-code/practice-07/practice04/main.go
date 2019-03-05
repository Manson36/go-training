package main

import "fmt"

//编写一个fbn(n int)函数
	func fbn(n int)([]uint64){

		fbnSlice := make([]uint64,n)
		fbnSlice[0] = 1
		fbnSlice[1] = 1

		for i:= 2; i<n ; i++{
			fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
		}
		return fbnSlice
	}

func main(){

	fbnSlice := fbn(20)
	fmt.Println("fbnSlice",fbnSlice)
}