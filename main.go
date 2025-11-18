package main

import (
	"fmt"
	"log"
	"net/http"
	"url-shortner/config"
	"url-shortner/handlers"
)

func main() {
	fmt.Println("Starting URL Shortner...")
	config.InitRedis()

	http.HandleFunc("/shorten", handlers.ShortenHandler)
	http.HandleFunc("/", handlers.RedirectHandler)

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
