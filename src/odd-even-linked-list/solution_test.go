package main

import (
	"reflect"
	"testing"
)

func Test_oddEvenList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"1",
			args{
				&ListNode{2, &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{6, &ListNode{4, &ListNode{7, nil}}}}}}},
			},
			&ListNode{2, &ListNode{3, &ListNode{6, &ListNode{7, &ListNode{1, &ListNode{5, &ListNode{4, nil}}}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddEvenList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("oddEvenList() = %v, want %v", got, tt.want)
			}
		})
	}
}
