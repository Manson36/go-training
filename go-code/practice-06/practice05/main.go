package main

import "fmt"

func main() {

	//演示for-range的遍历
	heros := [...]string{"武松", "林冲", "花荣"}

	for index, value := range heros {
		fmt.Printf("index=%v value=%v\n", index, value)
		fmt.Printf("heros[%d]=%v\n", index, value)
	}

	for _, value := range heros {
		fmt.Printf("元素的值是%v\n", value)
	}

	//创建一个数组放置A-Z字母
	var myChar [26]byte

	for i := 0; i < 26; i++ {
		myChar[i] = 'A' + byte(i)
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%c", myChar[i])

	}
}