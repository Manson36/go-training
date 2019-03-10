package main

import "fmt"

type Goods struct {
	Name string
	Price float64
}
type Brand struct {
	Name string
	Address string
}

type TV struct {
	Goods
	Brand
}

func main() {
	tv := TV{
		Goods{
			Price: 5000.2,
			Name: "电视机02",
		},
		Brand{
			Name: "夏普",
			Address: "北京",
		},
	}
	fmt.Println(tv.Goods.Name)
	fmt.Println(tv.Price)
}