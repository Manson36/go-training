package main

import "fmt"

type MethodUtil struct{

}

func (mu MethodUtil)Print() {
	for i:=1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtil)Print2(m int, n int ) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtil) JudgeNum(num int) {
		if num%2==0{
			fmt.Println(num,"shi oushu ")
		}else {
			fmt.Println(num,"shi ji shu ")
		}
	}


func main() {
	var mu MethodUtil
	mu.Print()
	mu.Print2(3,3)
	mu.JudgeNum(34)
	
}