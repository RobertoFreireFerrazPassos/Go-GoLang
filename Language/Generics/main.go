package main

import (
	"fmt"
)

type Number interface {
	int8 | int64 | float64
}

func sumNumbers[N Number](s []N) N {
	var total N
	for _, num := range s {
		total += num
	}
	return total
}

func main() {
	ints := []int64{32, 64, 96, 128}
	floats := []float64{32.0, 64.0, 96.1, 128.2}
	bytes := []int8{8, 16, 24, 32}

	fmt.Println(sumNumbers(ints))
	fmt.Println(sumNumbers(floats))
	fmt.Println(sumNumbers(bytes))
}
