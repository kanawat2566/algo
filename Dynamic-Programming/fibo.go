package dynamicprogramming

var tempTable [10001]int64

func Fibo(n int64) int64 {
	if n == 1 || n == 0 {
		return n
	}

	if tempTable[n] >= 2 {
		return tempTable[n]
	}

	result := Fibo(n-1) + Fibo(n-2)
	tempTable[n] = result
	return result
}

var tempTable2 [10001]int64

func FiboButtonUp(n int) int64 {
	tempTable2[0] = 0
	tempTable2[1] = 1
	var i int
	for i = 2; i <= n; i++ {
		tempTable2[i] = tempTable2[i-1] + tempTable2[i-2]
	}
	return tempTable2[n]
}

func FiboButtonUpV2(n int) int64 {
	var j1, j2, result int64
	if n == 0 || n == 1 {
		return int64(n)
	}
	j1 = 0
	j2 = 1
	var flag bool = true
	var i int
	for i = 2; i <= n; i++ {
		result = j1 + j2

		if flag {
			j1 = result
			flag = false
		} else {
			j2 = result
			flag = true
		}
	}

	return result
}

func FiboButtonUpV3(n int) int64 {
	var f1, f2, result int64
	if n == 0 || n == 1 {
		return int64(n)
	}
	f2 = 0
	f1 = 1
	var i int
	for i = 2; i <= n; i++ {
		result = f2 + f1
		f2 = f1
		f1 = result

	}

	return result
}
