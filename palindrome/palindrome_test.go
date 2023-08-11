package palindrome

import "testing"

func TestPalindrome(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"racecar", true},
		{"hello", false},
		{"", true},
		{"a", true},
		{"ab", false},
		{"aba", true},
	}

	for _, tc := range testCases {
		actual := Palindrome(tc.input)
		if actual != tc.expected {
			t.Errorf("Palindrome(%q) = %v, expected %v", tc.input, actual, tc.expected)
		}
	}
}
