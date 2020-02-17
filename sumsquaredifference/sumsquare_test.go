package sumsquaredifference

import "testing"

func Test_sumSquareDifference(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: string("sum of square"),
			args: args{n:10} ,
			want: 2640},
		{name: string("sum of square"),
			args: args{n:100} ,
			want: 25164150},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumSquareDifference(tt.args.n); got != tt.want {
				t.Errorf("sumSquareDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}