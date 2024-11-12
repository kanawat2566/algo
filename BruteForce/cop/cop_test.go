package cop

import (
	"reflect"
	"testing"
)

func TestMaxDiff(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want Aws
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{arr: []int{1, 2, 3, 4}},
			want: Aws{p: 0, q: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxDiff(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MaxDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
