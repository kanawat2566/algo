package dynamicprogramming

import (
	"reflect"
	"testing"
)

func TestKnapsack(t *testing.T) {
	type args struct {
		capacity int
		weights  []int
		values   []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "test case1",
			args: args{
				capacity: 50,
				weights:  []int{20, 20, 30, 30, 10},
				values:   []int{20, 30, 30, 40, 50},
			},
			want: [][]int{{2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Knapsack(tt.args.capacity, tt.args.weights, tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Knapsack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKnapsackV2(t *testing.T) {
	type args struct {
		W    int
		idx  int
		w    []int
		p    []int
		pick []bool
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantArr []int
	}{
		// TODO: Add test cases.
		{
			name: "test case1",
			args: args{
				W:    50,
				idx:  4,
				w:    []int{20, 20, 30, 30, 10},
				p:    []int{20, 30, 30, 40, 50},
				pick: []bool{true, true, true, true, true},
			},
			want:    99,
			wantArr: []int{1, 2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			KnapsackV2(tt.args.W, tt.args.idx, tt.args.w, tt.args.p, tt.args.pick)
			if Max != tt.want {
				t.Errorf("KnapsackV2() = %v, want %v", Max, tt.want)
			}
			if !reflect.DeepEqual(MaxArr, tt.wantArr) {
				t.Errorf("KnapsackV2() = %v, want %v", MaxArr, tt.wantArr)
			}
		})
	}
}
