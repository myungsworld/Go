package main

import (
	"log"
	"net/http"
	"time"

	"github.com/WEB4/decohandler"

	"github.com/WEB4/myapp"
)

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER1] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER1] Completed time:", time.Since(start).Milliseconds())
}

func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER2] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER2] Completed time:", time.Since(start).Milliseconds())
}

func NewHandler() http.Handler {
	mux := myapp.NewHandler()
	// Decorate Pattern
	h := decohandler.NewDecoHandler(mux, logger)
	h2 := decohandler.NewDecoHandler(h, logger2)
	return h2
}
func main() {
	mux := NewHandler()
	http.ListenAndServe(":3000", mux)
}
