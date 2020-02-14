package prime

import "testing"

func Test_findLargestPrimeFactor(t *testing.T) {
	type args struct {
		prime int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"largest prime factor for 13195",
			args{13195},
			29,
		},
		{
			"largest prime factor for 600851475143",
			args{600851475143},
			6857,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLargestPrimeFactor(tt.args.prime); got != tt.want {
				t.Errorf("findLargestPrimeFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}