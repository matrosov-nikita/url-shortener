package storage

type Storage interface {
	Add(url string) error
}

type Item struct {
	URL        string `json:"url"`
	ShortedURL string `json:"shorted"`
}
