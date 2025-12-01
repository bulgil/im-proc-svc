package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type RequestIDCtxKey struct{}

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := generateRequestID()

		ctx := context.WithValue(r.Context(), RequestIDCtxKey{}, id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func generateRequestID() string {
	return uuid.New().String()
}

func GetRequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	if requestID, ok := ctx.Value(RequestIDCtxKey{}).(string); ok {
		return requestID
	}

	return ""
}
