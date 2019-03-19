package main

import (
	"fmt"
	"reflect"
)

//反射的最佳实践
type Monster struct {
	Name string `json:"name"`
	Age int `json:"monster_name"`
	Score float64 `json:"成绩"`
	Sex string
}

//方法：返回两个数的和
func(s Monster)GetSum(n1,n2 int)int {

	return n1 + n2
}

//方法：接受四个值，给s 赋值
func(s Monster)Set(name string, age int, score float64, sex string){

	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

//方法：显示s的值
func(s Monster)Print(){

	fmt.Println("----Start----")
	fmt.Println(s)
	fmt.Println("----End----")

}

func TestStruct(a interface{}) {

	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()

	if kd != reflect.Struct {
		fmt.Println("except struct")
		return
	}
	num := val.NumField()

	fmt.Printf("struct has %d field\n", num)

	//遍历结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field%d:的值为%v\n", i, val.Field(i))

		//获取struct的标签
		tagStruct := typ.Field(i).Tag.Get("json")

		if tagStruct != "" {
			fmt.Printf("Field%d:的值为%v\n", i, tagStruct)
		}
	}

	//获取该结构体有多少个方法
	numOfMethod := val.NumMethod()
	fmt.Printf("struct has %d methods\n",numOfMethod)

	val.Method(1).Call(nil)//获取到第二个方法，方法排序是按照方法名的ASCII码

	var params []reflect.Value
	params = append(params,reflect.ValueOf(10))
	params = append(params,reflect.ValueOf(40))

	res := val.Method(0).Call(params)
	fmt.Println("res=",res[0].Int())
}


func main() {

	var a Monster = Monster{
		Name:"黄鼠狼精",
		Age:233,
		Score:233.3,
	}
	TestStruct(a)
	}