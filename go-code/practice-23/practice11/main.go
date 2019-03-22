package main

import "fmt"

func main() {

	age := ageMinOrMax("min",1,3,4,2)
	fmt.Println("最小的数是",age )

	ageArr := []int{7,8,4,5,1}
	age = ageMinOrMax("max",ageArr...)
	fmt.Println("最大的数是",age)
}

func ageMinOrMax(m string, a ...int)int {
	if len(a) == 0 {
		return 0
	}
	if m== "max" {
		max := a[0]
		for _,v := range a {
			if v > max {
				max =v
			}
		}
		return max
	}else if m == "min" {
		min := a[0]
		for _,v :=range a {
			if v < min {
				min = v
			}
		}
		return min
	}else {
		e := -1
		return e
	}
}