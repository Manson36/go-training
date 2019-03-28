package main

import "fmt"

//编写一个小学生考试系统
type Pupil struct {
	Name string
	Age int
	Score float64
}
//显示他的成绩
func (p *Pupil)ShowInfo() {
	fmt.Println("学生名",p.Name,"年龄",p.Age,"成绩",p.Score)
}

func (p *Pupil)SetScore(score float64) {
	p.Score = score
}

func (p *Pupil)Test() {
	fmt.Println("小学生正在考试")
}

type Graduate struct {
	Name string
	Age int
	Score float64
}
//显示他的成绩
func (p *Graduate)ShowInfo() {
	fmt.Println("学生名",p.Name,"年龄",p.Age,"成绩",p.Score)
}

func (p *Graduate)SetScore(score float64) {
	p.Score = score
}

func (p *Graduate)Test() {
	fmt.Println("大学生正在考试")
}

func main() {

	var pupil = &Pupil{
		Name:"Tom",
		Age:12,
	}
	pupil.Test()
	pupil.SetScore(336.2)
	pupil.ShowInfo()

	var graduate = &Graduate{
		Name:"Tom",
		Age:22,
	}
	graduate.Test()
	graduate.SetScore(336.2)
	graduate.ShowInfo()
}