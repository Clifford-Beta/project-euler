package multiples

import "testing"

func Test_findSumOfAllMultiplesNumbers(t *testing.T) {
	type args struct {
		numbers    []int
		upperLimit int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "multiples of 3 or 5 below 10",
			args: args{
				numbers:    []int{3, 5},
				upperLimit: 10,
			},
			want: 23},
		{name: "multiples of 3 or 5 below 10",
			args: args{
				numbers:    []int{3, 5,7},
				upperLimit: 10,
			},
			want: 30},
		{name: "multiples of 3 or 5 below 15",
			args: args{
				numbers:    []int{3, 5},
				upperLimit: 15,
			},
			want: 45},
		{name: "multiples of 3 or 5 below 20",
			args: args{
				numbers:    []int{3, 5},
				upperLimit: 20,
			},
			want: 78},
		{name: "multiples of 3 or 5 below 1000",
			args: args{
				numbers:    []int{3, 5},
				upperLimit: 1000,
			},
			want: 233168},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSumOfAllMultiplesNumbers(tt.args.numbers, tt.args.upperLimit); got != tt.want {
				t.Errorf("findSumOfAllMultiplesNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}