package main

import (
	"errors"
	"fmt"
)

//自定义错误
	func readConf(name string )(err error){

		if name == "config.ini"{
			//读取
			return nil
		}else {
			//返回自定义错误
			return errors.New("读取文件错误")
		}
	}

	 func test02(){

	 	err:= readConf("config2.ini")
	 	if err !=nil {
	 		//输出错误并终止程序
	 		panic(err)
		}
	 	fmt.Println("test02继续执行。。。")
	 }

func main(){

	test02()
	fmt.Println("main()")
}