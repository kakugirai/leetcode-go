package main

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"1",
			args{2.0, 3},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
