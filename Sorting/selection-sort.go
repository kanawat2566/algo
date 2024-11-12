package sorting

import "math"

func Selection_sort(arr []int) []int {
	var outs []int
	var n = len(arr)
	for i := 0; i < n; i++ {
		min := arr[i]
		for j := i; j < n; j++ {
			if arr[j] < min {
				tmp := min
				min = arr[j]
				arr[j] = tmp
			}
		}
		outs = append(outs, min)

	}

	return outs
}
func Selection_sortV2(arr []int) []int {
	var outs []int
	var n = len(arr)

	var min int = 0
	for i := 0; i < n; i++ {
		min = arr[i]
		for j := i + 1; j < n; j++ {
			if arr[j] < min {
				tmp := min
				min = arr[j]
				arr[j] = tmp
			}
		}
		outs = append(outs, min)
	}

	return outs
}

func Selection_sortV3(arr []int) []int {
	pos := len(arr) - 1
	for ; pos > 0; pos-- {
		max_idx := 0
		for i := 0; i <= pos; i++ {
			if arr[i] > arr[max_idx] {
				max_idx = i
			}
		}
		arr[pos], arr[max_idx] = arr[max_idx], arr[pos]
	}
	return arr
}

// [4,5,6,2,3]
// [5,4,3,2,1]
func Heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		temp := arr[i]
		arr[i] = arr[largest]
		arr[largest] = temp

		// เรียก Heapify ซ้ำสำหรับโหนดที่สลับตำแหน่ง
		Heapify(arr, n, largest)
	}
}

func BuildMaxHeap(array []int, n int) []int {
	for i := int(math.Floor(float64(n)/2)) - 1; i >= 0; i-- {
		Heapify(array, n, i)
	}
	return array
}
func Heap_sort(arr []int) []int {
	BuildMaxHeap(arr, len(arr))
	for i := len(arr) - 1; i > 0; i-- {
		temp := arr[0]
		arr[0] = arr[i]
		arr[i] = temp
		Heapify(arr, i, 0)
	}
	return arr
}
