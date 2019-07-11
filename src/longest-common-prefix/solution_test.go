package main

import "testing"

func Test_longestCommonPrefix(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				[]string{"flower", "flow", "flo"},
			},
			"flo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonPrefix2(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				[]string{"flower", "flow", "flo"},
			},
			"flo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix2(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonPrefix3(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				[]string{"flower", "flow", "flo"},
			},
			"flo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix3(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonPrefix4(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				[]string{"flower", "flow", "flo"},
			},
			"flo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix4(tt.args.strs); got != tt.want {
				t.Errorf("longestCommonPrefix4() = %v, want %v", got, tt.want)
			}
		})
	}
}
