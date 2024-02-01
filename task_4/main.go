package main

import (
	"flag"
	"fmt"
	"os"
)

var data = make(chan interface{})

func worker(ch <-chan interface{}) {
	for {
		fmt.Fprintln(os.Stdout, <-ch)
	}
}

func main() {
	n := flag.Int("n", 1, "Количество воркеров")
	flag.Parse()

	for i := 0; i < *n; i++ {
		go worker(data)
	}

	// тест
	for i := 0; i < 5; i++ {
		data <- i
	}

	for {
	}
}
