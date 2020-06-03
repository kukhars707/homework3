package main

import (
	"fmt"
)

func main() {
	var result []int
	values := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	result = append(result, values[0]...)

	for _, v := range values[1:] {
		result = append(result, v[len(v)-1])
	}

	last := reverseInts(values[len(values)-1][:len(values)-1])

	result = append(result, last...)
	result = append(result, values[1][0:len(values)-1]...)

	fmt.Println(result)
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
