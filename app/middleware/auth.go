package middleware

import (
	"fmt"
	"log"
	"net/http"
)

// Auth will check session
func (m Middleware) Auth(next http.HandlerFunc) http.HandlerFunc {
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
		currentSession, err := m.getSessionByToken(sessionToken)
		if err != nil {
			log.Println(fmt.Errorf("TOKEN (%s): %s", sessionToken, err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Unauthorized: Token invalid or already been expired."))
			return
		}

		if currentSession == "" {
			log.Println(fmt.Errorf(`TOKEN: "GET %s" returns nil`, sessionToken))
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized: User not found."))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m Middleware) getSessionByToken(token string) (string, error) {
	return m.Cache.Get(token).Result()
}
