package perfectnumber

import (
	"reflect"
	"testing"
)

func Test_perfectnumbers(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{a: 1, b: 7},
			want: []int{6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Perfectnumbers(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("perfectnumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPerfactNumber(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "case 1",
			args: args{x: 5},
			want: true,
		},
		{
			name: "case 2",
			args: args{x: 6},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPerfactNumber(tt.args.x); got != tt.want {
				t.Errorf("IsPerfactNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
