package main

import "fmt"

func getSwquence() func() int {
	i :=0
	return func() int {
		i++
		return i
	}
}

func main() {

	nextNumber := getSwquence()

	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	nextNumber1 := getSwquence()

	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

}