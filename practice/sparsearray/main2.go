package main

import "fmt"

type nodeVal struct {
	row int
	cal int
	val int
}

func main() {
	//创建原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2
	chessMap[3][4] = 1

	//查看
	for _, v1 := range chessMap {
		for _, v2 := range v1 {
			fmt.Printf("%3d", v2)
		}
		fmt.Println()
	}

	//转换为稀疏数组
	var sparseArray []nodeVal
	sparseArray = append(sparseArray, nodeVal{
		row: 11, cal: 11, val: 0,
	})

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				sparseArray = append(sparseArray, nodeVal{row: i, cal: j, val: v2})
			}
		}
	}

	//查看洗漱数组
	fmt.Println("当前的稀疏数组是：")
	for i, v := range sparseArray {
		fmt.Printf("%d: %d %d %d\n", i, v.row, v.cal, v.val)
	}

	//把稀疏数组转换为二维数组
	var chessMap2 [11][11]int
	for i, v := range sparseArray {
		if i != 0 {
			chessMap2[v.row][v.cal] = v.val
		}
	}

	fmt.Println("恢复后的源数据是：")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%3d", v2)
		}
		fmt.Println()
	}
}
