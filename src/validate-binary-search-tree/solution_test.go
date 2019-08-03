package main

import (
	"testing"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{&TreeNode{1,
				&TreeNode{2,
					&TreeNode{4,
						&TreeNode{7, nil, nil},
						nil},
					&TreeNode{5, nil, nil}},
				&TreeNode{3,
					&TreeNode{6,
						&TreeNode{8, nil, nil},
						&TreeNode{9, nil, nil}},
					nil}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
