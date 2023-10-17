package word

import (
	"testing"
	"unicode"
)

func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

func IsPalindromeV2(letters string) bool {
	n := len(letters) / 2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}

	return true
}

func BenchmarkIsPalindromeV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindromeV2("A man, a plan, a canal: Panama")
	}
}

func IsPalindromeV3(letters []rune) bool {
	n := len(letters) / 2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}

	return true
}

func BenchmarkIsPalindromeV3(b *testing.B) {
	b.ResetTimer()
	s := "A man, a plan, a canal: Panama"
	letters := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IsPalindromeV3(letters)
	}
}
