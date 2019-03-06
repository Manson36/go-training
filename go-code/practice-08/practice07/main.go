package main

import "fmt"

func main(){

	//map的遍历
	var cities =make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "广州"

	for k,v := range cities{
		fmt.Printf("k=%v v=%v",k,v)
	}

	studentMap := make(map[string]map[string]string,3)

	studentMap["stu1"] = make(map[string]string)
	studentMap["stu1"]["name"] = "Tom"
	studentMap["stu1"]["sex"] = "男"
	studentMap["stu1"]["address"] = "北京"

	studentMap["stu2"] = make(map[string]string,3)
	studentMap["stu2"]["name"] = "Mary"
	studentMap["stu2"]["sex"] = "女"
	studentMap["stu2"]["address"] = "上海"

	fmt.Println(studentMap)
	fmt.Println(studentMap["stu2"])
	fmt.Println(studentMap["stu2"]["address"])

	for k1,v1 := range studentMap{
		fmt.Println("k1=",k1)
		for k2,v2 := range v1{
			fmt.Printf("\t k2=%v v2=%v\n",k2,v2)
		}
		fmt.Println()
	}
}