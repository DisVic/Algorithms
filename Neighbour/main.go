package main

import (
	"fmt"
	"time"
)

func main() {
	var n, m int
	_, _ = fmt.Scan(&n, &m)
	matrix := make([][]int, n)
	counter := 0
	for i := range matrix {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			counter++
			matrix[i][j] = counter
		}
	}
	t := time.Now()
	findNeighbour(matrix)
	fmt.Print(time.Since(t))
}

func findNeighbour(numbers [][]int) {
	var corX, corY int
	_, _ = fmt.Scan(&corX, &corY)
	switch {
	case corX == 0 && corY == 0:
		fmt.Printf("%d %d", numbers[corX+1][corY], numbers[corX][corY+1])
		return
	case corX == len(numbers)-1 && corY == len(numbers)-1:
		fmt.Printf("%d %d", numbers[corX-1][corY], numbers[corX][corY-1])
		return
	case corX == 0 && corY == len(numbers)-1:
		fmt.Printf("%d %d", numbers[corX+1][corY], numbers[corX][corY-1])
		return
	case corX == len(numbers)-1 && corY == 0:
		fmt.Printf("%d %d", numbers[corX-1][corY], numbers[corX][corY+1])
		return
	case corX == len(numbers)-1:
		fmt.Printf("%d %d %d", numbers[corX-1][corY], numbers[corX][corY+1], numbers[corX][corY-1])
		return
	case corX == 0:
		fmt.Printf("%d %d %d", numbers[corX+1][corY], numbers[corX][corY+1], numbers[corX][corY-1])
		return
	case corY == 0:
		fmt.Printf("%d %d %d", numbers[corX+1][corY], numbers[corX-1][corY], numbers[corX][corY+1])
		return
	case corY == len(numbers)-1:
		fmt.Printf("%d %d %d", numbers[corX+1][corY], numbers[corX-1][corY], numbers[corX][corY-1])
		return
	default:
		fmt.Printf("%d %d %d %d", numbers[corX+1][corY], numbers[corX-1][corY], numbers[corX][corY-1], numbers[corX][corY+1])
		return
	}
}
