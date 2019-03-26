package main

import "fmt"

type MethodUtil struct {

}
func (mu MethodUtil)Print() {
	for i := 1; i <= 10; i++ {
		for j :=1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtil)Print2(m,n int) {
	for i := 1; i <= m; i++ {
		for j :=1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {

	var mu MethodUtil
	mu.Print()
	mu.Print2(3,2)
}