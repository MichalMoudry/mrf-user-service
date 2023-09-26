package util

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// A middleware for adding context from a JWT token.
func ParseJWT() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := parseBearerToken(r.Header)
			if tokenString == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			token, _, err := jwt.NewParser().ParseUnverified(tokenString, nil)
			if err != nil {
				WriteErrResponse(w, http.StatusBadRequest, err)
				return
			}

			userId, err := token.Claims.GetSubject()
			if err != nil {
				WriteErrResponse(w, http.StatusBadRequest, err)
				return
			}

			ctx := WithUserId(r.Context(), userId)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Functon for trimming recieved JWT.
func parseBearerToken(h http.Header) string {
	if h == nil {
		return ""
	}
	return strings.TrimPrefix(h.Get("Authorization"), "Bearer ")
}
