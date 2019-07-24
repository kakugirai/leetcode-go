package main

import (
	"reflect"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{
				[]int{5, 4, 2, 6},
			},
			[][]int{{5, 4, 2, 6}, {5, 4, 6, 2}, {5, 2, 4, 6}, {5, 2, 6, 4}, {5, 6, 4, 2}, {5, 6, 2, 4}, {4, 5, 2, 6}, {4, 5, 6, 2}, {4, 2, 5, 6}, {4, 2, 6, 5}, {4, 6, 5, 2}, {4, 6, 2, 5}, {2, 5, 4, 6}, {2, 5, 6, 4}, {2, 4, 5, 6}, {2, 4, 6, 5}, {2, 6, 5, 4}, {2, 6, 4, 5}, {6, 5, 4, 2}, {6, 5, 2, 4}, {6, 4, 5, 2}, {6, 4, 2, 5}, {6, 2, 5, 4}, {6, 2, 4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}
