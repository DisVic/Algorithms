package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	text := strings.Split(inputBufio(), " ")
	t := time.Now()
	fmt.Println(findLongestWord(text))
	fmt.Print(time.Since(t))
}

func inputBufio() string {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	str := scanner.Text()
	return str
}

func findLongestWord(text []string) (string, int) {
	counter := 0
	word := ""
	for i := 0; i < len(text); i++ {
		if len(text[i]) > counter {
			word = text[i]
			counter = len(text[i])
		}
	}
	return word, counter
}
