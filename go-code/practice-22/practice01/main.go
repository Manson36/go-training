package main

import (
	"fmt"
	"strconv"
)

func main() {

	orig := "233"
	fmt.Printf("orig的当前类型是%T，且操作系统是%d位。\n",orig,strconv.IntSize)

	num,err := strconv.Atoi(orig)
	fmt.Printf("num 的当前类型是%T，且数值是%d\n",num,num)
	fmt.Print("%s \n",err)
	num += 5
	newS := strconv.Itoa(num)
	fmt.Printf("%s \n",newS)
}