package model

import "fmt"

type person struct {
	Name string
	age int
	sal float64
}
//写一个工厂模式的
func NewPerson(name string) *person {
	return &person{
		Name:name,
	}
}

func (p *person)SetAge(age int) {
	if age > 0 && age <150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确")
	}
}
func (p *person)Getage ()int {
	return p.age
}

func (p *person)SetSal(sal float64) {
	if sal > 3000 && sal <300000 {
		p.sal = sal
	} else {
		fmt.Println("薪水范围不正确")
	}
}
func (p *person)GetSal ()float64 {
	return p.sal
}