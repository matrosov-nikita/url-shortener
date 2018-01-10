package storage

type Cacher interface {
	GetUniqueKey(key string) (int64, error)
	SetUniqueKey(key string) error
}

type Storage interface {
	AddURL(short, origin string) error
	GetURL(short string) (string, error)
	Count() (int64, error)
}
