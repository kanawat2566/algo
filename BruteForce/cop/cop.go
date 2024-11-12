package cop

import (
	"math"
)

type Aws struct {
	p int
	q int
}

func MaxDiff(arr []int) Aws {
	maxDiff := 0
	var result Aws

	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			diff := int(math.Abs(float64(arr[i] - arr[j])))
			if diff > maxDiff {
				maxDiff = diff
				result = Aws{i, j}
			}
		}
	}
	return result
}
