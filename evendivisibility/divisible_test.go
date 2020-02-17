package evendivisibility

import (
	"testing"
)

func TestGCD(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: string("GCD of 2 and 6"),
			args: args{
				a: 2,
				b: 6,
			},
			want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCD(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GCD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_computeLCM(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find LCM",
			args: args{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
			want: 2520,
		},
		{
			name: "find LCM",
			args: args{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,14,15,16,17,18,19,20}},
			want: 232792560,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := computeLCM(tt.args.nums...); got != tt.want {
				t.Errorf("LCM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLCM1(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "lcm",
			args: args{
				a: 2,
				b: 6,
			},
			want: 6,
		},
		{
			name: "lcm",
			args: args{
				a: 7,
				b: 9,
			},
			want: 63,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LCM(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("LCM() = %v, want %v", got, tt.want)
			}
		})
	}
}
