package main

import (
	"fmt"
	"sort"
)

func main() {

	map1 := make(map[int]int,10)
	map1[10] = 100
	map1[1] = 13
	map1[3] = 33
	map1[6] = 123

	fmt.Println(map1)

	var keys []int
	for k,_ := range map1 {
		keys = append(keys,k)
	}
	sort.Ints(keys)
	fmt.Println(keys)

	for _,k :=range keys {
		fmt.Printf("map1[%v]=%v\n",k,map1[k])
	}
}