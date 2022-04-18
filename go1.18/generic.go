/**
 * @Author: alessonhu
 * @Description:
 * @File:  generic
 * @Version: 1.0.0
 * @Date: 2022/3/31 4:35 PM
 */

package main

import "fmt"

type Number interface {
	int | int64 | float64
}

func main() {
	ins := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Println(Sum(ins))
	fmt.Println(Sum(floats))

	fmt.Println(SumNumber(1.0, 2.0))
}

func Sum[K comparable, V int | float32 | int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func SumNumber[N Number](a, b N) N {
	return a + b
}
