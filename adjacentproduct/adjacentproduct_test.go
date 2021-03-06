package adjacentproduct

import (
	"reflect"
	"testing"
)

const ThousandDigitNumber = "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"

func Test_calculateProduct(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 int
	}{
		{name: string("compute product"),
			args:  args{num: string("9989")},
			want:  []int{9, 9, 8, 9},
			want1: 5832},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := calculateProduct(tt.args.num)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateProduct() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("calculateProduct() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_adjacentNDigits(t *testing.T) {
	type args struct {
		number string
		n      int
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 int
	}{
		{name: string("largest adjacent product of 4 digits"),
			args: args{
				number: ThousandDigitNumber,
				n:      4,
			},
			want:  []int{9, 9, 8, 9},
			want1: 5832},
		{name: string("largest adjacent product of 13 digits"),
			args: args{
				number: ThousandDigitNumber,
				n:      13,
			},
			want:  []int{5, 5, 7, 6, 6, 8, 9, 6, 6, 4, 8, 9, 5},
			want1: 23514624000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := adjacentNDigits(tt.args.number, tt.args.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("adjacentNDigits() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("adjacentNDigits() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
