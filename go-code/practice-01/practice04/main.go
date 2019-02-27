package main

import "fmt"

func Test(char byte) byte{
	return char+1
}
func main(){
	//swith 语句的基本使用2
	var key byte
	fmt.Println("请输入一个字符 a,b,c,d")
	fmt.Scanf("%c",&key)

	switch Test(key)+1 {
	case 'a':
		fmt.Println("周一，猴子穿新衣")
	case 'b':
		fmt.Println("周二，猴子当小二")
	case 'c':
		fmt.Println("周三，猴子爬雪山")
	default:
		fmt.Println("输入有误")
	}
}