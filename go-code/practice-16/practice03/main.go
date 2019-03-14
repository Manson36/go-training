package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main () {

	//打开文件
	//file 叫对象 指针 文件句柄
	file,err := os.Open("d:/1/test.txt")
	if err != nil {
		fmt.Println("open file err=",err)
	}
	defer file.Close()

	//创建一个带缓冲的*Reader
	reader := bufio.NewReader(file)
	//循环读取文件的内容
	for {
		str,err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束")
}