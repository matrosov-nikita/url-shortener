package shortener

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const base = len(alphabet)

// Encode URL using given incremental key
func Encode(url string, key int) string {
	var shortURL string

	for key > 0 {
		shortURL = shortURL + string(alphabet[key%base])
		key /= base
	}

	return reverse(shortURL)
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
