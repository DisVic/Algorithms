package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var length int
	var strNumber string
	_, _ = fmt.Scan(&length)
	for i := 0; i < length; i++ {
		n := 0
		_, _ = fmt.Scan(&n)
		strNumber += strconv.Itoa(n)
	}
	var newNumber string
	var number1, number2 int
	_, _ = fmt.Scan(&newNumber)
	t := time.Now()
	number1, _ = strconv.Atoi(strNumber)
	number2, _ = strconv.Atoi(newNumber)
	number := number1 + number2
	strNumber = strconv.Itoa(number)
	for i := 0; i < len(strNumber); i++ {
		fmt.Printf("%c ", strNumber[i])
	}
	time.Sleep(time.Second)
	fmt.Println(time.Since(t))
}
