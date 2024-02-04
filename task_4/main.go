package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

// data - канал с данными
var data = make(chan interface{}, 10)

// Я использовал конекст с функцией отмены для того,
// чтобы разом завершить все воркеры.
//
// Вместо того чтобы пытаться явно остановить каждый
// воркер отдельно, просто отменяем контекст, и каждый воркер,
// который его слушает, получает сигнал о необходимости остановиться
func main() {
	var wg sync.WaitGroup
	workersCount := *(flag.Int("n", 5, "Количество воркеров"))

	flag.Parse()

	// Контекст, который слушают все воркеры
	workersCtx, cancel := context.WithCancel(context.TODO())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	// Горутина для отслеживания cигнала
	go func() {
		<-sigs
		fmt.Println()
		cancel()
	}()

	// Запуск нужного количества воркеров
	for i := 1; i <= workersCount; i++ {
		wg.Add(1)
		go worker(workersCtx, &wg, data, i)
	}

	// Генерация тестовых данных в канал data
	go func() {
		s := "Minim officia aliqua amet officia do lorem incididunt mollit ullamco ex amet veniam non. Consectetur velit ullamco irure adipiscing voluptate ut deserunt officia pariatur est eu sit labore cillum sint adipiscing commodo exercitation dolor. Fugiat ex enim consequat officia velit exercitation pariatur duis esse proident tempor aliquip. Qui ea sed deserunt ea ex cupidatat commodo laborum sint voluptate officia minim eiusmod anim occaecat mollit aliquip. Minim in quis duis dolor lorem eu irure proident ullamco consequat nulla aute labore enim deserunt dolore magna nulla consequat. Dolor id amet nulla aute labore elit aliqua nisi aliqua aliquip cillum aliquip aliqua labore elit do in. Magna excepteur et "

		input := strings.Split(s, " ")

		for i := 0; i < 5000; i++ {
			data <- input[rand.Intn(len(input))]
		}
	}()

	wg.Wait()
	fmt.Println("Все воркеры завершили работу")
}

func worker(ctx context.Context, wg *sync.WaitGroup, data <-chan interface{}, id int) {
	defer wg.Done()
	fmt.Println("Создан новый воркер")
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d завершает работу\n", id)
			return
		case msg := <-data:
			fmt.Fprintf(os.Stdout, "worker %d: %s\n", id, msg)
		}
	}
}
