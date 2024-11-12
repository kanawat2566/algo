package maxsum

import "fmt"

func MaxSumRange(arr []int, w int) int {
	start := 0
	sumMax := 0
	for i := 0; i < len(arr); i++ {
		sum := 0
		fmt.Printf("Round: %v ", i)
		for j := i; j < i+w-1 && j < len(arr); j++ {
			sum += arr[j]
			fmt.Printf(" %v", arr[j])
		}
		if sumMax < sum {
			sumMax = sum
			start = i
		}
		fmt.Println()
	}
	return start
}
