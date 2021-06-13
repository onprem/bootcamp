package main

import (
	"log"
	"net/http"

	"github.com/onprem/bootcamp/pkg/api"
	"github.com/onprem/bootcamp/pkg/store"
)

func main() {
	str, err := store.NewLocal("db.gob")
	if err != nil {
		log.Fatal(err)
	}

	apy := api.New(str)

	mux := http.NewServeMux()

	apy.Register(mux)

	http.ListenAndServe(":8080", mux)
}
