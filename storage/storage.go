package storage

type Storage interface {
	Add(url string) error
	GetUniqueKey() (int, error)
}

type Item struct {
	URL        string `json:"url"`
	ShortedURL string `json:"shorted"`
}
