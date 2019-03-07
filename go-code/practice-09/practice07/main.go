package main

import "fmt"

func main() {

	type Cat struct{
		Name string
		Age int
		Color string
		Hobby string
	}

	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	cat1.Hobby = "吃鱼"

	fmt.Println("猫的属性如下")
	fmt.Println("name=",cat1.Name)
	fmt.Println("Age=",cat1.Age)
	fmt.Println("Color=",cat1.Color)
	fmt.Println("Hobby=",cat1.Hobby)

}