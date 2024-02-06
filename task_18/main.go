package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu sync.Mutex
	c  int
}

// increment требует синхронизации потоков
func (c *Counter) increment() {
	c.c++
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.increment()
	c.mu.Unlock()
}

// Load метод не потокобезопасный
func (c *Counter) Load() int {
	return c.c
}

func main() {
	counter := Counter{}

	go worker(&counter)
	go worker(&counter)
	go worker(&counter)
	go worker(&counter)

	time.Sleep(4 * time.Second)

	fmt.Println(counter.Load())
}

func worker(c *Counter) {
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			c.Increment()
		}
	}
}
