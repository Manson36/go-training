package main

import "fmt"

func main() {

	//channel的读写练习
	var intChan chan int
	intChan = make(chan int ,3)
	intChan <-10
	intChan <-20
	intChan <-30

	num1 := <-intChan
	num2 := <-intChan
	num3 := <-intChan

	fmt.Printf("num1=%v,num1=%v,num1=%v\n",num1,num2,num3)

	var mapChan chan map[string]string
	mapChan =  make(chan map[string]string,4)

	m1 := make(map[string]string,20)
	m1["city1"] = "北京"
	m1["city2"] = "上海"

	m2 := make(map[string]string,20)
	m1["hero1"] = "武松"
	m1["hero2"] = "鲁智深"

	mapChan <- m1
	mapChan <- m2

	m11 := <- mapChan
	m22 := <- mapChan

	fmt.Printf("m11=%v,m22=%v",m11,m22)

	//存放任意类型的数据
	var allChan chan interface{}
	type Cat struct{
		Name string
		Age int
	}
	allChan = make(chan interface{},10)

	cat1 := Cat{Name:"tom",Age:2}
	cat2 := Cat{Name:"Jack",Age:3}
	allChan <- cat1
	allChan <- cat2
	allChan <- 10
	allChan <- "jack"

	p1 := <-allChan
	p2 := <-allChan
	p3 := <-allChan
	p4 := <-allChan

	fmt.Println(p1,p2,p3,p4)
}