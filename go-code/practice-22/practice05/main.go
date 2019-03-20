package main

import "fmt"

func main() {

	//打印空心金字塔
	var totalLevel =20

	for i := 1; i <= totalLevel; i++ {

		for j := 1; j <= totalLevel - i; j++ {

			fmt.Print(" ")
		}
		for k :=1; k <= 2 * i - 1; k ++ {
			if k == 1 || k == 2 * i -1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}