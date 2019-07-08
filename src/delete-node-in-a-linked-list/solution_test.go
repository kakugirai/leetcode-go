package main

import (
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		node *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"1",
			args{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}},
			&ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteNode(tt.args.node)
			if !reflect.DeepEqual(tt.args.node, tt.want) {
				t.Errorf("numDecodings() = %v, want %v", tt.args.node, tt.want)
			}
		})
	}
}
