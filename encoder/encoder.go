// Package encoder provides implementation for encoding based on given number
package encoder

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const base = int64(len(alphabet))

// TODO first short URL is empty
// Encode using given number
func Encode(n int64) string {
	var encoded string

	for n > 0 {
		encoded += string(alphabet[n%base])
		n /= base
	}

	return encoded
}
