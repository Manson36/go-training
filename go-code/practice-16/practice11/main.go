package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount int
	NumCount int
	SpaceCount int
	OtherCount int

}

func main() {

	fileName := "d:/1/abc.txt"
	file,err := os.Open(fileName)
	if err != nil {
		fmt.Println("open fail err=",err)
	}
	defer file.Close()

	var count CharCount
	reader := bufio.NewReader(file)

	for {
		str,err := reader.ReadString('\n')
		if err == io.EOF{
			break
		}
	for _,v :=range str {
		switch {
		case v>'a'&& v < 'z':
			fallthrough
		case v>'A'&& v < 'Z':
			count.ChCount++
		case v>='0'&& v <= '9':
			count.NumCount++
		case v == ' ' || v == '\t':
			count.SpaceCount++
		default:
			count.OtherCount++
		}

	}
	fmt.Printf("字符的个数为%v，数字的个数为%v，空格的个数为%v ，其他的个数为%v",
		count.ChCount,count.NumCount,count.SpaceCount,count.OtherCount)
	}
}