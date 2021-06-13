package api

import (
	"net/http"

	"github.com/onprem/bootcamp/pkg/store"
)

type API struct {
	str store.Store
}

func New(str store.Store) *API {
	return &API{
		str: str,
	}
}

func (a *API) Register(mux *http.ServeMux) {
	mux.HandleFunc("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong!\n"))
	}))

	mux.HandleFunc("/get", http.HandlerFunc(a.handleGet))
	mux.HandleFunc("/set", http.HandlerFunc(a.handleSet))
}
