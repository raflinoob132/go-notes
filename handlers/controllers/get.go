package controllers

import (
	"encoding/json"
	"net/http"

	//"replika-golang-fiber/handlers/controllers/crudmodel"
	"replika-golang-fiber/models"
	"replika-golang-fiber/query"
)

func GetNotesHandler(w http.ResponseWriter, r *http.Request) {
	// Parse body request
	var payload models.SearchPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, `{"message": "Invalid request payload", "error": "`+err.Error()+`"}`, http.StatusBadRequest)
		return
	}

	notes, total, err := query.GetNotes(payload)
	if err != nil {
		http.Error(w, `{"message": "Failed to fetch notes", "error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		return
	}
	// Buat response JSON
	response := map[string]interface{}{
		"notes": notes,
		"pagination": map[string]interface{}{
			"current_page": payload.Page,
			"total_pages":  (total + int64(payload.Limit) - 1) / int64(payload.Limit),
			"total_items":  total,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
