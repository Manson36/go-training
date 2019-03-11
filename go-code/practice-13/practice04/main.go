package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}
type Phone struct {
	name string
}

func (p Phone)Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone)Stop() {
	fmt.Println("手机停止工作。。。")
}
func (p Phone)Call() {
	fmt.Println("手机正在打电话。。。")
}

type Camera struct {
	name string
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
	//类型断言

	if phone, ok := usb.(Phone);ok {
		phone.Call()
	}
	usb.Stop()
}

func main () {

	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"xioami"}
	usbArr[2] = Camera{"nikang"}

	//遍历USBArr
	var computer Computer
	for _,v := range usbArr{
		computer.Working(v)
		fmt.Println()
	}


}