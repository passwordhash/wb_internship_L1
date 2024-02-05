package main

import "fmt"

func setBit(n int64, pos uint) int64 {
	n |= 1 << pos
	return n
}

func clearBit(n int64, pos uint) int64 {
	mask := ^(1 << pos)
	n &= int64(mask)
	return n
}

func main() {
	n := int64(10)

	fmt.Printf("Было: %b\n", n)

	n = setBit(n, 0)

	fmt.Printf("%b\n", n)

	n = clearBit(n, 1)

	fmt.Printf("%b\n", n)
}
