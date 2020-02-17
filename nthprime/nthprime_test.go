package nthprime

import "testing"

func Test_isPrime(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: string("is prime"),
			args: args{
			num: 6,
			},
			want: false},
		{name: string("is prime"),
			args: args{
				num: 5,
			},
			want: true},
		{name: string("is prime"),
			args: args{
				num: 7,
			},
			want: true},
		{name: string("is prime"),
			args: args{
				num: 11,
			},
			want: true},
		{name: string("is prime"),
			args: args{
				num: 13,
			},
			want: true},
		{name: string("is prime"),
			args: args{
				num: 17,
			},
			want: true},
		{name: string("is prime"),
			args: args{
				num: 19,
			},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(tt.args.num); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nthPrime(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: string("nth prime"),
			args: args{number:6},
			want: 13},
		{name: string("nth prime"),
			args: args{number:10001},
			want: 104743},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthPrime(tt.args.number); got != tt.want {
				t.Errorf("nthPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}