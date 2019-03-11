package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age int
}
type HeroSlice []Hero

func (hs HeroSlice)Len()int {
	return len(hs)
}

func (hs HeroSlice)Less(i,j int)bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice)Swap(i,j int){
	hs[i],hs[j] = hs[j],hs[i]
}

type Student struct {
	Name string
	Age int
	Score float64
}

func main() {

	var intSlice =[]int{0,2,44,5,677}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	var heros HeroSlice
	for i:=0; i<10; i++{
		hero := Hero{
			Name:fmt.Sprintf("英雄%d",rand.Intn(100)),
			Age:rand.Intn(100),
		}
		heros = append(heros,hero)
	}
	for _,v := range heros{
		fmt.Println(v)
	}
	sort.Sort(heros)
	fmt.Println("  ----paixuhou==--")

	for _,v := range heros{
		fmt.Println(v)
	}

	i:=10
	j:=20
	i,j = j,i
	fmt.Println(i,j)
}