package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	findDistance()

}

func findDistance() {
	var n int
	//var err error
	_, _ = fmt.Scan(&n)
	// if err != nil {
	// 	fmt.Println("ошибка")
	// 	return
	// }
	// if n <= 0 {
	// 	fmt.Println("Пустая строка данных")
	// 	return
	// }
	houses := make([]int, n)
	for i := 0; i < n; i++ {
		var number int
		_, _ = fmt.Scan(&number)
		houses[i] = number
	}
	t := time.Now()

	var wg sync.WaitGroup
	var mutex sync.Mutex
	counter := 0

	// Forward pass
	wg.Add(n)
	for i := 0; i < len(houses); i++ {
		go func(i int) {
			defer wg.Done()
			mutex.Lock()
			if houses[i] == 0 {
				counter = 0
			} else {
				counter++
				houses[i] = counter
			}
			mutex.Unlock()
		}(i)
	}

	wg.Wait() // Wait for all forward pass goroutines to complete

	// Reset counter
	counter = 0

	// Backward pass
	wg.Add(n)
	for i := len(houses) - 1; i >= 0; i-- {
		go func(i int) {
			defer wg.Done()
			mutex.Lock()
			if houses[i] == 0 {
				counter = 0
			} else if counter < houses[i] {
				counter++
				houses[i] = counter
			}
			mutex.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Print(houses)
	fmt.Println(time.Since(t))
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
