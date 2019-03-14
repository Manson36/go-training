package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//拷贝文件
func copyFile(dstFileName string,srcFileName string)(wtitten int64,err error){
	 	srcFile,err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open file err=", err)
	}
	 	defer srcFile.Close()

	 	reader := bufio.NewReader(srcFile)

	 	dstFile,err := os.OpenFile(dstFileName,os.O_WRONLY | os.O_APPEND,0666)
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}
	 	writer := bufio.NewWriter(dstFile)
	 	defer dstFile.Close()
	 	return io.Copy(writer,reader)
	}
func main () {

	srcFile := "d:/1/test.txt"
	dstFile := "d:/1/abc.txt"
	_,err := copyFile(dstFile,srcFile)
	if err == nil {
		fmt.Println("拷贝完成\n")
	}else {
		fmt.Println("拷贝失败 err=",err)
	}
}
