package query

import (
	"replika-golang-fiber/initialize"
	"replika-golang-fiber/models"
)

func GetFavorites(userID uint, payload models.SearchPayload) ([]map[string]interface{}, int64, error) {
	// Default nilai pagination
	if payload.Page < 1 {
		payload.Page = 1
	}
	if payload.Limit < 1 {
		payload.Limit = 10
	}

	offset := (payload.Page - 1) * payload.Limit

	// Query untuk mendapatkan favorit
	query := `
		SELECT notes.id, notes.title
		FROM favorites
		JOIN notes ON notes.id = favorites.note_id
		WHERE favorites.user_id = ?
	`
	var args []interface{}
	args = append(args, userID)

	// Filter berdasarkan judul
	if payload.Title != "" {
		query += " AND notes.title LIKE ?"
		args = append(args, "%"+payload.Title+"%")
	}

	// Paginasi
	query += " ORDER BY favorites.created_at DESC LIMIT ? OFFSET ?"
	args = append(args, payload.Limit, offset)

	rows, err := initialize.DB.Raw(query, args...).Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var favorites []map[string]interface{}
	for rows.Next() {
		var noteID uint
		var title string
		if err := rows.Scan(&noteID, &title); err != nil {
			return nil, 0, err
		}
		favorites = append(favorites, map[string]interface{}{
			"note_id": noteID,
			"title":   title,
		})
	}

	// Query untuk menghitung total data
	countQuery := `
		SELECT COUNT(*)
		FROM favorites
		JOIN notes ON notes.id = favorites.note_id
		WHERE favorites.user_id = ?
	`
	argsCount := append([]interface{}{userID}, args[1:len(args)-2]...) // Hapus limit dan offset
	var total int64
	err = initialize.DB.Raw(countQuery, argsCount...).Scan(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return favorites, total, nil
}
