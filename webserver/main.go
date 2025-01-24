package main

import (
	"net/http"
	"strconv"
	"sync/atomic"
)

func main() {
	cfg := &apiConfig{}
	servmux := http.NewServeMux()
	server := http.Server{
		Handler: servmux,
		Addr:    ":8080",
	}
	fileServerHandler := http.StripPrefix("/app", http.FileServer(http.Dir(".")))
	servmux.Handle("/app/", cfg.middlewareMetricsInc(fileServerHandler))
	servmux.HandleFunc("/healthz", handleStatus)
	hitcountHandler := cfg.handleHitCount
	servmux.HandleFunc("/metrics", hitcountHandler)
	hitcountResetHandler := cfg.resetMetrics
	servmux.HandleFunc("/reset", hitcountResetHandler)
	server.ListenAndServe()
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

func (cfg *apiConfig) handleHitCount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("Hits: " + strconv.Itoa(int(cfg.fileserverHits.Load()))))
}

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}

func (cfg *apiConfig) resetMetrics(w http.ResponseWriter, r *http.Request) {
	cfg.fileserverHits.Store(0)
}

// func (cfg *apiConfig) resetMetrics(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		cfg.fileserverHits.Store(0)
// 		next.ServeHTTP(w, r)
// 	})
// }
