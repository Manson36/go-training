package main

import (
	"fmt"
)

type nodeVal struct {
	row int
	col int
	val int
}

func main() {
	//1.创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //1代表黑子
	chessMap[2][3] = 2 //2代表白字

	//2.输出查看原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t",v2)
		}
		fmt.Println()
	}

	//3.转为稀疏数组 ：使用结构体的想法就是算法
	//思路：1）遍历chessMap，发现一个元素不为0，创建一个nodeVal结构体；2）将其放入对应的切片
	var sparseArray []nodeVal

	//标准的稀疏数组应该还有一个值结点，记录二维数组的规模（行、列、默认值）
	valNode := nodeVal{
		row: 11,
		col: 11,
		val: 0}

	sparseArray = append(sparseArray, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
			//创建一个值结点
			nodeVal := nodeVal{
				row: i,
				col: j,
				val: v2,
			}
			sparseArray = append(sparseArray, nodeVal)
		}
		}
	}

	//输出稀疏数组
	fmt.Println("当前的稀疏数组是：")
	for i, valNode := range sparseArray {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}

	//将稀疏数组存盘，在从盘中读取为二维数组，可以参考迷宫
	//这里直接将稀疏数组转换为二维数组
	var chessMap2 [11][11]int

	for i, valNode := range sparseArray {
		if i != 0 {
			chessMap2[valNode.row][valNode.col] = valNode.val
		}
	}

	fmt.Println("恢复后的原始数据")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}

