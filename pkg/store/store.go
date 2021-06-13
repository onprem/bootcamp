package store

type Store interface {
	Get(string) string
	Set(string, string)
}
