package main

import (
	"fmt"
	"math/big"
)

func main() {
	//a := big.NewInt(int64(math.Pow(2, 24)))
	//b := big.NewInt(int64(math.Pow(2, 21)))
	a, b := new(big.Int), new(big.Int)
	a.SetString("92233720368547758054", 10)
	b.SetString("9223372036854775801", 10)

	div := new(big.Int).Div(a, b)
	fmt.Println(div)

	mult := new(big.Int).Mul(a, b)
	fmt.Println(mult)

	sum := new(big.Int).Add(a, b)
	fmt.Println(sum)

	sub := new(big.Int).Sub(a, b)
	fmt.Println(sub)
}
