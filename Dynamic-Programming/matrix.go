package dynamicprogramming

import "fmt"

func MatrixChan() {
	var a = [][]int{
		{10, 8},
		{2, 4},
		{5, 6},
	}
	b := [][]int{
		{7, 1},
		{3, 9},
	}
	var c [3][2]int

	for i := 0; i < len(a); i++ {
		var sum int = 0
		for j := 0; j < len(c); j++ {

			for k := 0; k < len(b); k++ {
				sum += a[i][j] * b[j][i]
			}
			c[i][j] = sum
			fmt.Printf("c[%v][%v] = %v ", i, j, sum)
		}
		fmt.Println()
	}
}
