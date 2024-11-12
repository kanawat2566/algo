package perfectnumber

import "fmt"

func Perfectnumbers(a, b int) []int {
	if a > b {
		return []int{}
	}
	var ans []int
	for i := a; i < b; i++ {
		// check perfact number
		if IsPerfactNumber(i) {
			ans = append(ans, i)
		}

	}

	return ans
}

func IsPerfactNumber(x int) bool {
	var elm []int
	for i := 1; i < x; i++ {
		if x%i == 0 {
			elm = append(elm, i)
		}
	}

	sum := 0
	for i := 0; i < len(elm); i++ {
		sum += elm[i]
	}

	fmt.Printf("num:%v = %v sum=%v \n", x, elm, sum)

	if sum == x {
		return true
	} else {
		return false
	}

}
