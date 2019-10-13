package main

import "testing"

func Test_checkIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1",
			args{
				s: "title",
				t: "paper",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("checkIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
