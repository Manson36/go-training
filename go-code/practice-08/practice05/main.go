package main

import "fmt"

func main(){

	//map[string]map[string]string
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
}