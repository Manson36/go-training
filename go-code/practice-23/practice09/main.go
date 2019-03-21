package main

import "fmt"

func pipe(ff func()int)	int {
	return ff()
}

type FormaFunc func(s string,x,y int)string

func format(ff FormaFunc,s string ,x,y int)string {
	return ff(s,x,y)
}

func main() {

	s1 := pipe(func() int {
		return 100
	})
	s2 := format(func(s string, x, y int) string {
		return  fmt.Sprintf(s,x,y)
	},"%d,%d",10,20)
	fmt.Println(s1,s2)
}