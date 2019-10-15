package main

import (
	"reflect"
	"testing"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	headA := &ListNode{2, &ListNode{4, nil}}
	headB := &ListNode{9, &ListNode{1, headA}}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"1",
			args{
				headA: headA,
				headB: headB,
			},
			&ListNode{2, &ListNode{4, nil}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
