package main

import "fmt"

//请编写一个函数，swap(n1 *int,n2 *int)，可以交换n1和n2 的值

func swap(n1 *int,n2 *int){

	t:= *n1
	*n1 = *n2
	*n2 = t
}
func main(){

	a:= 10
	b:= 2
	swap(&a,&b)
	fmt.Println("a =",a,"b =",b)
}