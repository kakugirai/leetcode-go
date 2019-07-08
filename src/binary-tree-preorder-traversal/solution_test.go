package main

import (
	"reflect"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"1",
			args{
				&TreeNode{1,
					&TreeNode{2,
						&TreeNode{4,
							&TreeNode{7, nil, nil},
							nil},
						&TreeNode{5, nil, nil}},
					&TreeNode{3,
						&TreeNode{6,
							&TreeNode{8, nil, nil},
							&TreeNode{9, nil, nil}},
						nil}},
			},
			[]int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

