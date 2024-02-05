package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	var nums []int
	for i := 0; i < 20; i++ {
		nums = append(nums, i)
	}

	go func() {
		for _, x := range nums {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * 2
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}
