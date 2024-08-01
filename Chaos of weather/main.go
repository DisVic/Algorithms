package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	weather := make([]int, n)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(weather); i++ {
		weather[i] = rand.Intn(273)
		//_, _ = fmt.Scan(&weather[i])
	}
	t := time.Now()
	time.Sleep(time.Second)
	fmt.Println(findMaxDegrees(weather))
	fmt.Print(time.Since(t))
}

func findMaxDegrees(numbers []int) int {
	counter := 0
	if numbers[0] > numbers[1] {
		counter++
	}
	if numbers[len(numbers)-1] > numbers[len(numbers)-2] {
		counter++
	}
	for i := 1; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i-1] && numbers[i] > numbers[i+1] {
			counter++
		}
	}
	return counter
}
