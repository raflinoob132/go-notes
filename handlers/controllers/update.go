package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/raflinoob132/go-notes/initialize"
	"github.com/raflinoob132/go-notes/models"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func UpdateNoteHandler(w http.ResponseWriter, r *http.Request) {
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

	// Ambil ID dari URL parameter (misalnya: /notes/{id})
	id := chi.URLParam(r, "id")

	// Cek apakah note dengan ID tersebut ada di database
	var existingNote models.Note
	if err := initialize.DB.Where("id = ?", id).First(&existingNote).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, `{"error":"Catatan tidak ditemukan"}`, http.StatusNotFound)
		} else {
			http.Error(w, `{"error":"Gagal memeriksa catatan di database"}`, http.StatusInternalServerError)
		}
		return
	}

	// Update data catatan
	existingNote.Title = payload.Title
	existingNote.Content = payload.Content
	existingNote.Category = payload.Category
	existingNote.UpdatedAt = time.Now()

	// Simpan perubahan ke database
	if result := initialize.DB.Save(&existingNote); result.Error != nil {
		http.Error(w, `{"error":"Gagal memperbarui data ke database"}`, http.StatusInternalServerError)
		return
	}

	// Berikan response sukses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Note updated successfully",
		"note":    existingNote,
	})
}
