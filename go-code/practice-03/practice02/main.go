package main

import "fmt"

	//求出斐波那契数 1,1,2,3,5,8,13

	func fbn(n int)int {

		if(n == 1||n == 2){
			return 1
		}else{
			return fbn(n-1)+fbn(n-2)
		}
	}
func main(){

	res:=fbn(3)
	fmt.Println("res=",res)
	fmt.Println("res=",fbn(3))
	fmt.Println("res=",fbn(4))
	fmt.Println("res=",fbn(5))
	fmt.Println("res=",fbn(6))
	}