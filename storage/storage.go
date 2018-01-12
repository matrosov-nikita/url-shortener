// Package storage contains interfaces for Cacher and Storage
package storage

// Cacher interface
type Cacher interface {
	GetUniqueKey(key string) (int64, error)
	IncrUniqueKey(key string) error
}

// Storage interface
type Storage interface {
	AddURL(short, origin string) error
	GetURL(short string) (string, error)
	Count() (int64, error)
}
