package main

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"1",
			args{
				[]byte{'a', 'b', 'd', 'e'},
			},
			[]byte{'e', 'd', 'b', 'a'},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.s)
			if !reflect.DeepEqual(tt.args.s, tt.want) {
				t.Errorf("reverseString() = %v, want %v", tt.args.s, tt.want)
			}
		})
	}
}
