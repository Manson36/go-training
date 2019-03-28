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

type TV2 struct {
	*Goods
	*Brand
}

func main() {

	tv := TV{Goods{"电视机001",22222.2},Brand{"海尔","山东"}}

	tv2 := TV{
		Goods{
			"电视机002",
			111.11,
		},
		Brand{
			"夏普",
			"北京",
		},
	}
	fmt.Println("tv=",tv)
	fmt.Println("tv2=",tv2)

	tv3 := TV2{&Goods{"电视机003",22222.2},&Brand{"海尔","山东"}}

	tv4 := TV2{
		&Goods{
			"电视机004",
			111.11,
		},
		&Brand{
			"夏普",
			"北京",
		},
	}
	fmt.Println("tv3=",*tv3.Goods,*tv3.Brand)
	fmt.Println("tv4=",*tv4.Goods,*tv4.Brand)
}