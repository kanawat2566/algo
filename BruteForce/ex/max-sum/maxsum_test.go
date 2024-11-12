package maxsum

import "testing"

func TestMaxSumRange(t *testing.T) {
	type args struct {
		arr []int
		w   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{arr: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, w: 5},
			want: 6,
		},
		{
			name: "test2",
			args: args{arr: []int{1, 20, 20, 4, 5, 6, 100, 70, 9, 1, 100, 100, 100, 1, 1}, w: 5},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSumRange(tt.args.arr, tt.args.w); got != tt.want {
				t.Errorf("MaxSumRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
