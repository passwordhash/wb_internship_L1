package main

import "fmt"

func main() {
	x := 10
	y := 25

	x, y = y, x

	fmt.Println(x, y)
}
