package main

import (
	"fmt"
	"github.com/go-training/go-code/practice-29/practice07/model"
)

func main() {

	account := model.NewAccount("js1111","111111",40)

	if account != nil {
		fmt.Println("创建成功=",account)
	} else {
		fmt.Println("创建失败")
	}
}