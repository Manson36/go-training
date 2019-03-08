package main

import "fmt"

type box struct {
	len float64
	width float64
	height float64
}
func (box *box)getVolumn() float64{
	return box.len*box.width*box.height
}

func main(){

	var box box
	box.height = 1.2
	box.width =33
	box.len = 33
	volumn := box.getVolumn()
	fmt.Printf("体积为%.2f",volumn)
}