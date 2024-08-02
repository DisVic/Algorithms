package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	str := inputBufio()
	t := time.Now()
	fmt.Println(isPalindrom(str))
	fmt.Print(time.Since(t))
}

func inputBufio() string {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	str := scanner.Text()
	return str
}

func isPalindrom(str string) bool {
	re := regexp.MustCompile(`[[:punct:]]|[[:space:]]`)
	str = re.ReplaceAllString(str, "")
	str = strings.ToLower(str)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
