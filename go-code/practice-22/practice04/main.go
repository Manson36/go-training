package main

import (
	"fmt"
	"strings"
)

func main() {

	//字符串大小写的转换
	var orig = "How are you ?"
	var lower string
	var upper string

	fmt.Println(orig)
	lower = strings.ToLower(orig)
	fmt.Println(lower)
	upper = strings.ToUpper(orig)
	fmt.Println(upper)


}