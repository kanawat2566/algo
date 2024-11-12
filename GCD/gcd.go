package gcd

import "fmt"

func GCD(a, b int) int {
	for b > 0 {
		tmp := b
		b = a % b
		a = tmp
		fmt.Printf("b:%v \n", b)
	}
	return a
}
