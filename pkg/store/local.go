package store

import (
	"encoding/gob"
	"io"
	"os"
)

type LocalStore struct {
	db   map[string]string
	path string
}

func NewLocal(path string) (*LocalStore, error) {
	ls := &LocalStore{
		db:   make(map[string]string),
		path: path,
	}

	return ls, ls.load()
}

func (s *LocalStore) Set(key string, value string) {
	s.db[key] = value
	s.save()
}

func (s *LocalStore) Get(key string) string {
	return s.db[key]
}

func (s *LocalStore) save() error {
	f, err := os.Create(s.path)
	if err != nil {
		return err
	}
	defer f.Close()

	return gob.NewEncoder(f).Encode(s.db)
}

func (s *LocalStore) load() error {
	f, err := os.Open(s.path)
	if err != nil {
		// File does not exists, skip returning error.
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer f.Close()

	err = gob.NewDecoder(f).Decode(&s.db)
	if err == io.EOF {
		s.db = make(map[string]string)
		err = nil
	}
	return err
}
