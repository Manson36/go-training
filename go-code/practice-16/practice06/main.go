package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//打开一个存在的文件，追加"ABC,EGLISH"
	filePath := "d:/1/abc.txt"

	file,err := os.OpenFile(filePath,os.O_WRONLY| os.O_APPEND,0666)
	if err != nil {
		fmt.Println("open file err=",err)
		return
	}
	defer file.Close()
	//准备写入5句"Hello,Manson"
	str := "ABC,EGLISH\r\n"
	writer := bufio.NewWriter(file)
	for i:=0; i<5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}