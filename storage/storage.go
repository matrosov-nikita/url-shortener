// Package storage contains interfaces for Cacher and Storage
package storage

// Cacher interface
type Cacher interface {
	// GetUniqueKey return value of given key from redis.
	GetUniqueKey() (int64, error)
	// IncrUniqueKey method increments given key from redis.
	IncrUniqueKey() error
}

// Storage interface
type Storage interface {
	// AddURL method adds short and origin versions of URL to database.
	AddURL(short, origin string) error
	// GetURL method returns original version of URL by short version.
	GetURL(short string) (string, error)
	// Count method returns total cound of rows in the database.
	Count() (int64, error)
}
