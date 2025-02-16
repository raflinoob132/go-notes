package controllers

import (
	"encoding/json"
	"net/http"
	"replika-golang-fiber/initialize"
	"replika-golang-fiber/models"
	"time"

	"gorm.io/gorm"
	// Sesuaikan dengan struktur project Anda
)

type FavoriteRequest struct {
	NoteID uint `json:"note_id"`
}

func PostFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	// Ambil user_id dari context (ditambahkan oleh AuthMiddleware)
	userID := r.Context().Value("userID").(uint)

	// Decode request body
	var favoriteReq FavoriteRequest
	err := json.NewDecoder(r.Body).Decode(&favoriteReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validasi apakah note_id ada
	var note models.Note
	if err := initialize.DB.First(&note, favoriteReq.NoteID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Note not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Error checking note", http.StatusInternalServerError)
		return
	}

	// Cek apakah favorit sudah ada untuk user dan note ini
	var existingFavorite models.Favorite
	err = initialize.DB.Where("user_id = ? AND note_id = ?", userID, favoriteReq.NoteID).First(&existingFavorite).Error
	if err == nil {
		http.Error(w, "Favorite already exists", http.StatusBadRequest)
		return
	}

	// Simpan ke tabel favorites
	favorite := models.Favorite{
		UserID:    userID,
		NoteID:    favoriteReq.NoteID,
		CreatedAt: time.Now(),
	}

	if err := initialize.DB.Create(&favorite).Error; err != nil {
		http.Error(w, "Failed to save favorite", http.StatusInternalServerError)
		return
	}

	// Response success
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Favorite added successfully"})
}
