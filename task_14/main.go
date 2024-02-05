package main

import (
	"fmt"
	"reflect"
)

func determine(i interface{}) string {
	return reflect.TypeOf(i).String()
}

func main() {
	fmt.Println(determine(10))
	fmt.Println(determine("hello"))
	fmt.Println(determine(true))
	fmt.Println(determine(make(chan int)))
}
