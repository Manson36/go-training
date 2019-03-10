package main

import (
	"fmt"
)

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
type TV2 struct {
	*Goods
	*Brand
}

func main() {

	tv := TV{Goods{"电视机01",5000.2},Brand{"海尔","山东"}}
	tv2 := TV{
		Goods{
			Price: 5000.2,
			Name: "电视机02",
		},
		Brand{
			Name: "夏普",
			Address: "北京",
		},
	}
	fmt.Println("tv=",tv)
	fmt.Println("tv2",tv2)

	tv3 := TV2{&Goods{"电视机3",43322},&Brand{"创维","河南"}}
	tv4 := TV2{
		&Goods{
			Name: "电视机04",
			Price: 33333,
		},
		&Brand{
			Name: "长虹",
			Address:"四川",
		},
	}
	fmt.Println("tv3",*tv3.Goods,*tv3.Brand)
	fmt.Println("tv4",*tv4.Goods,*tv4.Brand)
}