package storage

type Storage map[string]string

func NewStorage() *Storage {
	storage := make(Storage)
	return &storage
}

func (s Storage) Add(full string) string {
	s[defaultShortURL] = full
	return defaultShortURL
}

func (s Storage) Get(short string) string {
	return s[short]
}
