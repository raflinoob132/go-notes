package controllers

import (
	"net/http"
	"strconv"

	"github.com/raflinoob132/go-notes/initialize"

	"github.com/go-chi/chi/v5"
)

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Delete note by ID
	//var notes models.Note
	id := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, `{"error":"Invalid ID"}`, http.StatusBadRequest)
		return
	}
	query := "DELETE FROM notes WHERE id = ?"

	if result := initialize.DB.Exec(query, intID); result.Error != nil {
		http.Error(w, `{"error":"Failed to delete note"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}
