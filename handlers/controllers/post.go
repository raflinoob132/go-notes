package controllers

import (
	"encoding/json"
	"net/http"

	//"replika-golang-fiber/controllers/crudmodel"
	"replika-golang-fiber/initialize"
	"replika-golang-fiber/models"
	"time"

	"gorm.io/gorm"
)

func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var payload *models.PostPayload

	// Decode body request
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, `{"error": "Invalid payload"}`, http.StatusBadRequest)
		return
	}

	// Validasi input
	if payload.Title == "" {
		http.Error(w, `{"error":"Judul tidak boleh kosong"}`, http.StatusBadRequest)
		return
	}

	// Cek apakah judul sudah ada di database
	var existingNote models.Note
	if err := initialize.DB.Where("title = ?", payload.Title).First(&existingNote).Error; err == nil {
		http.Error(w, `{"error":"Judul yang sama sudah ada"}`, http.StatusBadRequest)
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		http.Error(w, `{"error":"Gagal memeriksa judul di database"}`, http.StatusInternalServerError)
		return
	}

	// Buat catatan baru
	now := time.Now()
	newNote := models.Note{
		Title:     payload.Title,
		Content:   payload.Content,
		Category:  payload.Category,
		CreatedAt: now,
	}

	if result := initialize.DB.Create(&newNote); result.Error != nil {
		http.Error(w, `{"error":"Gagal menyimpan data ke database"}`, http.StatusInternalServerError)
		return
	}

	// Berikan response sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Note created successfully",
		"notes":   newNote,
	})
}
