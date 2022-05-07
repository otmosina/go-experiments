package main

import (
	"fmt"
)

var arr []int = []int{1, 2, 3}

func reverseDefer(a []int) (res []int) {
	for _, i := range a {
		defer func(i int) {
			res = append(res, i)
		}(i)
	}
	return res
}

func reverseSimple(a []int) (res []int) {
	for i := len(a) - 1; i >= 0; i-- {
		res = append(res, a[i])
	}
	return res
}
func main() {
	fmt.Println(reverseDefer(arr))
	fmt.Println(reverseSimple(arr))
}
