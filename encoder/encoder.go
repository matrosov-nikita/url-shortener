// Package encoder provides implementation for encoding based on given number.
package encoder

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
