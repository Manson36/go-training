package main

import (
	"fmt"
	"reflect"
)

type TagType struct { //tags
	field1 bool    "是否有存货"
	field2 string  "商品名称"
	field3 int     "商品价格"
}

func main() {

	tt := TagType{true, "LPhone X",1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
}

func refTag(tt TagType, ix int) {

	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Println(ixField.Tag)
}
