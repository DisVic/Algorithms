package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	t := time.Now()
	isPower(n)
	time.Sleep(time.Second)
	fmt.Println(time.Since(t))
}

func isPower(number int) {
	for i := 0; math.Pow(4, float64(i)) <= float64(number); i++ {
		if math.Pow(4, float64(i)) == float64(number) {
			fmt.Println("True")
			return
		}
	}
	fmt.Println("False")
	return
}
