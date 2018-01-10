package encoder

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const base = int64(len(alphabet))

// Encode URL using given incremental key
func Encode(n int64) string {
	var encoded string

	for n > 0 {
		encoded += string(alphabet[n%base])
		n /= base
	}

	return encoded
}
