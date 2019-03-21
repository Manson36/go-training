package main

import "fmt"

func main() {

	grade := "E"
	marks := 90

	switch {

	case marks >= 90:
		grade = "A"
	case marks >= 80:
		grade = "B"
	case marks >= 70:
		grade = "C"
	case marks >= 60:
		grade = "D"
	default:
		grade = "E"
	}

	switch {
	case grade == "A":
		fmt.Println("成绩优秀")
	case grade == "B":
		fmt.Println("成绩良好")
	case grade == "C",grade == "D":
		fmt.Println("成绩及格")
	default:
		fmt.Println("成绩不及格")
	}
	fmt.Println("你的成绩是",grade)
}