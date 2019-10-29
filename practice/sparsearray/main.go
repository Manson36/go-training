package main

import "fmt"

type nodeVal struct {
	row int
	col int
	val int
}

func main() {
	//创建一个chessmap
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[3][2] = 2
	chessMap[5][6] = 1
	chessMap[7][8] = 2

	//查看chessMap
	//for i := 0; i < len(chessMap); i++ {
	//	for j := 0; j < len(chessMap[0]); j++ {
	//		fmt.Printf("%3d", chessMap[i][j])
	//	}
	//	fmt.Println()
	//}
	for _, v1 := range chessMap {
		for _, v2 := range v1 {
			fmt.Printf("%3d", v2)
		}
		fmt.Println()
	}

	//转换稀疏数组
	var sparseArray []nodeVal
	valNode := nodeVal{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArray = append(sparseArray, valNode)

	for i, v1 := range chessMap {
		for j, v2 := range v1 {
			if v2 != 0 {
				valNode := nodeVal{
					row: i,
					col: j,
					val: v2,
				}
				sparseArray = append(sparseArray, valNode)
			}
		}
	}

	fmt.Println("当前的稀疏数组是：")
	for i, valNode := range sparseArray {
		fmt.Printf("%d: %d, %d, %d\n", i, valNode.row, valNode.col, valNode.val)
	}

	var chessMap2 [11][11]int
	for i, v := range sparseArray {
		if i != 0 {
			chessMap2[v.row][v.col] = v.val
		}
	}

	for _, v1 := range chessMap2 {
		for _, v2 := range v1 {
			fmt.Printf("%3d", v2)
		}
		fmt.Println()
	}
}
