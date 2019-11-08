package main

import "testing"

func Test_minimumSwap(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				s1: "xxyyxyxyxx",
				s2: "xyyxyxxxyx",
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumSwap(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("minimumSwap() = %v, want %v", got, tt.want)
			}
		})
	}
}
