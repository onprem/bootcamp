package store

type InMemStore struct {
	db map[string]string
}

func NewInMem() *InMemStore {
	return &InMemStore{
		db: make(map[string]string),
	}
}

func (s *InMemStore) Set(key string, value string) {
	s.db[key] = value
}

func (s *InMemStore) Get(key string) string {
	return s.db[key]
}
