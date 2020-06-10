package main

import (
	"fmt"
)

func main() {
	values := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	var result []int
	var step int

	for i := 0; i < len(values); i++ {
		result = append(result, values[0][i])
		step++
	}

	for i := 1; i < len(values); i++ {
		result = append(result, values[i][len(values)-1])
		step++
	}

	for i := len(values) - 2; i >= 0; i-- {
		result = append(result, values[len(values)-1][i])
		step++
	}

	for i := len(values) - 2; i > 0; i-- {
		result = append(result, values[i][0])
		step++
	}

	for i := 1; i <= len(values)*len(values); i++ {
		fmt.Println(i)
	}

	fmt.Println("result", result)
	fmt.Println("step", step)
}
