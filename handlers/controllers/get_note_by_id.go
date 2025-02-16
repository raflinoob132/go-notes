package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/raflinoob132/go-notes/initialize"
	"github.com/raflinoob132/go-notes/models"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	// _ "github.com/go-sql-driver/mysql"
)

func GetNoteByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, `{"error":"Invalid ID"}`, http.StatusBadRequest)
		return
	}

	var note models.Note
	if err := initialize.DB.First(&note, intID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, `{"error":"Note not found"}`, http.StatusNotFound)
		} else {
			http.Error(w, `{"error":"Failed to get note"}`, http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(note)
}
