package main

import (
	"flag"
	"fmt"
	"time"
)

var ch = make(chan int)

func main() {
	duration := *(flag.Duration("t", 5, "Продолжительность работы программы в сек."))

	values := []int{5, 2, 12, 52, 16, 77, 0, -4, 723, 5, 32, -55, 512}

	go func() {
		for _, v := range values {
			ch <- v
			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

	time.Sleep(duration * time.Second)
	fmt.Println("Завершение программы")
}
