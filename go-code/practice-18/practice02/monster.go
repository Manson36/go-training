package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {

	Name string
	Age int
	Skill string
}

//给Monster绑定方法store，将一个Monster序列化保存到文件中
func (this *Monster)Store ()bool {

	data,err := json.Marshal(this)
	if err != nil {
		fmt.Println("err =",err)
		return false
	}
	//保存到文件
	filePath := "d:/1/monster.ser"
	err =ioutil.WriteFile(filePath,data,0666)
	if err != nil {
		fmt.Println("write fail err=",err)
		return false
	}
	return true
}

//给Monster 绑定方法 Restore，将文件读取出来，并反序列化结果
func (this *Monster)Restore ()bool {

	filePath := "d:/1/monster.ser"
	data,err :=ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read fail err=",err)
		return false
	}

	err = json.Unmarshal(data,this)
	if err != nil {
		fmt.Println("Unmarshal err =",err)
		return false
	}
	return true
}