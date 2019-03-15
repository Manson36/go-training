package main

import (
	"encoding/json"
	"fmt"
)

//json 的序列化

	//结构体的序列化
	type Monster struct {
		Name string
		Age int
		Birthday string
		Sal float64
		Skill string
	}

func testStruct()  {
	//演示
	monster := Monster{
		Name:"牛魔",
		Age:500,
		Birthday:"2111-22-33",
		Sal:33333.3,
		Skill:"小牛拳",
	}
	//monster 序列化
	data,err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化错误 err=\n",err)
	}
	fmt.Println("struct 序列化结果\n",string(data))
}

	//map 序列化
	func testMap() {
	var a map[string]interface{}
	a = make(map[string]interface{})

	a["name"] = "红孩儿"
	a["age"] = 19
	a["address"] = "火焰山"

	data,err := json.Marshal(a)
	if err != nil {
	fmt.Println("序列化错误 err=\n",err)
	}
	fmt.Println("map 序列化结果\n",string(data))
	}

	//对切片序列化
func testSlice() {

	var slice []map[string]interface{}
	var m1 map[string]interface{}

	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = "22"
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}

	m2 = make(map[string]interface{})
	m2["name"] = "Sam"
	m2["age"] = "265"
	m2["address"] = [2]string{"夏威夷","波士顿"}
	slice = append(slice, m2)

	data,err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化错误 err=\n",err)
	}
	fmt.Println("切片 序列化结果\n",string(data))
}

func main() {

	testMap()
	testSlice()
	testStruct()
}