package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//打开一个存在的文件，将原来的内容覆盖为"你好，中国"
	filePath := "d:/1/abc.txt"

	file,err := os.OpenFile(filePath,os.O_WRONLY| os.O_TRUNC,0666)
	if err != nil {
		fmt.Println("open file err=",err)
		return
	}
	defer file.Close()
	//准备写入5句"Hello,Manson"
	str := "你好，中国\r\n"
	writer := bufio.NewWriter(file)
	for i:=0; i<5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}