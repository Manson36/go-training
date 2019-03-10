package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}
type Phone struct {

}

func (p Phone)Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone)Stop() {
	fmt.Println("手机停止工作。。。")
}

type Camera struct {

}

func (c Camera)Start() {
	fmt.Println("手机开始工作。。。")
}
func (c Camera)Stop() {
	fmt.Println("手机停止工作。。。")
}

type Computer struct {

}
func (c Computer)Working(usb Usb){
	usb.Start()
	usb.Stop()
}

func main () {

	computer := Computer{}
	phone := Phone{}
	camera :=Camera{}

	computer.Working(phone)
	computer.Working(camera)

}