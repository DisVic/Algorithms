package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	var a, b string
	_, _ = fmt.Scan(&a, &b)
	t := time.Now()
	number1 := new(big.Int)
	number2 := new(big.Int)
	number1.SetString(a, 2)
	number2.SetString(b, 2)
	sum := new(big.Int).Add(number1, number2)
	fmt.Println(sum.Text(2))
	fmt.Print(time.Since(t))
}
