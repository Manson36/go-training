package main

import (
	"encoding/json"
	"fmt"
)

//反射场景的引出
type Monster struct {
	Name string `json:"monsterName"`
	Age int `json:"monsterAge"`
	Sal float64 `json:"monsterSal"`
	Sex string `json:"monsterSex"`
}

func main() {

	m := Monster{
		Name:"玉兔精",
		Age:255,
		Sal:23333.3,
		Sex:"女",
	}
	data,_ := json.Marshal(m)

	fmt.Println("json的结果为",string(data))
}