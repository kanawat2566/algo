package ex2

import (
	"reflect"
	"testing"
)

func TestFindRectangle(t *testing.T) {
	type args struct {
		A [][]bool
	}
	tests := []struct {
		name    string
		args    args
		wantRes [2]PointX
	}{
		// TODO: Add test cases.
		{
			name: "case1",
			args: args{[][]bool{
				{false, false, false},
				{false, true, false},
				{false, false, false},
			}},
			wantRes: [2]PointX{{1, 1}, {1, 1}},
		},
		{
			name: "case2",
			args: args{[][]bool{
				{true, false, false},
				{false, true, false},
				{false, false, true},
			}},
			wantRes: [2]PointX{{0, 0}, {2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := FindRectangle(tt.args.A); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("FindRectangle() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
