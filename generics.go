/**
 * @Author: alessonhu
 * @Description:
 * @File:  generics
 * @Version: 1.0.0
 * @Date: 2022/4/8 5:14 PM
 */

package main

import "fmt"

func print[T any](arr []T) {
	for _, v := range arr {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println()
}

func find[T comparable](arr []T, elem T) int {
	for i, v := range arr {
		if v == elem {
			return i
		}
	}
	return -1
}

type stack[T any] []T

func (s *stack[T]) push(elem T) {
	*s = append(*s, elem)
}

func main() {
	str := []string{"A", "B", "C"}
	num := []int{1, 2, 3}
	print(str)
	print(num)

	fmt.Println(find(num, 2))

	sta := stack[string]{}
	sta.push("111")
	sta.push("222")
	fmt.Println(sta)
}

func gMap[T1 any, T2 any](arr []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(arr))
	for i, elem := range arr {
		result[i] = f(elem)
	}
	return result
}

func gReduce[T1 any, T2 any](arr []T1, init T2, f func(T2, T1) T2) T2 {
	result := init
	for _, elem := range arr {
		result = f(result, elem)
	}
	return result
}
