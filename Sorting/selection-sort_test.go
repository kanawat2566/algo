package sorting

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

func TestSelection_sort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "case 1",
			args: args{[]int{1, 32, 3, 55, 5, 2, 43, 2, 3, 4, 5, 6}},
			want: []int{1, 2, 2, 3, 3, 4, 5, 5, 6, 32, 43, 55},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := memoryUsed()
			if got := Selection_sort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Selection_sort() = %v, want %v", got, tt.want)
			}
			after := memoryUsed()
			fmt.Printf("Memory used by function Selection_sort: %d bytes\n", after-before)
		})
	}
}
func TestSelection_sortV2(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "case 1",
			args: args{[]int{1, 32, 3, 55, 5, 2, 43, 2, 3, 4, 5, 6}},
			want: []int{1, 2, 2, 3, 3, 4, 5, 5, 6, 32, 43, 55},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			before := memoryUsed()
			if got := Selection_sortV2(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Selection_sortV2() = %v, want %v", got, tt.want)
			}
			after := memoryUsed()
			fmt.Printf("Memory used by function Selection_sortV2: %d bytes\n", after-before)
		})
	}
}

func TestSelection_sortV3(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "case 1",
			args: args{[]int{1, 32, 3, 55, 5, 2, 43, 2, 3, 4, 5, 6}},
			want: []int{1, 2, 2, 3, 3, 4, 5, 5, 6, 32, 43, 55},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			before := memoryUsed()
			if got := Selection_sortV3(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Selection_sortV3() = %v, want %v", got, tt.want)
			}
			after := memoryUsed()
			fmt.Printf("Memory used by function Selection_sortV3: %d bytes\n", after-before)
		})
	}
}

func memoryUsed() uint64 {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	return memStats.Alloc
}

func TestBuildMaxHeap(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "case 1",
			args: args{array: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			before := memoryUsed()
			if got := Heap_sort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildMaxHeap() = %v, want %v", got, tt.want)
			}
			after := memoryUsed()
			fmt.Printf("Memory used by function BuildMaxHeap: %d bytes\n", after-before)
		})
	}
}
