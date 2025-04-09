package middleware

import (
	"context"
	"net/http"

	authpb "cec-project/delivery-app/auth/proto/gen"
)

func Auth(authClient authpb.AuthServiceClient) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Skip auth for login/register
			if r.URL.Path == "/api/v1/auth/login" ||
				r.URL.Path == "/api/v1/auth/register" {
				next.ServeHTTP(w, r)
				return
			}

			token := r.Header.Get("Authorization")
			if token == "" {
				http.Error(w, "missing auth token", http.StatusUnauthorized)
				return
			}

			// Validate token with auth service
			res, err := authClient.ValidateToken(r.Context(), &authpb.ValidateTokenRequest{
				Token: token,
			})

			if err != nil || !res.Valid {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			// Add user info to context
			ctx := context.WithValue(r.Context(), "userID", res.UserId)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
