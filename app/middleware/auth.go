package middleware

import (
	"fmt"
	"log"
	"net/http"
)

// Auth will check session
func (h Handler) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Every protected request should send session_token in cookie
		c, err := r.Cookie("session_token")
		if err != nil && err == http.ErrNoCookie {
			log.Println(fmt.Errorf("COOKIE: %s", err.Error()))
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized: Token not found."))
			return
		}

		sessionToken := c.Value
		authenticated, err := h.isAuthenticated(sessionToken)
		if err != nil {
			log.Println(fmt.Errorf("TOKEN: %s", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Opps... Something went wrong."))
			return
		}

		if authenticated == false {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized: Token invalid."))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (h Handler) isAuthenticated(token string) (bool, error) {
	value, err := h.Cache.Get("session_token").Result()
	if err != nil {
		return false, err
	}

	return value == token, nil
}
