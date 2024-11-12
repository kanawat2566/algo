package dynamicprogramming

import "testing"

func TestMaxSubSum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name    string
		args    args
		wantSum int
		wantP   int
		wantQ   int
	}{
		// TODO: Add test cases.
		{name: "case 1",
			args:    args{arr: []int{1, 4, 2, 3}},
			wantSum: 10,
			wantP:   1,
			wantQ:   4,
		},
		{name: "case 2",
			args:    args{arr: []int{-2, -1, -3, -5}},
			wantSum: -1,
			wantP:   2,
			wantQ:   2,
		},
		{name: "case 2",
			args:    args{arr: []int{2, 3, -6, 4, -2, 3, -5, -4, 3}},
			wantSum: 5,
			wantP:   1,
			wantQ:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSum, gotP, gotQ := MaxSubSum(tt.args.arr)
			if gotSum != tt.wantSum {
				t.Errorf("MaxSubSum() gotSum = %v, want %v", gotSum, tt.wantSum)
			}
			if gotP != tt.wantP {
				t.Errorf("MaxSubSum() gotP = %v, want %v", gotP, tt.wantP)
			}
			if gotQ != tt.wantQ {
				t.Errorf("MaxSubSum() gotQ = %v, want %v", gotQ, tt.wantQ)
			}
		})
	}
}
