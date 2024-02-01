package main

import (
	"fmt"
	"sync"
)

// Способ 1: буферизированный канал
func main1() {
	nums := []int{2, 4, 6, 8, 10}
	squares := make(chan int, len(nums))
	var wg sync.WaitGroup
	var sum int

	for _, x := range nums {
		wg.Add(1)
		go func(x int) {
			squares <- x * x
			wg.Done()
		}(x)
	}

	wg.Wait()
	close(squares)

	for sq := range squares {
		sum += sq
	}

	fmt.Print(sum)
}

// Способ 2: mutex
func main() {
	nums := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	var mu sync.Mutex
	sum := 0

	for _, x := range nums {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			res := x * x
			mu.Lock()
			sum += res
			mu.Unlock()
		}(x)
	}

	wg.Wait()

	fmt.Println(sum)
}
