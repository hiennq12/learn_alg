package main

import (
	"reflect"
	"testing"
)

func TestLongestSubString(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"qqqswedaasassasaasasas", 6},
	}

	for _, tt := range tests {
		got := lengthOfLongestSubstring(tt.input)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("lengthOfLongestSubstring(%v) = %v; want %v", tt.input, got, tt.expected)
		}
	}

}

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	l := 0
	mapChar := make(map[string]int)
	for r, ch := range s {
		if preIdx, ok := mapChar[string(ch)]; ok && preIdx >= l {
			// ky tu da gap, con nam trong cua so
			l = preIdx + 1 // tang left vuot qua index do
		}

		mapChar[string(ch)] = r
		if maxLen < r-l+1 {
			maxLen = r - l + 1
		}

	}
	return maxLen
}
