package main

import "testing"

func Test_countBattleships(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				[][]byte{
					{'X', '.', '.', 'X'},
					{'.', '.', '.', 'X'},
					{'.', '.', '.', 'X'},
					{'.', '.', '.', 'X'},
				},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBattleships(tt.args.board); got != tt.want {
				t.Errorf("countBattleships() = %v, want %v", got, tt.want)
			}
		})
	}
}
