package query

import (
	"replika-golang-fiber/initialize"
	"replika-golang-fiber/models"
	"time"
)

func GetNotes(payload models.SearchPayload) ([]map[string]interface{}, int64, error) {
	// Default nilai pagination
	if payload.Page < 1 {
		payload.Page = 1
	}
	if payload.Limit < 1 {
		payload.Limit = 10
	}

	offset := (payload.Page - 1) * payload.Limit

	query := `
		SELECT id,title, category, created_at
		FROM notes
		WHERE 1=1
	`
	var args []interface{}

	// Filter berdasarkan title
	if payload.Title != "" {
		query += " AND title LIKE ?"
		args = append(args, "%"+payload.Title+"%")
	}

	// Filter berdasarkan semua kolom
	if payload.SearchAll != "" {
		query += " AND (title LIKE ? OR category LIKE ?)"
		args = append(args, "%"+payload.SearchAll+"%", "%"+payload.SearchAll+"%")
	}
	// } else if payload.SearchTitle != "" {
	// 	query += " AND (title LIKE ?)"
	// 	args = append(args, "%"+payload.SearchTitle+"%")
	// } else if payload.SearchCat != "" {
	// 	query += " AND (category LIKE ?)"

	// 	args = append(args, "%"+payload.SearchCat+"%")
	// }

	query += " LIMIT ? OFFSET ?"
	args = append(args, payload.Limit, offset)

	rows, err := initialize.DB.Raw(query, args...).Rows()
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var notes []map[string]interface{}
	for rows.Next() {
		var noteID uint
		var title, category string
		var createdAt time.Time
		if err := rows.Scan(&noteID, &title, &category, &createdAt); err != nil {
			return nil, 0, err
		}
		notes = append(notes, map[string]interface{}{
			"id":         noteID,
			"title":      title,
			"category":   category,
			"created_at": createdAt,
		})
	}

	// Query untuk menghitung total data
	countQuery := `
		SELECT COUNT(*)
		FROM notes
		WHERE 1=1
	`
	argsCount := append([]interface{}{}, args[:len(args)-2]...) // Hapus limit dan offset
	if payload.Title != "" {
		countQuery += " AND title LIKE ?"
	}

	if payload.SearchAll != "" {
		countQuery += " AND (title LIKE ? OR category LIKE ?)"
	}

	var total int64
	err = initialize.DB.Raw(countQuery, argsCount...).Scan(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return notes, total, nil
}
