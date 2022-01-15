package auth

import (
	"encoding/base64"
	"github.com/leganck/simple-dns/config"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func BasicAuth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		credentials := config.GetCredentials()
		if len(credentials) == 0 {
			handler.ServeHTTP(w, r)
			return
		}
		auth := r.Header.Get("Authorization")
		if auth != "" {
			username, password, ok := parseBasicAuth(auth)
			if ok {
				validPassword, userFound := credentials[username]
				if userFound && validPassword == password {
					handler.ServeHTTP(w, r)
					return
				}
				log.Info("user: %s fail login", username)
			}
		}
		unauthorized(w)
		return
	})
}

func unauthorized(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
	// 401 status code
	w.WriteHeader(http.StatusUnauthorized)
}

func parseBasicAuth(auth string) (username, password string, ok bool) {
	const prefix = "Basic "
	// Case insensitive prefix match. See Issue 22736.
	if len(auth) < len(prefix) || !strings.EqualFold(auth[:len(prefix)], prefix) {
		return
	}
	c, err := base64.StdEncoding.DecodeString(auth[len(prefix):])
	if err != nil {
		return
	}
	cs := string(c)
	s := strings.IndexByte(cs, ':')
	if s < 0 {
		return
	}
	return cs[:s], cs[s+1:], true
}
