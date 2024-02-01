package main

import "fmt"

// 1 способ: буферизированный канал
//func main() {
//	var wg sync.WaitGroup
//	nums := []int{2, 4, 6, 8, 10}
//	write := make(chan int, len(nums))
//
//	for _, num := range nums {
//		wg.Add(1)
//		go func(group *sync.WaitGroup, num int) {
//			defer group.Done()
//			write <- num * num
//		}(&wg, num)
//	}
//
//	wg.Wait()
//	close(write)
//
//	for num := range write {
//		fmt.Println(num)
//	}
//}

// 2 способ: pipeline
func main() {
	nums := []int{2, 4, 6, 8, 10}
	toSquare := make(chan int)
	toPrint := make(chan int)

	go func() {
		for _, num := range nums {
			toSquare <- num
		}
		close(toSquare)
	}()

	go func() {
		for num := range toSquare {
			toPrint <- num * num
		}
		close(toPrint)
	}()

	for num := range toPrint {
		fmt.Println(num)
	}
}
