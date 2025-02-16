package credential

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // Sama dengan secret key pada login

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ambil token dari header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, `{"message": "Missing token"}`, http.StatusUnauthorized)
			return
		}

		// Pastikan format "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, `{"message": "Invalid token format"}`, http.StatusUnauthorized)
			return
		}
		tokenString := parts[1]

		// Parse dan verifikasi token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Pastikan algoritma signing cocok
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return jwtSecret, nil
		})
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, `{"message": "Invalid token"}`, http.StatusUnauthorized)
			return
		}

		// Ambil klaim dari token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, `{"message": "Invalid token claims"}`, http.StatusUnauthorized)
			return
		}

		// Ambil user_id dari klaim
		userID, ok := claims["user_id"].(float64) // JWT menyimpan angka sebagai float64
		if !ok {
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, `{"message": "Invalid user ID in token"}`, http.StatusUnauthorized)
			return
		}

		// Simpan user_id ke context
		ctx := context.WithValue(r.Context(), "userID", uint(userID))

		// Lanjutkan ke handler berikutnya dengan context baru
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
