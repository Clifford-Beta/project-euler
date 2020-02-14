package palindromeproduct

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"reverse",
			args {"racecar"},
			"racecar",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.str); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_is_palindrome(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"isPalindrome",
			args {"racecar"},
			true,
		},
		{
			"isPalindrome",
			args {"cool"},
			false,
		},
		{
			"isPalindrome",
			args {"xxx"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.str); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPalindromeProduct(t *testing.T) {
	tests := []struct {
		name  string
		want  int
	}{
		{
			"palindrome_product",
			906609,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findPalindromeProduct()
			if got != tt.want {
				t.Errorf("findPalindromeProduct() got = %v, want %v", got, tt.want)
			}
		})
	}
}