package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	time.Sleep(time.Second)
	compareOnEven(1, 2, -3)
	compareOnEven(7, 11, 7)
	compareOnEven(6, -2, 0)
	fmt.Print(time.Since(t))
}

func compareOnEven(a, b, c int) {
	if a%2 == 0 && b%2 == 0 && c%2 == 0 || a%2 != 0 && b%2 != 0 && c%2 != 0 {
		fmt.Print("WIN")
	} else {
		fmt.Print("FAIL")
	}
}
