package main
//package evenfibonnaci

import "testing"

//func TestFibonacci(t *testing.T) {
//	type args struct {
//		n uint
//	}
//	tests := []struct {
//		name string
//		args args
//		want uint64
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := Fibonacci(tt.args.n); got != tt.want {
//				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestFibonacci(t *testing.T) {
	data := []struct {
		n    uint
		want uint64
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {10, 55}, {42, 267914296},
	}

	for _, d := range data {
		if got := Fibonacci(d.n); got != d.want {
			t.Errorf("Invalid Fibonacci value for N: %d, got: %d, want: %d", d.n, got, d.want)
		}
	}
}