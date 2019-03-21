package main

import "fmt"

func main() {

	m := map[string]int{"a":1,"b":2}

	for i,v := range m {
		fmt.Println(i,v)
	}
	number := []int {1,2,3,4}
	for i,v := range number {
		fmt.Printf("第%d次，值为%d",i,v)
	}
}