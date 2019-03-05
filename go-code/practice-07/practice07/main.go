package main

import "fmt"

func BinaryFind(arr *[6]int,leftIndex int,rightIndex int,findVal int){
		if leftIndex > rightIndex{
			fmt.Println("找不到")
			return
		}
		middle := (leftIndex + rightIndex)/2

		if (*arr)[middle] > findVal{
			BinaryFind(arr,leftIndex,middle-1,findVal)
		}else if (*arr)[middle] < findVal{
			BinaryFind(arr,middle+1,rightIndex,findVal)
		}else{
			fmt.Printf("找到了，下标%v",middle)
		}
	}

func main(){

	//二分法查找
	arr := [6]int{11,22,33,44,55,66}

	BinaryFind(&arr,0,len(arr)-1,-6)

}