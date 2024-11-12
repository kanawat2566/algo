package dynamicprogramming

import (
	"fmt"
	"testing"

	helper "github.com/kanawat2566/go-algo/Helper"
)

func TestFibo(t *testing.T) {

	tests := []struct {
		name string
		args int
		want int64
	}{
		{name: "case 1",
			args: 0,
			want: 0,
		},
		{name: "case 2",
			args: 1,
			want: 1,
		},
		{name: "case 3",
			args: 2,
			want: 1,
		},
		{name: "case 4",
			args: 3,
			want: 2,
		},
		{name: "case 5",
			args: 4,
			want: 3,
		},
		{name: "case 6",
			args: 5,
			want: 5,
		},
		{name: "case 7",
			args: 6,
			want: 8,
		},
		{name: "case 100",
			args: 100,
			want: 3736710778780434371,
		},
		{name: "case 100",
			args: 100,
			want: 3736710778780434371,
		},
		{name: "case 1000",
			args: 1000,
			want: 817770325994397771,
		},
		{name: "case 5000",
			args: 5000,
			want: 535601498209671957,
		},
	}

	before2 := helper.MemoryUsed()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FiboButtonUp(tt.args); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
	after2 := helper.MemoryUsed()
	fmt.Printf("Memory used by function FiboButtonUp: %d bytes\n", after2-before2)

}

func TestFiboV2(t *testing.T) {

	tests := []struct {
		name string
		args int
		want int64
	}{
		{name: "case 1",
			args: 0,
			want: 0,
		},
		{name: "case 2",
			args: 1,
			want: 1,
		},
		{name: "case 3",
			args: 2,
			want: 1,
		},
		{name: "case 4",
			args: 3,
			want: 2,
		},
		{name: "case 5",
			args: 4,
			want: 3,
		},
		{name: "case 6",
			args: 5,
			want: 5,
		},
		{name: "case 7",
			args: 6,
			want: 8,
		},
		{name: "case 100",
			args: 100,
			want: 3736710778780434371,
		},
		{name: "case 100",
			args: 100,
			want: 3736710778780434371,
		},
		{name: "case 1000",
			args: 1000,
			want: 817770325994397771,
		},
		{name: "case 5000",
			args: 5000,
			want: 535601498209671957,
		},
	}

	before := helper.MemoryUsed()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FiboButtonUpV2(tt.args); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
	after := helper.MemoryUsed()
	fmt.Printf("Memory used by function FiboButtonUpV1: %d bytes\n", after-before)

}

func TestFiboV3(t *testing.T) {

	tests := []struct {
		name string
		args int
		want int64
	}{
		{name: "case 1",
			args: 0,
			want: 0,
		},
		{name: "case 2",
			args: 1,
			want: 1,
		},
		{name: "case 3",
			args: 2,
			want: 1,
		},
		{name: "case 4",
			args: 3,
			want: 2,
		},
		{name: "case 5",
			args: 4,
			want: 3,
		},
		{name: "case 6",
			args: 5,
			want: 5,
		},
		{name: "case 7",
			args: 6,
			want: 8,
		},
		{name: "case 100",
			args: 100,
			want: 3736710778780434371,
		},
		{name: "case 100",
			args: 100,
			want: 3736710778780434371,
		},
		{name: "case 1000",
			args: 1000,
			want: 817770325994397771,
		},
		{name: "case 5000",
			args: 5000,
			want: 535601498209671957,
		},
	}

	before := helper.MemoryUsed()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FiboButtonUpV3(tt.args); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
	after := helper.MemoryUsed()
	fmt.Printf("Memory used by function FiboButtonUpV3: %d bytes\n", after-before)

}
