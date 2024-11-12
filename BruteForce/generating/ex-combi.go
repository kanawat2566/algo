package generating

import "fmt"

func Combi(n int, sol [4]int, len int) {
	if len < n {
		sol[len] = 1
		Combi(n, sol, len+1)

		sol[len] = 0
		Combi(n, sol, len+1)

	} else {
		fmt.Printf("%v", sol)
	}
}
