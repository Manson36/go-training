package main

import "fmt"

func main(){

	//map的增删改查
	var cities =make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "广州"

	fmt.Println(cities)

	cities["no2"] = "上海~"
	fmt.Println(cities)

	delete(cities,"no2")
	fmt.Println(cities)

	delete(cities,"no4")
	fmt.Println(cities)

	cities =make(map[string]string)
	fmt.Println(cities)

	val,ok := cities["no2"]
	if ok{
		fmt.Printf("有no2 值为%v",val)
	}else {
		fmt.Println("没有no2")
	}

}