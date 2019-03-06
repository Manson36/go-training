package main

import "fmt"

func main(){

	//map声明
	var a map[string]string
	a= make(map[string]string,10)

	a["no1"] = "宋江"
	a["no2"] = "吴用"
	a["no1"] = "武松"
	a["no3"] = "吴用"

	fmt.Println(a)

	//第二种方式
	var cities =make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "广州"

	fmt.Println(cities)

	//第三种方式
	heros := map[string]string{
		"hero1" : "宋江",
		"hero2" : "武松",
	}
	heros["hero3"] = "吴用"

	fmt.Println(heros)
}