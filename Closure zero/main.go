package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	findDistance()

}

func findDistance() {
	var n int
	var err error
	_, err = fmt.Scan(&n)
	if err != nil {
		fmt.Println("ошибка")
		return
	}
	if n <= 0 {
		fmt.Println("Пустая строка данных")
		return
	}
	houses := make([]int, n)
	for i := 0; i < n; i++ {
		var number int
		_, _ = fmt.Scan(&number)
		houses[i] = number
	}
	t := time.Now()
	counter := 0
	for i := 0; i < len(houses); i++ {
		time.Sleep(time.Second)
		if houses[i] == 0 {
			counter = 0
		} else {
			counter++
			houses[i] = counter
		}
	}
	for i := len(houses) - 1; i > 0; i-- {
		if houses[i] == 0 {
			counter = 0
		} else if counter < houses[i] {
			counter++
			houses[i] = counter
		} else {
			continue
		}
	}
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
