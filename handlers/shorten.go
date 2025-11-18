package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortner/config"
	"url-shortner/utils"

	"github.com/redis/go-redis/v9"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

// ShortenHandler handles POST /shorten
func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.URL == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	longURL := req.URL

	//checking if long URL already has a short code -> idempotency
	existingCode, err := config.RedisClient.Get(config.Ctx, "long:"+longURL).Result()

	if err == nil {
		// shortcode already exists -> return the same short URL
		shortURL := fmt.Sprintf("http://localhost:8080/%s", existingCode)
		json.NewEncoder(w).Encode(ShortenResponse{ShortURL: shortURL})
		return
	}

	// generate a new shortcode
	shortCode := utils.GenerateShortCode(6)

	// Collision handling - ensure the generated short code is not used before

	for {
		_, err := config.RedisClient.Get(config.Ctx, shortCode).Result()
		if err == redis.Nil {
			// short code is unused before -> Okay to proceed
			break
		}

		shortCode = utils.GenerateShortCodeWithSalt(longURL)
	}

	// Save in Redis :short -> long
	if err := config.RedisClient.Set(config.Ctx, shortCode, longURL,0).Err(); err != nil {
		http.Error(w, "Failed to store URL", http.StatusInternalServerError)
		return
	}

	//SAve reverse mapping: long -> short
	config.RedisClient.Set(config.Ctx, "long:"+longURL, shortCode,0)

	// Construct short URL (localhost:8080/<code>)
	shortURL := fmt.Sprintf("http://localhost:8080/%s", shortCode)

	resp := ShortenResponse{ShortURL: shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
