package dynamicprogramming

func MaxSubSum(arr []int) (sum, p, q int) {
	var tmp [3]int
	tmp[0] = arr[0]
	for i := 0; i < len(arr); i++ {
		sum := arr[i]
		for j := i + 1; j < len(arr); j++ {
			sum += arr[j] 
			if sum > tmp[0] {
				tmp[0] = sum
				tmp[1] = i
				tmp[2] = j
			}
		}
		if arr[i] > tmp[0] {
			tmp[0] = arr[i]
			tmp[1] = i
			tmp[2] = i
		}
	}
	return tmp[0], tmp[1] + 1, tmp[2] + 1
}
