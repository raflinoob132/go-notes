package credential

import (
	"encoding/json"
	"log"
	"net/http"

	//"github.com/raflinoob132/go-notes/credential/usermodel"
	"time"

	"github.com/raflinoob132/go-notes/initialize"
	"github.com/raflinoob132/go-notes/models"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var payload models.User

	// Decode JSON request ke struct
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, `{"message": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}
	log.Printf("Payload: %+v\n", payload)

	// Validasi input
	if payload.Username == "" || payload.Email == "" || payload.Password == "" {
		http.Error(w, `{"message": "All fields are required"}`, http.StatusBadRequest)
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, `{"message": "Failed to hash password"}`, http.StatusInternalServerError)
		return
	}

	//--

	query := `INSERT INTO users (username, email, password, created_at) VALUES (?, ?, ?, ?)`
	if err := initialize.DB.Exec(query, payload.Username, payload.Email, string(hashedPassword), time.Now()).Error; err != nil {
		http.Error(w, `{"message": "Failed to save user"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "User registered successfully"}`))
}

// Simpan user ke database
// user := usermodel.User{
// 	Username:  payload.Username,
// 	Email:     payload.Email,
// 	Password:  string(hashedPassword),
// 	CreatedAt: time.Now(),
// }
//versi gorm
// package credential

// import (
// 	"encoding/json"
// 	"net/http"
// 	"github.com/raflinoob132/go-notes/credential/usermodel"
// 	"github.com/raflinoob132/go-notes/initialize"
// 	"time"

// 	"golang.org/x/crypto/bcrypt"
// )

// func RegisterHandler(w http.ResponseWriter, r *http.Request) {
// 	var payload usermodel.User

// 	// Decode JSON request ke struct
// 	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
// 		http.Error(w, {"message": "Invalid request payload"}, http.StatusBadRequest)
// 		return
// 	}

// 	// Validasi input
// 	if payload.Username == "" || payload.Email == "" || payload.Password == "" {
// 		http.Error(w, {"message": "All fields are required"}, http.StatusBadRequest)
// 		return
// 	}

// 	// Hash password
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		http.Error(w, {"message": "Failed to hash password"}, http.StatusInternalServerError)
// 		return
// 	}

// 	// Simpan user ke database
// 	user := usermodel.User{
// 		Username:  payload.Username,
// 		Email:     payload.Email,
// 		Password:  string(hashedPassword),
// 		CreatedAt: time.Now(),
// 	}

// 	if err := initialize.DB.Create(&user).Error; err != nil {
// 		http.Error(w, {"message": "Failed to save user"}, http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	w.Write([]byte({"message": "User registered successfully"}))
// }
