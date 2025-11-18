package handlers

import (
	"net/http"
	"url-shortner/config"
)

// RedirectHandler handles GET /<shortCode>
func RedirectHandler(w http.ResponseWriter, r *http.Request) {

	// Get everything after the first slash ("/aB9xYz")
	shortCode := r.URL.Path[1:]

	if shortCode == "" {
		http.Error(w, "Short code missing", http.StatusBadRequest)
		return
	}

	originalURL, err := config.RedisClient.Get(config.Ctx, shortCode).Result()
	if err != nil {
		http.Error(w, "Short URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)

}