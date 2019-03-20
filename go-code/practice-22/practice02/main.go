package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "The quick brown fox jumps over the lazy dog 中文"
	strSli := strings.Fields(str)
	fmt.Println(strSli)

	for _,val := range strSli {
		fmt.Println(val)
	}
	fmt.Println()
	str2 := strings.Join(strSli,";")
	fmt.Println(str2)
}