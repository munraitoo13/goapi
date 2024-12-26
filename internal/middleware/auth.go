package middleware

import (
	"context"
	"net/http"
	"strings"

	db "github.com/munraitoo13/goapi/database"
)

type ctx string

const ctxKey ctx = "client" // defined own type to avoid collisions

func TokenAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var clientId = r.URL.Query().Get("clientId") // gets clientId by query
		client, ok := db.Clients[clientId]           // gets client from db using query id

		// error checking
		if !ok || clientId == "" {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		token := r.Header.Get("Authorization") // gets the token from header
		if !isValidToken(client, token) {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}

		ctx := context.WithValue(r.Context(), ctxKey, client) // new ctx
		r = r.WithContext(ctx)                                // assign the context to the request variable

		next.ServeHTTP(w, r) // proceeds the code execution

	}
}

// checks validity of the token
func isValidToken(client db.Client, token string) bool {
	if strings.HasPrefix(token, "Bearer") {
		return strings.TrimPrefix(token, "Bearer ") == client.Token
	}

	return false
}
