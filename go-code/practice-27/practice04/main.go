package main

import "fmt"

func main() {

	cities := make(map[string]string)

	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "天津"

	for k,v := range cities {
		fmt.Printf("k=%v, v=%v\n",k,v)
	}

	//遍历一个复杂的结构
	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string,3)
	studentMap["stu01"]["name"] = "Tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "北京长安街"

	studentMap["stu02"] = make(map[string]string,3)
	studentMap["stu02"]["name"] = "Mary"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "上海黄浦江"

	for k1,v1 := range studentMap {
		fmt.Println("k1=", k1)
		for k2,v2 := range v1 {
			fmt.Printf("\t k2 =%v,v2=%v\n",k2,v2)
		}
		fmt.Println()
	}
}