package main

import (
	"fmt"
	"sync"
	"time"
)

func loader(data []int, m *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, x := range data {
		m.Store(time.Now().String(), x)
	}
}

func main() {
	var wg sync.WaitGroup
	var m sync.Map

	n := 100

	var data []int
	for i := 0; i < n; i++ {
		data = append(data, i)
	}

	wg.Add(1)
	go loader(data[:n/2], &m, &wg)

	wg.Add(1)
	go loader(data[n/2:], &m, &wg)

	wg.Wait()
	m.Range(func(key, value any) bool {
		fmt.Printf("%s: %d\n", key, value)
		return true
	})
}
