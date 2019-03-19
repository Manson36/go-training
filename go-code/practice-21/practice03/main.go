package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn,err := net.Dial("tcp","192.168.20.253:8888")

	if err != nil {
		fmt.Println("cient dial err =",err)
		return
	}
	reader := bufio.NewReader(os.Stdin)

	line,err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("readString err=",err)

	}
	n,err := conn.Write([]byte(line))
	if err != nil {
		fmt.Println("conn.Write err=",err)
	}
	fmt.Printf("客户端发送了 %d 字节的数据,并退出",n )
}