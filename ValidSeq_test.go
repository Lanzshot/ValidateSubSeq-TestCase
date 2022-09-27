/*
2 sets of array where need to determine if the 2nd array is a subset of the 1st array.
*/

package main

import "testing"

func TestIsValidSubseq(t *testing.T) {
	type args struct {
		array    []int
		sequence []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Hello",
			args: args{
				array:    []int{5, 1, 22, 25, 6, -1, 8, 10},
				sequence: []int{25, 10},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidSubseq(tt.args.array, tt.args.sequence); got != tt.want {
				t.Errorf("IsValidSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}
