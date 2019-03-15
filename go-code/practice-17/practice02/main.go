package main

import (
	"flag"
	"fmt"
)

func main() {

	//flag 包 解析命令行参数
	var user string
	var pwd string
	var host string
	var port string

	flag.StringVar(&user,"u","","用户名，默认为空")
	flag.StringVar(&pwd,"pwd","","密码，默认为空")
	flag.StringVar(&host,"h","localhost","主机名，默认为空")
	flag.StringVar(&port,"port","3306","端口号，默认为3306")

	//转换必须调用该方法
	flag.Parse()

	//输出结果
	fmt.Printf("user=%v,pwd=%v,host=%v,port=%v",
		user,pwd,host,port)
}