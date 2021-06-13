package store

import (
	"encoding/gob"
	"os"
)

type LocalStore struct {
	db   map[string]string
	file *os.File
}

func NewLocal(path string) (*LocalStore, error) {
	file, err := os.Open(path)
	if err != nil {
		file, err = os.Create(path)
		if err != nil {
			return nil, err
		}
	}

	ls := &LocalStore{
		db:   make(map[string]string),
		file: file,
	}

	err = gob.NewDecoder(file).Decode(&ls.db)
	if err != nil {
		ls.db = make(map[string]string)
	}

	return ls, nil
}

func (s *LocalStore) Set(key string, value string) {
	s.db[key] = value

	_ = s.file.Truncate(0)
	_ = gob.NewEncoder(s.file).Encode(s.db)
	_ = s.file.Sync()
}

func (s *LocalStore) Get(key string) string {
	return s.db[key]
}
