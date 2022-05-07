package main

import (
	"fmt"
)

var arr []int = []int{1, 2, 3}

func reverse(a []int) (res []int) {
	for _, i := range a {
		defer func(i int) {
			res = append(res, i)
		}(i)
	}
	return res
}
func main() {
	fmt.Println(reverse(arr))
}
