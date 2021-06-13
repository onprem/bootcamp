package main

import (
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"https://www.google.com",
	"https://twitter.com",
	"https://github.com",
	"https://archlinux.org",
	"https://golang.org",
	"https://www.google.com",
	"https://twitter.com",
	"https://github.com",
	"https://archlinux.org",
	"https://golang.org",
}

func req(i int, url string, wg *sync.WaitGroup, ch chan int) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("[ERROR] index: %d, url: %s", i, url)
		return
	}

	ch <- resp.StatusCode

	// log.Printf("[SUCCESS] index: %d, url: %s, status: %d", i, url, resp.StatusCode)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	ch := make(chan int)

	go func() {
		for val := range ch {
			log.Printf("[SUCCESS] status: %d", val)
		}
	}()

	go func() {
		for val := range ch {
			log.Printf("[SUCCESS2] status: %d", val)
		}
	}()

	for i, u := range urls {
		wg.Add(1)
		go req(i, u, &wg, ch)
	}

	// Blocks until done has been called for all.
	wg.Wait()

	close(ch)
}
