package credential

import (
	"encoding/json"
	"net/http"

	//"replika-golang-fiber/credential/models"
	"replika-golang-fiber/initialize"
	"replika-golang-fiber/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// kelemahan: token tidak berganti ganti
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	payload := models.LoginPayload{}

	// Decode JSON request
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, `{"message": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	// Cek apakah username/email ada di database menggunakan raw query
	query := `SELECT id, username, email, password FROM users WHERE username = ? OR email = ?`
	row := initialize.DB.Raw(query, payload.UsernameOrEmail, payload.UsernameOrEmail).Row()

	var userID uint
	var username, email, hashedPassword string
	if err := row.Scan(&userID, &username, &email, &hashedPassword); err != nil {
		http.Error(w, `{"message": "Invalid username/email or password"}`, http.StatusUnauthorized)
		return
	}

	// Verifikasi password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(payload.Password)); err != nil {
		http.Error(w, `{"message": "Invalid username/email or password"}`, http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, `{"message": "Failed to generate token"}`, http.StatusInternalServerError)
		return
	}

	// Kirimkan token ke client
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"token": "` + tokenString + `"}`))
}

//versi orm
// func LoginHandler(w http.ResponseWriter, r *http.Request) {
// 	var payload struct {
// 		UsernameOrEmail string json:"username_or_email"
// 		Password        string json:"password"
// 	}
// 	var dbUser models.User

// 	// Decode JSON request
// 	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
// 		http.Error(w, {"message": "Invalid request payload"}, http.StatusBadRequest)
// 		return
// 	}

// 	// Cek apakah username/email ada di database
// 	if err := initialize.DB.Where("username = ? OR email = ?", payload.UsernameOrEmail, payload.UsernameOrEmail).First(&dbUser).Error; err != nil {
// 		http.Error(w, {"message": "Invalid username/email or password"}, http.StatusUnauthorized)
// 		return
// 	}

// 	// Verifikasi password
// 	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(payload.Password)); err != nil {
// 		http.Error(w, {"message": "Invalid username/email or password"}, http.StatusUnauthorized)
// 		return
// 	}

// 	// Generate JWT token
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"user_id": dbUser.ID,
// 		"exp":     time.Now().Add(24 * time.Hour).Unix(),
// 	})

// 	tokenString, err := token.SignedString(jwtSecret)
// 	if err != nil {
// 		http.Error(w, {"message": "Failed to generate token"}, http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write([]byte({"token": " + tokenString + "}))
// }
