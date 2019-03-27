package main

import (
	"fmt"
	"github.com/go-training/go-code/practice-11/practice03"
)

func main() {

	p:= model.Newperson("smith")
	p.SetAge(18)
	p.SetSal(4000)
	fmt.Println(p)
	fmt.Println(p.Name,"age = ",p.GetAge(),"sal =",p.GetSal())
}