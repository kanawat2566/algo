package main

import (
	"fmt"

	twopointer "github.com/kanawat2566/twopointer"
)

//wordcount "github.com/kanawat2566/go-algo/WordCount"

func main() {
	nums := []int{1, 1, 2, 2, 3, 4, 4, 5}
	k := twopointer.RemoveDuplicates(nums)
	fmt.Println("Number of unique elements:", k)
	fmt.Println("Modified array:", nums[:k])

}
