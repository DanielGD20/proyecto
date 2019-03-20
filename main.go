package main

import (
	"net/http"
	"os"
)

func main() {
	fileServer := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fileServer))
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
