package main

import (
	"fmt"
	"sync"
)

var count int

func main() {
	mu := &sync.Mutex{}
	size := 100
	countMap := make(map[int]int)
	wg := &sync.WaitGroup{}
	for i := 0; i < size; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer mu.Unlock()
			count++
			countMap[i] = i
			defer wg.Add(-1) //similar as wg.Done
		}()
	}
	wg.Wait()
	fmt.Println(count)
	// fmt.Println(reverseSimple(arr))
}
