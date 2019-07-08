package main

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name string
		args args
		want MyCircularQueue
	}{
		{
			"1",
			args{3},
			MyCircularQueue{[]int{0, 0, 0}, 3, -1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_EnQueue(t *testing.T) {
	type fields struct {
		elements []int
		capacity int
		front    int
		rear     int
	}
	type args struct {
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"1",
			fields{
				[]int{0, 0, 0},
				3,
				-1,
				-1,
			},
			args{1},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cq := &MyCircularQueue{
				elements: tt.fields.elements,
				capacity: tt.fields.capacity,
				front:    tt.fields.front,
				rear:     tt.fields.rear,
			}
			if got := cq.EnQueue(tt.args.value); got != tt.want {
				t.Errorf("MyCircularQueue.EnQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_DeQueue(t *testing.T) {
	type fields struct {
		elements []int
		capacity int
		front    int
		rear     int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"1",
			fields{
				[]int{0, 0, 0},
				3,
				-1,
				-1,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cq := &MyCircularQueue{
				elements: tt.fields.elements,
				capacity: tt.fields.capacity,
				front:    tt.fields.front,
				rear:     tt.fields.rear,
			}
			if got := cq.DeQueue(); got != tt.want {
				t.Errorf("MyCircularQueue.DeQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_Front(t *testing.T) {
	type fields struct {
		elements []int
		capacity int
		front    int
		rear     int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"1",
			fields{
				[]int{1, 1, 1},
				3,
				-1,
				-1,
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cq := &MyCircularQueue{
				elements: tt.fields.elements,
				capacity: tt.fields.capacity,
				front:    tt.fields.front,
				rear:     tt.fields.rear,
			}
			if got := cq.Front(); got != tt.want {
				t.Errorf("MyCircularQueue.Front() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_Rear(t *testing.T) {
	type fields struct {
		elements []int
		capacity int
		front    int
		rear     int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			"1",
			fields{
				[]int{1, 1, 1},
				3,
				-1,
				-1,
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cq := &MyCircularQueue{
				elements: tt.fields.elements,
				capacity: tt.fields.capacity,
				front:    tt.fields.front,
				rear:     tt.fields.rear,
			}
			if got := cq.Rear(); got != tt.want {
				t.Errorf("MyCircularQueue.Rear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_IsEmpty(t *testing.T) {
	type fields struct {
		elements []int
		capacity int
		front    int
		rear     int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"1",
			fields{
				[]int{0, 0, 0},
				3,
				-1,
				-1,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cq := &MyCircularQueue{
				elements: tt.fields.elements,
				capacity: tt.fields.capacity,
				front:    tt.fields.front,
				rear:     tt.fields.rear,
			}
			if got := cq.IsEmpty(); got != tt.want {
				t.Errorf("MyCircularQueue.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyCircularQueue_IsFull(t *testing.T) {
	type fields struct {
		elements []int
		capacity int
		front    int
		rear     int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"1",
			fields{
				[]int{0, 0, 0},
				3,
				-1,
				-1,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cq := &MyCircularQueue{
				elements: tt.fields.elements,
				capacity: tt.fields.capacity,
				front:    tt.fields.front,
				rear:     tt.fields.rear,
			}
			if got := cq.IsFull(); got != tt.want {
				t.Errorf("MyCircularQueue.IsFull() = %v, want %v", got, tt.want)
			}
		})
	}
}
