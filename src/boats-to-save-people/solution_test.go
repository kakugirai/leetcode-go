package main

import "testing"

func Test_numRescueBoats(t *testing.T) {
	type args struct {
		people []int
		limit  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{[]int{3, 5, 3, 4}, 5},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numRescueBoats(tt.args.people, tt.args.limit); got != tt.want {
				t.Errorf("numRescueBoats() = %v, want %v", got, tt.want)
			}
		})
	}
}
