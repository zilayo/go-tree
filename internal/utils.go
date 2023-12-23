package utils

import "strings"

// reverseString reverses a string.
func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// ReplaceLast replaces the last occurrence of 'old' in 's' with 'new'.
func ReplaceLast(s, old, new string) string {
	rev := reverseString(s)
	newRev := reverseString(new)

	if strings.Contains(rev, old) {
		replaced := strings.Replace(rev, old, newRev, 1)
		return reverseString(replaced)
	}
	return s
}
