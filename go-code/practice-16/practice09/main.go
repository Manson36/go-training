package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	filePath1 := "d:/1/test.txt"
	filePath2 := "d:/2/kkk.txt"

	data, err := ioutil.ReadFile(filePath1)
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	err = ioutil.WriteFile(filePath2,data,0666)
	if err != nil {
		fmt.Println("write fail err=",err)
	}
}