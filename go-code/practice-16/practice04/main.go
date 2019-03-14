package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	file := "d:/1/test.txt"
	content,err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("read file err=",err)
	}

	fmt.Println(string(content))
}
