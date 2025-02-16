package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/raflinoob132/go-notes/models"
	"github.com/raflinoob132/go-notes/query"
)

func GetFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	// Ambil user_id dari context (ditambahkan oleh AuthMiddleware)
	userID := r.Context().Value("userID").(uint)

	// Parse payload
	var payload models.SearchPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, `{"message": "Invalid request payload", "error": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	// Panggil fungsi query untuk mendapatkan data favorit
	favorites, total, err := query.GetFavorites(userID, payload)
	if err != nil {
		http.Error(w, `{"message": "Failed to fetch favorites", "error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}

	// Buat response JSON
	response := map[string]interface{}{
		"favorites": favorites,
		"pagination": map[string]interface{}{
			"current_page": payload.Page,
			"total_pages":  (total + int64(payload.Limit) - 1) / int64(payload.Limit),
			"total_items":  total,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
