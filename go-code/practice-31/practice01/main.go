package main

import "fmt"

type ValNode struct {
	raw int
	col int
	val int
}

func main() {

	//创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	//输出查看原始数组
	for _,v := range chessMap {
		for _,v2 := range v {
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}

	//转成稀疏数组
	var sparseArr []ValNode

	valNode := ValNode{
		raw:11,
		col:11,
		val:0,
	}
	sparseArr = append(sparseArr,valNode)

	for i,v := range chessMap {
		for j,v2 := range v {
			if v2 != 0 {
				valNode = ValNode{
					raw:i,
					col:j,
					val:v2,
				}
				sparseArr = append(sparseArr,valNode)
			}
		}
	}

	//输出稀疏数组
	fmt.Println("当前的稀疏数组是。。。")
	for i,valNode := range sparseArr {
		fmt.Printf("%d:%d %d %d\n",i,valNode.raw,valNode.col,valNode.val)
	}

	//将稀疏数组恢复原始数组
	var chessMap2 [11][11]int

	for i,valNode := range sparseArr {
		if i != 0 {
			chessMap2[valNode.raw][valNode.col] = valNode.val
		}
	}

	//输出chessMap2 的值
	fmt.Println("恢复后的原始数据")
	for _,v := range chessMap {
		for _,v2 := range v {
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}
}