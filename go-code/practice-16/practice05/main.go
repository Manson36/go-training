package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//创建一个新文件，输入5句“hello,Manson"
	filePath := "d:/1/abc.txt"

	file,err := os.OpenFile(filePath,os.O_WRONLY| os.O_CREATE,0666)
	if err != nil {
		fmt.Println("open file err=",err)
		return
	}
	defer file.Close()
	//准备写入5句"Hello,Manson"
	str := "Hello,Manson\t\n"
	writer := bufio.NewWriter(file)
	for i:=0; i<5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}