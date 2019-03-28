package main

import "fmt"

//编写一个学生考试系统
type Student struct {
	Name string
	Age int
	Score float64
}
//显示他的成绩
func (p *Student)ShowInfo() {
	fmt.Println("学生名",p.Name,"年龄",p.Age,"成绩",p.Score)
}

func (p *Student)SetScore(score float64) {
	p.Score = score
}

type Pupil struct {
	Student
}

func (p *Pupil)Test() {
	fmt.Println("小学生正在考试...")
}

type Graduate struct {
	Student
}
func (p *Graduate)Test() {
	fmt.Println("大学生正在考试...")
}

func main() {

	pupil := &Pupil{}
	pupil.Student.Name = "sam"
	pupil.Student.Age = 11
	pupil.Test()
	pupil.Student.SetScore(4.4)
	pupil.Student.ShowInfo()

	graduate := &Graduate{}
	graduate.Student.Name = "sam"
	graduate.Student.Age = 31
	graduate.Test()
	graduate.Student.SetScore(8.4)
	graduate.Student.ShowInfo()
}