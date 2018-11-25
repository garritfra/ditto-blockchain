package main

import (
	"net/http"
	"strings"
)

// ServeHTTP starts the HTTP server on port 8080
func StartServer() {
	http.HandleFunc("/", sayHello)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	w.Write([]byte(message))
}
