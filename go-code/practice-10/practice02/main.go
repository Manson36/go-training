package main

import "fmt"

type Person struct {

	name string
}

func (p Person)test(){

	fmt.Println("test()=",p.name)
}

func (p Person) speak (){
	fmt.Println(p.name,"is a goodman")
}

func (p Person) jisuan (){

	res := 0
	for i:=1; i<= 1000; i++{
		res+= i
	}
	fmt.Println(p.name,"jisuan jieguo",res )
}

func (p Person) jisuan2 (n int ){

	res := 0
	for i:=1; i<= n; i++{
		res+= i
	}
	fmt.Println(p.name,"jisuan jieguo",res )
}

func (p Person) Getsum(n1 int, n2 int )int{
	return n1+n2
}

func main() {

	var p Person
	p.name = "Tom"
	p.test()
	p.speak()
	p.jisuan()
	p.jisuan2(20)
	res := p.Getsum(12,21)
	fmt.Println(res )
}