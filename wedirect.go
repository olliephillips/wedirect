package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	// start HTTP server and redirect to www of requested
	// domain with 301 redirect

	// handler for everything
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Host, ("www.")) {
			// not handling www on this server
			http.NotFound(w, r)
		} else {
			// redirect to www version
			http.Redirect(w, r, "http://www."+r.Host+r.URL.String(), http.StatusMovedPermanently)
		}
	})

	// specify http server defaults
	srv := &http.Server{
		Addr:         ":80",
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("Starting HTTP server")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
