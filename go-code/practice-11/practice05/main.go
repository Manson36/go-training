package main

import "fmt"

type Pupil struct {
	Name string
	Age int
	Score int
}
//显示成绩
func (p *Pupil)ShowInfo(){
	fmt.Printf("学生名%v 学生年龄%v 学生分数%v",p.Name,p.Age,p.Score)
}

func (p *Pupil)SetScore (score int){
	p.Score = score
}

func (p *Pupil)testing(){
	fmt.Println("小学生正在开始中。。。")
}

type Graduate struct {
	Name string
	Age int
	Score int
}
//显示成绩
func (p *Graduate)ShowInfo(){
	fmt.Printf("学生名%v 学生年龄%v 学生分数%v",p.Name,p.Age,p.Score)
}

func (p *Graduate)SetScore (score int){
	p.Score = score
}

func (p *Graduate)testing(){
	fmt.Println("小学生正在开始中。。。")
}

func main() {

var pupil =Pupil{
	Name:"Tom",
	Age:10,
}

pupil.testing()
pupil.SetScore(90)
pupil.ShowInfo()

	var graduate =Graduate{
		Name:"Tom",
		Age:10,
	}

	graduate.testing()
	graduate.SetScore(90)
	graduate.ShowInfo()
}