package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func (app *application) BasicAuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// read the authorization header
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				app.unauthorizedBasicErrorResponse(w, r, fmt.Errorf("authorization header is missing"))
				return
			}

			// parse the authorization header -> get the base64
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Basic" {
				app.unauthorizedBasicErrorResponse(w, r, fmt.Errorf("invalid authorization header"))
				return
			}

			// decode it
			decoded, err := base64.StdEncoding.DecodeString(parts[1])
			if err != nil {
				app.unauthorizedBasicErrorResponse(w, r, err)
				return
			}

			username := app.config.auth.basic.user
			password := app.config.auth.basic.pass

			// check the credentials
			creds := strings.SplitN(string(decoded), ":", 2)
			if len(creds) != 2 || creds[0] != username || creds[1] != password {
				app.unauthorizedBasicErrorResponse(w, r, fmt.Errorf("invalid credentials"))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
