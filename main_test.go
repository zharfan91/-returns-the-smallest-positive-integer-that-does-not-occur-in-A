package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		want int
	}{
		{
			name: "sample with gap in middle",
			a:    []int{1, 3, 6, 4, 1, 2},
			want: 5,
		},
		{
			name: "consecutive positives",
			a:    []int{1, 2, 3},
			want: 4,
		},
		{
			name: "only negatives",
			a:    []int{-1, -3},
			want: 1,
		},
		{
			name: "missing one",
			a:    []int{2, 3, 4},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Solution(tt.a); got != tt.want {
				t.Fatalf("Solution(%v) = %d, want %d", tt.a, got, tt.want)
			}
		})
	}
}
