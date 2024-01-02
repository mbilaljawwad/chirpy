package main

import (
	"fmt"
	"net/http"
)

const port = ":8888"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/", func(writer http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/api/" {
			http.NotFound(writer, req)
			return
		}
		fmt.Fprintf(writer, "Welcome to the home page!")
	})

	corsMux := middlewareCors(mux)

	server := http.Server{
		Handler: corsMux,
		Addr: port,
	}

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("HTTP server ListenAndServe: %v", err)
	}
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(writer http.ResponseWriter, req *http.Request) {
			writer.Header().Set("Access-Control-Allow-Origin", "*")
			writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
			writer.Header().Set("Access-Control-Allow-Headers", "*")

			if req.Method == "OPTIONS" {
				writer.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(writer, req)
		})
}