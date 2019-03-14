package main

import (
	"fmt"
	"os"
)

func main () {

	//打开文件
	//file 叫对象 指针 文件句柄
	file,err := os.Open("d:/1/test.txt")
	if err != nil {
		fmt.Println("open file err=",err)
	}
	fmt.Println("file=",file)

	err = file.Close()
	if err != nil {
		fmt.Println("close file err=",err)
	}
}