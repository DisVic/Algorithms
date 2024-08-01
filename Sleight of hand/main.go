package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var k int
	_, _ = fmt.Scan(&k)

	game := make([]string, 4)

	for i := 0; i < len(game); i++ {
		_, _ = fmt.Scan(&game[i])
	}
	t := time.Now()
	counterDigit := make(map[rune]int)

	for _, str := range game {
		for _, ch := range str {
			if ch != '.' {
				counterDigit[ch]++
			}
		}
	}

	points := 0

	for _, counter := range counterDigit {
		if k*2 >= counter {
			points++
		}
	}

	//fmt.Print(points)
	fmt.Print(time.Since(t))
	printMemStats()
}

func printMemStats() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	fmt.Printf("Allocated Memory: %v bytes\n", memStats.Alloc)
	fmt.Printf("Total Allocated Memory: %v bytes\n", memStats.TotalAlloc)
	fmt.Printf("Heap Memory: %v bytes\n", memStats.HeapAlloc)
	fmt.Printf("Heap System Memory: %v bytes\n", memStats.HeapSys)
	fmt.Printf("Garbage Collector Memory: %v bytes\n", memStats.GCSys)
}
