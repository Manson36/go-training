package main

import (
	"fmt"
)

func writeData(intChan chan int){

	for i:=1; i <= 50; i++ {
		intChan <- i
		fmt.Println("writeData",i)
	}
	close(intChan)
}
func readData(intChan chan int,ExitChan chan bool){

	for {
		v,ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("Data读取的数据",v)
	}
	ExitChan <- true
	close(ExitChan)

}

func main() {

	intChan := make(chan int,50)
	exitChan := make(chan bool,1)

	go writeData(intChan)
	go readData(intChan,exitChan)

	for {
		_,ok := <- exitChan
		if !ok {
			break
		}
	}
}