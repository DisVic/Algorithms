package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {
	t := time.Now()
	time.Sleep(time.Second)
	fmt.Println(quadraticFunction(-8, -5, -2, 7))
	fmt.Println(quadraticFunction(8, 2, 9, -10))
	fmt.Println(quadraticFunction(-85, 7184, -2786, 25))
	fmt.Print(time.Since(t))
	printMemStats()
}

func quadraticFunction(a, b, c, x float64) float64 {
	return math.Pow(x, 2)*a + x*b + c
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
