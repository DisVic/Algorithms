package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	t := time.Now()
	fmt.Println(decToBin(n))
	fmt.Print(time.Since(t))
}

func decToBin(n int) string {
	str := ""
	for n > 0 {
		str += strconv.Itoa(n % 2)
		n /= 2
	}
	binNumber := ""
	for i := 0; i < len(str); i++ {
		binNumber += string(str[len(str)-1-i])
	}
	return binNumber
}
