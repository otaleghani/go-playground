package main

import (
  "fmt"
	"net/http"
  "github.com/otaleghani/go-playground/embed-ui/ui"
)

// This embeds it directly 
// //go:embed ui/dist/*
// var content embed.FS
// 
// func main() {
// 	distFS, _ := fs.Sub(content, "ui/dist")
// 	http.Handle("/", http.FileServer(http.FS(distFS)))
// 
// 	// Start the HTTP server
// 	port := "8080"
// 	println("Server is running at http://localhost:" + port)
// 	http.ListenAndServe(":"+port, nil)
// }

func main() {
  fs, err := ui.Ui()
  if err != nil {
    fmt.Println(err)
  }

 	http.Handle("/", http.FileServer(http.FS(fs)))
 
 	// Start the HTTP server
 	port := "8080"
 	println("Server is running at http://localhost:" + port)
 	http.ListenAndServe(":"+port, nil)
}
