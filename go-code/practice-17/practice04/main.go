package main

import (
	"encoding/json"
	"fmt"
)

//json 的反序列化

type Monster struct {
	Name string
	Age int
	Birthday string
	Sal float64
	Skill string
}

func unmarshalStruct()  {

	str := "{\"Name\":\"牛魔\",\"Age\":500,\"Birthday\":\"2111-22-33\",\"Sal\":33333.3,\"Skill\":\"小牛拳\"}"
	var monster Monster

	err := json.Unmarshal([]byte(str),&monster)
	if err != nil {
		fmt.Println("反序列化错误 err=\n",err)
	}
	fmt.Printf("反序列化结果，monster=%v,monster.Name =%v\n",monster,monster.Name)
}

//map 的反序列化
func unmarshalMap() {
	str := "{\"address\":\"火焰山\",\"age\":19,\"name\":\"红孩儿\"}"
	var a map[string]interface{}

	err := json.Unmarshal([]byte(str),&a)
	if err != nil {
		fmt.Println("反序列化错误 err=\n",err)
	}
	fmt.Printf("反序列化结果a =%v\n",a)
}

//对切片的反序列化
func unmarshalSlice() {
	str :="[{\"address\":\"北京\",\"age\":\"22\",\"name\":\"jack\"}," +
		"{\"address\":[\"夏威夷\",\"波士顿\"],\"age\":\"265\",\"name\":\"Sam\"}]"

	var slice []map[string]interface{}

	err := json.Unmarshal([]byte(str),&slice)
	if err != nil {
		fmt.Println("反序列化错误 err=\n",err)
	}
	fmt.Printf("反序列化结果 slice=%v\n",slice)
}

func main() {

	unmarshalMap()
	unmarshalSlice()
	unmarshalStruct()
}