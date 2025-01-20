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
	fileServerHandler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	servmux.Handle("/app/", fileServerHandler)
	servmux.HandleFunc("/healthz", handleStatus)
	server.ListenAndServe()
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}
