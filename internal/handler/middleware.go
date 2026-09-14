package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/SaH1l191/ticket-system/internal/auth"
)

type contextKey string
const userIDKey contextKey = "user_id"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		header := r.Header.Get("Authorization")

		if header == "" {
			writeJSON(w, http.StatusUnauthorized, map[string]string{
				"error": "missing authorization header",
			})
			return
		}

		//  Bearer <token>
		parts := strings.Fields(header)

		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			writeJSON(w, http.StatusUnauthorized, map[string]string{
				"error": "invalid authorization format",
			})
			return
		}

		userID, err := auth.ValidateToken(parts[1])

		if err != nil {
			writeJSON(w, http.StatusUnauthorized, map[string]string{
				"error": "invalid or expired token",
			})
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
