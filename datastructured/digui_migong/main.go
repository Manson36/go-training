package main

import "fmt"

//完成老鼠找路。使用引用，保证是同一个地图；i，j表示测试的点
func SetWay(myMap *[8][7]int, i, j int) bool {
	//分析什么情况找到出路，终点标记为2
	if myMap[6][5] == 2 {
		return true
	} else {
		//说明要继续查找
		if myMap[i][j] == 0 {//说明这个点可以探测
			//假设这个点可以通，需要继续探测；因为在左下方，使用下右上左更好
			myMap[i][j] = 2
			if SetWay(myMap, i + 1, j) {
				return true
			} else if SetWay(myMap, i, j+1) {
				return true
			} else if SetWay(myMap, i - 1, j) {
				return true
			} else if SetWay(myMap, i, j-1) {
				return true
			} else {//说明是死路
				myMap[i][j] = 3
				return false
			}
		}else {//说明这个点不能测，是墙
			return false
		}
	}
}

func main() {
	//先创建一个二维数组，模拟迷宫，规则：
	//数字0代表为走过，1代表墙，2代表通路，3代表走过但是不通，
	var myMap[8][7]int

	//将最上和最下设为1
	for i:= 0; i < len(myMap[0]); i++ {
		myMap[0][i] = 1
		myMap[len(myMap)-1][i] = 1
	}
	//将最左和最右设为1
	for i := 0; i < len(myMap); i++ {
		myMap[i][0] = 1
		myMap[i][len(myMap[0])-1] = 1
	}

	//设置墙
	myMap[3][1] = 1
	myMap[3][2] = 1
	myMap[2][1] = 1
	myMap[2][2] = 1

	//查看迷宫
	for i:= 0; i < len(myMap); i++ {
		for j := 0; j <  len(myMap[0]); j++ {
			fmt.Printf("%3d", myMap[i][j])
		}
		fmt.Println()
	}

	//开始测试
	SetWay(&myMap, 1,1)
	fmt.Println("测试结果")
	for i:= 0; i < len(myMap); i++ {
		for j := 0; j <  len(myMap[0]); j++ {
			fmt.Printf("%3d", myMap[i][j])
		}
		fmt.Println()
	}
}
