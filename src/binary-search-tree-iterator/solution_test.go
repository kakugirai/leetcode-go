package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want BSTIterator
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
			BSTIterator{
				-1,
				[]int{7, 4, 2, 5, 1, 8, 6, 9, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}
