package main

import "fmt"

func QuickSort(value []int) {
	if len(value) <= 1 {
		return
	}

	mid, i := value[0], 1
	head, tail := 0, len(value) -1

	for head < tail  {
		if value[i] > mid {
			value[i], value[tail] = value[tail], value[i]
			tail--
		} else {
				value[i], value[head] = value[head], value[i]
				head++
				i++
		}
	}

	QuickSort(value[:head])
	QuickSort(value[head+1:])
}

func main() {
	var values = []int{2, 0, 33, 55, 22, 32, 89, 11}
	QuickSort(values)
	fmt.Println(values)
}
