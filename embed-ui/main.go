package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed ui/dist/*
var content embed.FS

func main() {
	distFS, _ := fs.Sub(content, "ui/dist")
	http.Handle("/", http.FileServer(http.FS(distFS)))

	// Start the HTTP server
	port := "8080"
	println("Server is running at http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
