package main

import "fmt"

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

func main() {

	r1 := Rect{Point{1,2},Point{22,11}}

	fmt.Printf("r1.leftUp.x的地址%q，r1.leftUp.y的地址%q，r1.rightDown.x的地址%q，r1.rightDown.y的地址%q\n",
		&r1.leftUp.x,&r1.leftUp.y,&r1.rightDown.x,&r1.rightDown.y)

	r2 := Rect2{&Point{10,20},&Point{220,110}}

	fmt.Printf("r2.leftUp 本身的地址%q，r2.rightDown 本身的地址%q\n",
		&r2.leftUp,&r2.rightDown)
	fmt.Printf("r2.leftUp 本身的地址%q，r2.rightDown 本身的地址%q\n",
		r2.leftUp,r2.rightDown)
}