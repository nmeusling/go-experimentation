package main

import (
	"net/http"
)

func main() {
	servmux := http.NewServeMux()
	server := http.Server{
		Handler: servmux,
		Addr:    ":8080",
	}
	servmux.Handle("/", http.FileServer(http.Dir(".")))
	server.ListenAndServe()
}
