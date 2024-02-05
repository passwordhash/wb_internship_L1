package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// 1: завершение горутины с помощью канала
func withChan() {
	done := make(chan struct{})

	go func() {
		defer fmt.Println("I did !")
		for {
			select {
			case <-done:
				return
			default:
				fmt.Println("doing something ...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// Завершнаем горутину спустя 4 секунды
	time.Sleep(4 * time.Second)
	done <- struct{}{}

	fmt.Println("End of main goroutine")
}

// 2: завершение с помощью контекста с отменой
func withCancelContext() {
	ctx, cancel := context.WithCancel(context.TODO())

	go func(ctx context.Context) {
		defer fmt.Println("I did !")
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("doing something ...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	// Завершнаем горутину спустя 4 секунды
	time.Sleep(4 * time.Second)
	cancel()

	fmt.Println("End of main goroutine")
}

// 3: завершения с помощью контекста с таймаутом
//
// Горутина завершит свое выполнение через 3 секунды
func withCancelTimeout() {
	ctx, cancel := context.WithTimeout(context.TODO(), 3*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("doing something ...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	<-ctx.Done()
	fmt.Println("End of main goroutine")
}

// 4: завершение горутины с помощью канала сигналов
//
// Горутина будет работать пока не будет нажатa ctrl + C
func withChanOfSigs() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		for {
			select {
			case <-sigs:
				return
			default:
				fmt.Println("doing something ...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	<-sigs
	fmt.Println("End of main goroutine")
}

// 5: завершнеие горутины с завершением main
func main_() {
	go func() {
		for {
			fmt.Println("doing something ...")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(6 * time.Second)
	fmt.Println("End of main goroutine")
}

// 6: завершение горутины по истечении таймера
func withTimer() {
	timer := time.NewTimer(5 * time.Second)

	go func(timer *time.Timer) {
		for {
			select {
			case <-timer.C:
				return
			default:
				fmt.Println("doing something ...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(timer)

	<-timer.C
	fmt.Println("End of main goroutine")
}
