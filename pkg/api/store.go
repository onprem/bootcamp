package api

import (
	"encoding/json"
	"net/http"
)

func (a *API) handleGet(w http.ResponseWriter, r *http.Request) {
	var inp struct {
		Key string `json:"key"`
	}

	err := json.NewDecoder(r.Body).Decode(&inp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(a.str.Get(inp.Key)))
}

// `
// {
// 	"key": "11hh",
// 	"value": "aa"
//  "id": 122,
//  "user": {
//    "name": "haha"
//  }
// }
// `

func (a *API) handleSet(w http.ResponseWriter, r *http.Request) {
	var inp struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	err := json.NewDecoder(r.Body).Decode(&inp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("X-Error", err.Error())
		w.Write([]byte("something went wrong"))
		return
	}

	a.str.Set(inp.Key, inp.Value)
	w.Write([]byte("key-val stored!\n"))
}
