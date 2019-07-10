package main

import (
	"reflect"
	"testing"
)

func Test_stack_Len(t *testing.T) {
	tests := []struct {
		name string
		stk  *stack
		want int
	}{
		{
			"1",
			&stack{1, 2},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stk.Len(); got != tt.want {
				t.Errorf("stack.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_Empty(t *testing.T) {
	tests := []struct {
		name string
		stk  *stack
		want bool
	}{
		{
			"1",
			new(stack),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stk.Empty(); got != tt.want {
				t.Errorf("stack.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_stack_Peek(t *testing.T) {
	tests := []struct {
		name string
		stk  *stack
		want int
	}{
		{
			"1",
			&stack{1, 2, 3},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.stk.Peek(); got != tt.want {
				t.Errorf("stack.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want MyQueue
	}{
		{
			"1",
			 MyQueue{
			 	new(stack),
			 	new(stack),
			 },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyQueue_Peek(t *testing.T) {
	type fields struct {
		PushStack *stack
		PopStack  *stack
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"1",
			fields{
				&stack{},
				&stack{1},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MyQueue{
				PushStack: tt.fields.PushStack,
				PopStack:  tt.fields.PopStack,
			}
			if got := q.Peek(); got != tt.want {
				t.Errorf("MyQueue.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyQueue_Empty(t *testing.T) {
	type fields struct {
		PushStack *stack
		PopStack  *stack
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"1",
			fields{
				&stack{},
				&stack{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MyQueue{
				PushStack: tt.fields.PushStack,
				PopStack:  tt.fields.PopStack,
			}
			if got := q.Empty(); got != tt.want {
				t.Errorf("MyQueue.Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

