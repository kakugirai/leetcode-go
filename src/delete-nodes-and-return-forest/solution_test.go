package main

import (
	"reflect"
	"testing"
)

func Test_delNodes(t *testing.T) {
	type args struct {
		root     *TreeNode
		toDelete []int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			"1",
			args{
				root:     &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}},
				toDelete: []int{5, 3},
			},
			[]*TreeNode{{6, nil, nil}, {7, nil, nil}, {1, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, nil}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := delNodes(tt.args.root, tt.args.toDelete); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("delNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
