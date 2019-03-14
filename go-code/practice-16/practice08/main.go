package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	//读出原来的内容，并追加5句"hello,Beijing."
	failPath := "d:/1/test.txt"

	file,err := os.OpenFile(failPath,os.O_RDWR | os.O_APPEND,0666)
	if err != nil {
		fmt.Println("read fail err=",err)
		return
	}
	defer file.Close()

	//先读取内容到终端
	reader := bufio.NewReader(file)
	for {
		str,err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	str := "hello,Beijing\r\n"
	writer := bufio.NewWriter(file)
	for i:=0 ;i<5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()

}