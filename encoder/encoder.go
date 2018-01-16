// Package encoder provides implementation for encoding based on given number.
package encoder

import "strings"

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const base = int64(len(alphabet))

// Encode using given number
func Encode(n int64) string {
	var encoded string

	if n == 0 {
		return string(alphabet[0])
	}

	for n > 0 {
		encoded += string(alphabet[n%base])
		n /= base
	}

	return encoded
}

// IsValidHash validates hash
func IsValidHash(hash string) bool {
	if len(hash) == 0 {
		return false
	}

	for _, c := range hash {
		if !strings.Contains(alphabet, string(c)) {
			return false
		}
	}

	return true
}
