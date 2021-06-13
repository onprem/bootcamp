package main

import (
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}

func sayPong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Custom-Header", "ping-pong")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("pong!\n"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", sayHello)
	mux.HandleFunc("/ping", sayPong)

	http.ListenAndServe("0.0.0.0:8080", mux)
}
