package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var str1, str2 string
	_, _ = fmt.Scan(&str1, &str2)
	t := time.Now()
	for i := 0; i < len(str1); i++ {
		str2 = strings.Replace(str2, string(str1[i]), "", 1)
	}
	fmt.Println(str2)
	time.Sleep(time.Second)
	fmt.Print(time.Since(t))
}
