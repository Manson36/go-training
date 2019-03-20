package main

import (
	"fmt"
	"strings"
)

func main() {

	//分割
	fmt.Printf("%q \n",strings.Split("a,b,c",","))
	fmt.Printf("%q \n",strings.Split("a boy a girl a dog a cat","a "))
	fmt.Printf("%q\n",strings.Split("xyz",""))

	//修剪
	fmt.Printf("%q\n",strings.Trim(" !!! Golang !!! ","! "))
	fmt.Printf("%q\n",strings.Trim(" !!! Golang !!! ","!"))

	fmt.Printf("%q\n",strings.TrimLeft(" !!! Golang !!! ","! "))
	fmt.Printf("%q\n",strings.TrimLeft(" !!! Golang !!! ","!"))

	fmt.Println(strings.TrimSpace(" \t \n 这是\t一句话 \n\t  "))
}