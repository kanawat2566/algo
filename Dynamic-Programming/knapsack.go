package dynamicprogramming

func Knapsack(capacity int, weights, values []int) [][]int {
	// Define the capacity of the knapsack
	//var capacity int = 50
	// Define the weights and values of the items
	//var weights = []int{10, 20, 30}
	//var values = []int{100, 60, 120}

	var sumWeight int
	var sumMaxValue int
	var sumMaxArr = [][]int{}
	for i := 0; i < len(weights); i++ {
		for j := i + 1; j < len(weights); j++ {
			sumWeight = weights[i] + weights[j]
			if sumWeight <= capacity {
				if values[i]+values[j] > sumMaxValue {
					sumMaxValue = values[i] + values[j]
					sumMaxArr = append(sumMaxArr, []int{i + 1, j + 1})
				}
			}
		}
	}
	return sumMaxArr[len(sumMaxArr)-1:]
}

var Max int = 0
var MaxArr []int

func KnapsackV2(W, idx int, w, p []int, pick []bool) {
	if idx == 0 {
		sum_price := 0
		sum_weight := 0
		for i := 0; i < len(pick); i++ {
			if pick[i] {
				sum_price += p[i]
				sum_weight += w[i]
				if sum_weight <= W && sum_price > Max {
					Max = sum_price
					MaxArr = append(MaxArr, i)
				}
			}
		}
		return
	}

	pick[idx] = false
	KnapsackV2(W, idx-1, w, p, pick)
	pick[idx] = true
	KnapsackV2(W, idx-1, w, p, pick)

}


func KnapsackV3(W, idx int, w, p []int, remain int) {
	

}
